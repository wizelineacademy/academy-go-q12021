package controller

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	mocks "github.com/javiertlopez/golang-bootcamp-2020/mocks/usecase"
	"github.com/javiertlopez/golang-bootcamp-2020/model"
)

func Test_eventController_CreateEvent(t *testing.T) {
	events := &mocks.Events{}
	e := &eventController{
		events,
	}

	event := model.Event{
		Description: "Wedding",
		Status:      "NEW",
	}

	wrongEvent := model.Event{
		Description: "Graduation",
		Status:      "NEW",
	}

	expectedEvent := model.Event{
		ID:          "123",
		Description: "Wedding",
		Status:      "NEW",
	}

	events.On("Create", event).Return(expectedEvent, nil)
	events.On("Create", wrongEvent).Return(model.Event{}, errors.New("failed"))

	tests := []struct {
		name         string
		expectedCode int
		expectedBody string
		body         string
	}{
		{
			"Success",
			201,
			`{"id":"123","description":"Wedding","type":"","status":"NEW","created_at":null,"updated_at":null,"event_date":null,"event_location":"","name":"","phone":"","email":""}`,
			`{"description":"Wedding","status":"NEW"}`,
		},
		{
			"Error",
			500,
			`{"message":"Internal server error","status":500}`,
			`{"description":"Graduation","status":"NEW"}`,
		},
		{
			"Bad request",
			400,
			`{"message":"Bad request","status":400}`,
			`{"description":123,"status":"NEW}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler.
			req, err := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(tt.body)))
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(e.CreateEvent)

			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}
		})
	}
}

func Test_eventController_GetEventByID(t *testing.T) {
	events := &mocks.Events{}
	e := &eventController{
		events,
	}

	expectedEvent := model.Event{
		ID:          "123",
		Description: "Wedding",
		Status:      "NEW",
	}

	events.On("GetByID", "123").Return(expectedEvent, nil)
	events.On("GetByID", "456").Return(model.Event{}, errors.New("failed"))
	events.On("GetByID", "789").Return(model.Event{}, nil)
	events.On("GetReservations", "123").Return(nil, nil)
	events.On("GetReservations", "789").Return(nil, errors.New("failed"))

	tests := []struct {
		name         string
		id           string
		expectedCode int
		expectedBody string
	}{
		{
			"Success",
			"123",
			200,
			`{"id":"123","description":"Wedding","type":"","status":"NEW","created_at":null,"updated_at":null,"event_date":null,"event_location":"","name":"","phone":"","email":""}`,
		},
		{
			"Not found",
			"456",
			404,
			`{"message":"Not found","status":404}`,
		},
		{
			"Error",
			"789",
			500,
			`{"message":"Internal server error","status":500}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler.
			req, err := http.NewRequest("GET", fmt.Sprintf("/events/%s/reservations", tt.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/events/{id}/reservations", e.GetEventByID)

			// Change to Gorilla Mux router to pass variables
			router.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}
		})
	}
}

func Test_eventController_GetReservations(t *testing.T) {
	events := &mocks.Events{}
	e := &eventController{
		events,
	}

	arrival := time.Date(2020, 1, 1, 4, 0, 0, 0, time.UTC)
	departure := time.Date(2020, 1, 3, 18, 0, 0, 0, time.UTC)
	reservations := []model.Reservation{
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
	}

	events.On("GetByID", "123").Return(model.Event{}, nil)
	events.On("GetByID", "456").Return(model.Event{}, errors.New("failed"))
	events.On("GetByID", "789").Return(model.Event{}, nil)
	events.On("GetReservations", "123").Return(reservations, nil)
	events.On("GetReservations", "789").Return(nil, errors.New("failed"))

	tests := []struct {
		name         string
		id           string
		expectedCode int
		expectedBody string
	}{
		{
			"Success",
			"123",
			201,
			`{"id":"123","description":"Wedding","type":"","status":"NEW","created_at":null,"updated_at":null,"event_date":null,"event_location":"","name":"","phone":"","email":""}`,
		},
		{
			"Error",
			"789",
			500,
			`{"message":"Internal server error","status":500}`,
		},
		{
			"Not found",
			"456",
			404,
			`{"message":"Not found","status":404}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request to pass to our handler.
			req, err := http.NewRequest("GET", fmt.Sprintf("/events/%s", tt.id), nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/events/{id}", e.GetReservations)

			// Change to Gorilla Mux router to pass variables
			router.ServeHTTP(rr, req)

			// Check the content type is what we expect.
			expected := "application/json; charset=UTF-8"
			m := rr.Header()
			if contentType := m.Get("Content-Type"); contentType != expected {
				t.Errorf(
					"handler returned wrong content type: got %v want %v",
					contentType,
					expected,
				)
			}

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf(
					"handler returned wrong status code: got %v want %v",
					status,
					tt.expectedCode,
				)
			}

			// Check the response body is what we expect.
			if rr.Body.String() != tt.expectedBody {
				t.Errorf(
					"handler returned unexpected body: got %v want %v",
					rr.Body.String(),
					tt.expectedBody,
				)
			}
		})
	}
}
