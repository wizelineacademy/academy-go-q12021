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
	pokeMap map[int32]*model.Pokemon
	file string
}

func NewCsvStorage(file string) (*CsvPokeStorage, error) {
	source, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(source)
	pokeMap := make(map[int32]*model.Pokemon)
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
			fmt.Errorf("Incorrect amount of fields(5) in :", record)
		}
		var id int32
		if val, err := strconv.Atoi(record[0]); err != nil {
			fmt.Errorf("Error  parsing line ", record)
			continue
		} else {
			id = int32(val)
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

func (c *CsvPokeStorage) FindById(id int32) *model.Pokemon {
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
	target, err := os.OpenFile(c.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer target.Close()
	if err != nil {
		return nil, err
	}
	writer := csv.NewWriter(target)
	defer writer.Flush()
	err = writer.Write([]string{
		fmt.Sprintf("%d",pokemon.Id),
		pokemon.Species,
		pokemon.Sprite,
		strings.Join(pokemon.Types,","),
		pokemon.FlavorText})
	if err != nil {
		return nil, err
	}
	c.pokeMap[pokemon.Id] = pokemon
	return pokemon, nil
}

func (c *CsvPokeStorage) FindAllWorkers(typeStr string, items int, itemsPerWorker int) ([]*model.Pokemon, error) {
	poolSize := runtime.GOMAXPROCS(0)
	log.Println("Max pool size of ", poolSize)
	var wg sync.WaitGroup
	wg.Add(poolSize)
	key := make(chan int)
	latestIdx := make(chan int)
	shutdown := make(chan struct{})
	keys := []int{}
	for key := range c.pokeMap {
		keys = append(keys, int(key))
	}
	mod := 1
	if typeStr == "odd"{
		mod = 2
	}
	for i := 0; i < poolSize; i++ {
		w := &Worker{
			id:                i,
			key:               key,
			latestIdx:         latestIdx,
			mod:			   mod,
			itemsProcessed:    0,
			maxItemsPerWorker: itemsPerWorker,
			keys:              keys,
			shutdown:          shutdown,
			wg:                &wg,
		}
		go w.Work()
	}
	pokemons := []*model.Pokemon{}
	latestIdx <- 0
	for ;len(pokemons) <= items; {
		k := <-key
		if k == -1 {
			break
		}
		pokemons = append(pokemons, c.pokeMap[int32(keys[k])])
		latestIdx <- k
	}
	fmt.Println("Receiver sending shutdown signal")
	close(shutdown)
	wg.Wait()
	return pokemons, nil
}