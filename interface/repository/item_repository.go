package repository

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/repository"
	"database/sql"
	"fmt"
	"strings"
)

type itemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) repository.ItemRepository {
	return &itemRepository{db}
}

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


func (ir *itemRepository) Create(item *model.Item) (*model.Item, error) {


	resulset, err := ir.db.Exec("","")
	if err != nil {
		return nil, err
	}

	resulset.RowsAffected()

	return item, nil
}