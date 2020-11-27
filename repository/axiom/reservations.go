package axiom

import (
	"context"
	"log"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// collection keeps the table name
const collection = "reservations"

type reservations struct {
	db    string
	mongo *mongo.Client
}

type mongoReservation struct {
	ID        string     `bson:"_id"`
	Status    string     `bson:"status"`
	Plan      string     `bson:"plan"` // how can I predefine values here?
	Adults    int        `bson:"adults"`
	Minors    int        `bson:"minors"`
	AdultFee  float64    `bson:"adultFee"`
	MinorFee  float64    `bson:"minorFee"`
	Arrival   *time.Time `bson:"arrival"`
	Departure *time.Time `bson:"departure"`
	CreatedAt *time.Time `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`

	// Event information
	EventID string `bson:"event_id"`

	// Customer information
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
}

// NewReservationRepo returns the ReservationRepository implementation
func NewReservationRepo(db string, m *mongo.Client) repository.ReservationRepository {
	return &reservations{
		db,
		m,
	}
}

func (r *reservations) Create(eventID string, reservation model.Reservation) (model.Reservation, error) {
	collection := r.mongo.Database(r.db).Collection(collection)

	insert := &mongoReservation{
		ID:        reservation.ID,
		Status:    reservation.Status,
		Plan:      reservation.Plan,
		Adults:    reservation.Adults,
		Minors:    reservation.Minors,
		AdultFee:  reservation.AdultFee,
		MinorFee:  reservation.MinorFee,
		Arrival:   reservation.Arrival,
		Departure: reservation.Departure,
		CreatedAt: reservation.CreatedAt,
		UpdatedAt: reservation.UpdatedAt,
		EventID:   eventID,
		Name:      reservation.Name,
		Phone:     reservation.Phone,
		Email:     reservation.Email,
	}

	_, err := collection.InsertOne(context.Background(), insert)
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

func (r *reservations) GetByEventID(id string) ([]model.Reservation, error) {
	collection := r.mongo.Database(r.db).Collection(collection)

	filter := bson.M{"event_id": id}

	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	var response []model.Reservation

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {
		// create a value into which the single document can be decoded
		var item mongoReservation
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}

		response = append(response, item.toModel())
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.Background())

	return response, nil
}

func (r mongoReservation) toModel() model.Reservation {
	return model.Reservation{
		ID:        r.ID,
		Status:    r.Status,
		Plan:      r.Plan,
		Adults:    r.Adults,
		Minors:    r.Minors,
		AdultFee:  r.AdultFee,
		MinorFee:  r.MinorFee,
		Arrival:   r.Arrival,
		Departure: r.Departure,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Name:      r.Name,
		Phone:     r.Phone,
		Email:     r.Email,
	}
}
