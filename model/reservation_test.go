package model

import (
	"testing"
	"time"
)

func TestReservation_totalFee(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)

	type fields struct {
		Adults    int
		Minors    int
		AdultFee  float64
		MinorFee  float64
		Arrival   *time.Time
		Departure *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			"Adults",
			fields{
				Adults:    2,
				Minors:    0,
				AdultFee:  7,
				MinorFee:  0,
				Arrival:   &arrival,
				Departure: &departure,
			}, 2 * 7 * 2,
		},
		{
			"Adults and Minors",
			fields{
				Adults:    2,
				Minors:    2,
				AdultFee:  7,
				MinorFee:  1,
				Arrival:   &arrival,
				Departure: &departure,
			}, ((2 * 7) + (2 * 1)) * 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reservation{
				Adults:    tt.fields.Adults,
				Minors:    tt.fields.Minors,
				AdultFee:  tt.fields.AdultFee,
				MinorFee:  tt.fields.MinorFee,
				Arrival:   tt.fields.Arrival,
				Departure: tt.fields.Departure,
			}
			if got := r.totalFee(); got != tt.want {
				t.Errorf("Reservation.totalFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReservation_nights(t *testing.T) {
	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)

	departureSameDay := time.Date(2020, 1, 1, 18, 0, 0, 0, time.UTC)
	departureLessThan24Hours := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	departureMoreThan24Hours := time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)
	departureLessThan48Hours := time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)
	departureMoreThan48Hours := time.Date(2020, 1, 3, 12, 0, 0, 0, time.UTC)

	type fields struct {
		Arrival   *time.Time
		Departure *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"Same day",
			fields{
				Arrival:   &arrival,
				Departure: &departureSameDay,
			}, 0,
		},
		{
			"Less than 24 hours",
			fields{
				Arrival:   &arrival,
				Departure: &departureLessThan24Hours,
			}, 1,
		},
		{
			"More than 24 hours",
			fields{
				Arrival:   &arrival,
				Departure: &departureMoreThan24Hours,
			}, 1,
		},
		{
			"Less than 48 hours",
			fields{
				Arrival:   &arrival,
				Departure: &departureLessThan48Hours,
			}, 2,
		},
		{
			"More than 48 hours",
			fields{
				Arrival:   &arrival,
				Departure: &departureMoreThan48Hours,
			}, 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reservation{
				Arrival:   tt.fields.Arrival,
				Departure: tt.fields.Departure,
			}
			if got := r.nights(); got != tt.want {
				t.Errorf("Reservation.nights() = %v, want %v", got, tt.want)
			}
		})
	}
}
