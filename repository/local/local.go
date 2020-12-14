package local

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"
)

// collection keeps the table name
const collection = "reservations"

type reservations struct {
	csvFile *os.File
}

// NewReservationRepo returns the ReservationRepository implementation
func NewReservationRepo(csvFile *os.File) repository.ReservationRepository {
	return &reservations{
		csvFile,
	}
}

func (r *reservations) Create(eventID string, reservation model.Reservation) (model.Reservation, error) {
	writer := csv.NewWriter(r.csvFile)

	insert := []string{
		eventID,
		reservation.ID,
		reservation.Status,
		reservation.Plan,
		strconv.Itoa(reservation.Adults),
		strconv.Itoa(reservation.Minors),
		strconv.FormatFloat(reservation.AdultFee, 'f', -1, 64),
		strconv.FormatFloat(reservation.MinorFee, 'f', -1, 64),
		strconv.Itoa(int(reservation.Arrival.Unix())),
		strconv.Itoa(int(reservation.Departure.Unix())),
		reservation.Name,
		reservation.Phone,
		reservation.Email,
	}

	err := writer.Write(insert)
	if err != nil {
		return model.Reservation{}, err
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return model.Reservation{}, err
	}

	return reservation, nil
}

func (r *reservations) GetByEventID(id string) ([]model.Reservation, error) {
	reader := csv.NewReader(r.csvFile)

	var response []model.Reservation

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if record[0] == id {
			reservation, err := recordToReservation(record)
			if err != nil {
				return nil, err
			}

			response = append(response, reservation)
		}
	}

	_, err := r.csvFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func recordToReservation(record []string) (model.Reservation, error) {
	adults, err := strconv.Atoi(record[4])
	if err != nil {
		return model.Reservation{}, err
	}

	minors, err := strconv.Atoi(record[5])
	if err != nil {
		return model.Reservation{}, err
	}

	adultFee, err := strconv.ParseFloat(record[6], 64)
	if err != nil {
		return model.Reservation{}, err
	}

	minorFee, err := strconv.ParseFloat(record[7], 64)
	if err != nil {
		return model.Reservation{}, err
	}

	arrivalTimestamp, err := strconv.Atoi(record[8])
	if err != nil {
		return model.Reservation{}, err
	}
	arrival := time.Unix(int64(arrivalTimestamp), 0)

	departureTimestamp, err := strconv.Atoi(record[8])
	if err != nil {
		return model.Reservation{}, err
	}
	departure := time.Unix(int64(departureTimestamp), 0)

	return model.Reservation{
		ID:        record[1],
		Status:    record[2],
		Plan:      record[3],
		Adults:    adults,
		Minors:    minors,
		AdultFee:  adultFee,
		MinorFee:  minorFee,
		Arrival:   &arrival,
		Departure: &departure,
		Name:      record[10],
		Phone:     record[11],
		Email:     record[12],
	}, nil
}
