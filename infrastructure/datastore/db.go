package datastore

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"	

	"api-booking-time/domain/model"
	"api-booking-time/usecase/repository"
	ir "api-booking-time/interface/repository"
)

type DbRepository struct {
	Centres interface{ repository.CentreRepository }
}

func OpenDb() *DbRepository {
	centres := initialize()

	return &DbRepository{Centres: centres}
}

func initialize() repository.CentreRepository {
	filename := "./assets/centres.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	return ir.OpenCentreRepository(loadData(file))
}

func loadData(r io.Reader) *[]*model.Centre {
	reader := csv.NewReader(r)

	ret := make([]*model.Centre, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		id, _ := strconv.Atoi(row[0])
		capacity, _ := strconv.Atoi(row[6])
		openness, _ := strconv.Atoi(row[7])

		if err != nil {
			log.Println(err)
		}
		centre := &model.Centre{
			Id:			id,
			Capacity:	capacity,
			Openness:	openness,
			Name:		row[1],
			Address:	row[2],
			Email:		row[3],
			Phone:		row[4],
			Line:		row[5],
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, centre)
	}
	return &ret
}