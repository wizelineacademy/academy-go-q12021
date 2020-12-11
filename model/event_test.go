package model

import (
	"reflect"
	"testing"
	"time"
)

func TestEvent_CalculateTotalFee(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)

	type fields struct {
		Reservations []Reservation
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			"Adults",
			fields{
				[]Reservation{
					{Adults: 2,
						Minors:    0,
						AdultFee:  7,
						MinorFee:  0,
						Arrival:   &arrival,
						Departure: &departure,
					},
					{
						Adults:    2,
						Minors:    2,
						AdultFee:  7,
						MinorFee:  1,
						Arrival:   &arrival,
						Departure: &departure,
					},
				},
			}, (2 * 7 * 2) + (((2 * 7) + (2 * 1)) * 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Reservations: tt.fields.Reservations,
			}
			e.CalculateTotalFee()
		})
	}
}

func TestEvent_filterReservations(t *testing.T) {
	type fields struct {
		Reservations []Reservation
	}
	type args struct {
		status string
	}
	tests := []struct {
		name   string
		fields fields
		status string
		want   int
	}{
		{
			"Simple search",
			fields{
				[]Reservation{
					{Status: "CREATED"},
					{Status: "CREATED"},
					{Status: "CANCELED"},
					{Status: "PAID"},
				},
			},
			"CREATED",
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				Reservations: tt.fields.Reservations,
			}
			if got := e.filterReservations(tt.status); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("Event.filterReservations() = %v, want %v", got, tt.want)
			}
		})
	}
}
