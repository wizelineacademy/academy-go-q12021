package repository

import (
	"bootcamp/domain/model"
	"bootcamp/interface/controller/vo"
	"bootcamp/usecase/repository"
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

// itemRepository struct to SQL DB
type itemRepository struct {
	db *sql.DB
}

// NewItemRepository returns an ItemRepository
func NewItemRepository(db *sql.DB) repository.ItemRepository {
	return &itemRepository{db}
}

// FindAll returns all Items stored in the datastore.
func (ir *itemRepository) FindAll(items []*model.Item) ([]*model.Item, error) {

	fmt.Println("lenght=== %s", len(items))

	// SQL query to get all records from CSV file
	queryString := "SELECT id, name FROM items"

	if len(items) > 0 {
		var ids []string
		for _, v := range items {
			ids = append(ids,fmt.Sprint(v.ID))
		}
		queryString = fmt.Sprintf("%s WHERE id IN (%s)", queryString, strings.Join(ids, ","))
	}

	items = []*model.Item{}

	fmt.Println("QUERY === %s", queryString)

	// Execute query
	rows, err := ir.db.Query(queryString)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := model.Item{}
		if err = rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil

}

// Create adds a new Item in the datastore and return it.
func (ir *itemRepository) Create(item *model.Item) (*model.Item, error) {


	resulset, err := ir.db.Exec("","")
	if err != nil {
		return nil, err
	}

	resulset.RowsAffected()

	return item, nil
}

func (ir *itemRepository) FindAllPaged(items []*model.Item, paged *vo.Paged) ([]*model.Item, error) {

	queryString := "SELECT count(id) FROM items"

	// Get among of items
	rs, err := ir.db.Query(queryString)
	r := 0
	if rs.Next() {
		rs.Scan(&r)
	}

	workers := r/paged.ItemsPerWorkers

	//  get all items
	queryString = "SELECT id, name FROM items"
	rows, err := ir.db.Query(queryString)
	if err != nil {
		return nil, err
	}

	// channels
	chanIn := make(chan model.Item, r)
	chanOut := make(chan model.Item, paged.Items)

	var waitGroup sync.WaitGroup
	waitGroup.Add(workers)

	for i:=0; i<workers;i++ {
		go findOddOrEven(chanIn, chanOut, *paged, &waitGroup)
	}

	for rows.Next() {
		item := model.Item{}
		if err = rows.Scan(&item.ID, &item.Name); err != nil {
			return nil, err
		}
		chanIn <- item
	}
	waitGroup.Wait()

	close(chanIn)
	close(chanOut)

	for i:=0; i<paged.Items;i++{
		item := <-chanOut
		if item.ID>0 {
			items = append(items, &item)
		}
	}


	fmt.Println("Items ", len(items))

	return items, nil
}

func findOddOrEven(chanIn <-chan model.Item, chanOut chan<- model.Item, paged vo.Paged, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	fmt.Println("worker working...")

	for i:=0; i < paged.ItemsPerWorkers; i++ {
		itm := <- chanIn
		module := itm.ID%2
		if paged.Type == "odd" && module == 0 {
			chanOut <- itm
		} else if paged.Type == "even" && module != 0 {
			chanOut <- itm
		}
	}
}
