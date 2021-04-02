package datastore

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/ToteEmmanuel/academy-go-q12021/domain/model"
)

type CsvPokeStorage struct {
	pokeMap map[int]*model.Pokemon
	file    string
}

func NewCsvStorage(file string) (PokeStorage, error) {
	source, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(source)
	pokeMap := make(map[int]*model.Pokemon)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(record) != 5 {
			fmt.Errorf("Incorrect amount of fields(5) in : %s", record)
		}
		var id int
		if val, err := strconv.Atoi(record[0]); err != nil {
			fmt.Errorf("Error  parsing line %s", record)
			continue
		} else {
			id = int(val)
		}
		pokeMap[id] = &model.Pokemon{
			Id:         id,
			Species:    record[1],
			Sprite:     record[2],
			Types:      strings.Split(record[3], ","),
			FlavorText: record[4],
		}
	}
	return &CsvPokeStorage{pokeMap, file}, nil
}

func (c *CsvPokeStorage) FindById(id int) *model.Pokemon {
	return c.pokeMap[id]
}

func (c *CsvPokeStorage) FindAll() []*model.Pokemon {
	allPokes := make([]*model.Pokemon, 0, len(c.pokeMap))
	for _, v := range c.pokeMap {
		allPokes = append(allPokes, v)
	}
	return allPokes
}

func (c *CsvPokeStorage) Save(pokemon *model.Pokemon) (*model.Pokemon, error) {
	if c.pokeMap[pokemon.Id] != nil {
		return c.pokeMap[pokemon.Id], nil
	}
	target, err := os.OpenFile(c.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer target.Close()
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(target)
	defer writer.Flush()
	err = writer.Write([]string{
		fmt.Sprintf("%d", pokemon.Id),
		pokemon.Species,
		pokemon.Sprite,
		strings.Join(pokemon.Types, ","),
		pokemon.FlavorText})
	if err != nil {
		return nil, err
	}
	c.pokeMap[pokemon.Id] = pokemon
	return pokemon, nil
}

func (c *CsvPokeStorage) FindAllWorkers(typeStr string, items int, itemsPerWorker int) ([]*model.Pokemon, error) {
	pokemons := []*model.Pokemon{}
	if items == 0 || itemsPerWorker == 0 {
		return pokemons, nil
	}

	keys := []int{}
	for key := range c.pokeMap {
		keys = append(keys, int(key))
	}
	odd := typeStr == "odd"
	poolSize := runtime.GOMAXPROCS(0)
	wg, keyIdx, latestIdx, shutdown := c.prepareWorkers(poolSize, itemsPerWorker, keys, odd)
	k := 0
	for len(pokemons) < items && len(pokemons) < poolSize*itemsPerWorker {
		log.Println("Sending", k)
		latestIdx <- k
		k = <-keyIdx
		log.Println("Received", k)
		if k == -1 {
			break
		}
		pokemons = append(pokemons, c.pokeMap[int(keys[k])])
	}
	log.Println("Receiver sending shutdown signal")
	close(shutdown)
	log.Println("Waiting for workers to shutdown.")
	wg.Wait()
	return pokemons, nil
}

func (c *CsvPokeStorage) prepareWorkers(poolSize, itemsPerWorker int, keys []int, odd bool) (*sync.WaitGroup, chan int, chan int, chan struct{}) {
	log.Println("Max pool size of ", poolSize)
	var wg sync.WaitGroup
	wg.Add(poolSize)
	keyIdx := make(chan int)
	latestIdx := make(chan int)
	shutdown := make(chan struct{})

	for i := 0; i < poolSize; i++ {
		w := &Worker{
			id:                i,
			key:               keyIdx,
			latestIdx:         latestIdx,
			odd:               odd,
			itemsProcessed:    0,
			maxItemsPerWorker: itemsPerWorker,
			keys:              keys,
			shutdown:          shutdown,
			wg:                &wg,
		}
		go w.Work()
	}
	return &wg, keyIdx, latestIdx, shutdown
}
