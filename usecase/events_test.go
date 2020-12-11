package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	mocks "github.com/javiertlopez/golang-bootcamp-2020/mocks/repository"
	"github.com/javiertlopez/golang-bootcamp-2020/model"
)

func Test_events_Create(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	reservation := model.Reservation{
		Adults:    2,
		Minors:    0,
		AdultFee:  7,
		MinorFee:  0,
		Arrival:   &arrival,
		Departure: &departure,
	}

	tests := []struct {
		name    string
		event   model.Event
		want    model.Event
		wantErr bool
	}{
		{
			"Success (with reservations)",
			model.Event{
				Reservations: []model.Reservation{reservation},
			},
			model.Event{
				ID:           "123",
				Reservations: []model.Reservation{reservation},
				TotalFee:     2 * 7 * 2,
			},
			false,
		},
		{
			"Success",
			model.Event{
				Reservations: nil,
			},
			model.Event{
				ID:           "123",
				Reservations: nil,
			},
			false,
		},
		{
			"Error",
			model.Event{},
			model.Event{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventRepo := &mocks.EventRepository{}
			reservationRepo := &mocks.ReservationRepository{}
			if tt.wantErr {
				eventRepo.On("Create", tt.event).Return(tt.want, errors.New("failed"))
			} else {
				eventRepo.On("Create", tt.event).Return(tt.want, nil)
				reservationRepo.On("Create", "123", reservation).Return(reservation, nil)
			}

			e := &events{
				eventRepo:       eventRepo,
				reservationRepo: reservationRepo,
			}

			got, err := e.Create(tt.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("events.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("events.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_GetByID(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	reservations := []model.Reservation{
		{
			Adults:    2,
			Minors:    0,
			AdultFee:  7,
			MinorFee:  0,
			Arrival:   &arrival,
			Departure: &departure,
		},
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		id      string
		want    model.Event
		wantErr bool
	}{
		{
			"Success (with reservations)",
			"123",
			model.Event{
				ID:           "123",
				Reservations: reservations,
				TotalFee:     2 * 7 * 2,
			},
			false,
		},
		{
			"Success",
			"456",
			model.Event{
				ID:           "456",
				Reservations: []model.Reservation{},
				TotalFee:     0,
			},
			false,
		},
		{
			"Error",
			"123",
			model.Event{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventRepo := &mocks.EventRepository{}
			reservationRepo := &mocks.ReservationRepository{}
			if tt.wantErr {
				eventRepo.On("GetByID", tt.id).Return(tt.want, errors.New("failed"))
			} else {
				eventRepo.On("GetByID", tt.id).Return(tt.want, nil)
				reservationRepo.On("GetByEventID", "123").Return(reservations, nil)
				reservationRepo.On("GetByEventID", "456").Return([]model.Reservation{}, nil)
			}

			e := &events{
				eventRepo:       eventRepo,
				reservationRepo: reservationRepo,
			}
			got, err := e.GetByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("events.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("events.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_GetAll(t *testing.T) {
	allEvents := []model.Event{
		{
			ID: "123",
		},
		{
			ID: "456",
		},
	}

	tests := []struct {
		name    string
		want    []model.Event
		wantErr bool
	}{
		{
			"Success",
			[]model.Event{
				{
					ID: "123",
				},
				{
					ID: "456",
				},
			},
			false,
		},
		{
			"Error",
			nil,
			true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventRepo := &mocks.EventRepository{}
			reservationRepo := &mocks.ReservationRepository{}
			if tt.wantErr {
				eventRepo.On("GetAll").Return(nil, errors.New("failed"))
			} else {
				eventRepo.On("GetAll").Return(allEvents, nil)
			}

			e := &events{
				eventRepo:       eventRepo,
				reservationRepo: reservationRepo,
			}

			got, err := e.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("events.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("events.GetAll() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_events_AddReservations(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	reservation := model.Reservation{
		Adults:    2,
		Minors:    0,
		AdultFee:  7,
		MinorFee:  0,
		Arrival:   &arrival,
		Departure: &departure,
	}
	reservationResult := model.Reservation{
		ID:        "456",
		Adults:    2,
		Minors:    0,
		AdultFee:  7,
		MinorFee:  0,
		Arrival:   &arrival,
		Departure: &departure,
	}

	type args struct {
		id           string
		reservations []model.Reservation
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Reservation
		wantErr bool
	}{
		{
			"Success",
			args{
				"123",
				[]model.Reservation{
					{
						Adults:    2,
						Minors:    0,
						AdultFee:  7,
						MinorFee:  0,
						Arrival:   &arrival,
						Departure: &departure,
					},
				},
			},
			[]model.Reservation{
				{
					ID:        "456",
					Adults:    2,
					Minors:    0,
					AdultFee:  7,
					MinorFee:  0,
					Arrival:   &arrival,
					Departure: &departure,
				},
			},
			false,
		},
		{
			"Error",
			args{
				"456",
				[]model.Reservation{
					{
						Adults:    2,
						Minors:    0,
						AdultFee:  7,
						MinorFee:  0,
						Arrival:   &arrival,
						Departure: &departure,
					},
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservationRepo := &mocks.ReservationRepository{}
			if tt.wantErr {
				reservationRepo.On("Create", tt.args.id, reservation).Return(model.Reservation{}, errors.New("failed"))
			} else {
				reservationRepo.On("Create", tt.args.id, reservation).Return(reservationResult, nil)
			}

			e := &events{
				reservationRepo: reservationRepo,
			}
			got, err := e.AddReservations(tt.args.id, tt.args.reservations)
			if (err != nil) != tt.wantErr {
				t.Errorf("events.AddReservations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("events.AddReservations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_GetReservations(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	allReservations := []model.Reservation{
		{
			ID:        "123",
			Adults:    2,
			Minors:    0,
			AdultFee:  7,
			MinorFee:  0,
			Arrival:   &arrival,
			Departure: &departure,
		},
		{
			ID:        "456",
			Adults:    2,
			Minors:    0,
			AdultFee:  7,
			MinorFee:  0,
			Arrival:   &arrival,
			Departure: &departure,
		},
	}

	tests := []struct {
		name    string
		id      string
		want    []model.Reservation
		wantErr bool
	}{
		{
			"Success",
			"123",
			[]model.Reservation{
				{
					ID:        "123",
					Adults:    2,
					Minors:    0,
					AdultFee:  7,
					MinorFee:  0,
					Arrival:   &arrival,
					Departure: &departure,
				},
				{
					ID:        "456",
					Adults:    2,
					Minors:    0,
					AdultFee:  7,
					MinorFee:  0,
					Arrival:   &arrival,
					Departure: &departure,
				},
			},
			false,
		},
		{
			"Error",
			"456",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservationRepo := &mocks.ReservationRepository{}
			if tt.wantErr {
				reservationRepo.On("GetByEventID", tt.id).Return(nil, errors.New("failed"))
			} else {
				reservationRepo.On("GetByEventID", tt.id).Return(allReservations, nil)
			}

			e := &events{
				reservationRepo: reservationRepo,
			}
			got, err := e.GetReservations(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("events.GetReservations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("events.GetReservations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_events_validateReservation(t *testing.T) {
	tests := []struct {
		name         string
		reservations []model.Reservation
		wantErr      bool
	}{
		{
			"Success",
			[]model.Reservation{
				{
					Adults:   2,
					Minors:   0,
					AdultFee: 7,
					MinorFee: 0,
				},
			},
			false,
		},
		{
			"res.Adults < 1",
			[]model.Reservation{
				{
					Adults: 0,
				},
			},
			true,
		},
		{
			"res.AdultFee <= 0",
			[]model.Reservation{
				{
					Adults:   2,
					AdultFee: 0,
				},
			},
			true,
		},
		{
			"res.Minors < 0",
			[]model.Reservation{
				{
					Adults:   2,
					Minors:   -1,
					AdultFee: 7,
				},
			},
			true,
		},
		{
			"res.Minors > 0 && res.MinorFee <= 0",
			[]model.Reservation{
				{
					Adults:   2,
					Minors:   1,
					AdultFee: 7,
					MinorFee: 0,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &events{}
			if err := e.validateReservation(tt.reservations); (err != nil) != tt.wantErr {
				t.Errorf("events.validateReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
