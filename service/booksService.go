package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/joseantoniovz/academy-go-q12021/model"
	"github.com/joseantoniovz/academy-go-q12021/util"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Title  string `json:"title"`
	Author string `json:"authors"`
	Isbn13 int    `json:"isbn13"`
	Price  string `json:"price"`
	Image  string `json:"image"`
}

const pathFile = "./booklist.csv"

func GetAll() ([]model.Book, *model.Error) {
	f, err := util.Open(pathFile)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return nil, &err
	}

	books, errorReading := util.ReadCsv(f)

	if errorReading != nil {
		errorReading := model.Error{
			Code:    http.StatusInternalServerError,
			Message: errorReading,
		}
		return nil, &errorReading
	}
	return books, nil
}

func GetById(id string) (model.Book, *model.Error) {
	var books, err = GetAll()
	var bookResult model.Book
	for _, book := range books {
		if book.Id == id {
			bookResult = book
		}
	}
	return bookResult, err
}

func ConsumeAPI(id string) (model.Book, *model.Error) {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.itbook.store/1.0/books/9781617294136", nil)
	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return model.Book{}, &err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err := model.Error{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
		return model.Book{}, &err
	}

	var responseObject Response
	var result model.Book
	json.Unmarshal(bodyBytes, &responseObject)

	result.Id = strconv.Itoa(rand.Int())
	result.Title = responseObject.Title
	result.Author = responseObject.Author
	result.Format = "Digital"
	result.Price = responseObject.Price

	util.WriteInCSV(result, pathFile)
	fmt.Println("A new row has been added to the CSV file")
	return result, nil

}

func ConcurrencyBooks(typeNumber string, items, itemsPerWorker int) ([]model.Book, *model.Error) {
	log.Info("type: " + typeNumber)
	log.Info("items: " + strconv.Itoa(items))
	log.Info("workers: " + strconv.Itoa(itemsPerWorker))
	var books []model.Book = nil
	var wg sync.WaitGroup

	results := make(chan model.Book, items)

	for i := 0; i < items; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for data := range results {
				var book, err = GetById(data.Id)
				if err != nil {
					log.Error(err)
				}
				books = append(books, book)
			}
		}()
	}

	close(results)
	wg.Wait()

	return books, nil
}
