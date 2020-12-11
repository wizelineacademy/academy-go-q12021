package axiom

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/javiertlopez/golang-bootcamp-2020/model"
	"github.com/javiertlopez/golang-bootcamp-2020/repository"

	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection keeps the collection name
const Collection = "events"

type events struct {
	db    string
	mongo *mongo.Client
}

type mongoEvent struct {
	ID            string     `bson:"_id"`
	Description   string     `bson:"description"`
	Type          string     `bson:"type"` // how can I predefine values here?
	Status        string     `bson:"status"`
	CreatedAt     *time.Time `bson:"createdAt"`
	UpdatedAt     *time.Time `bson:"updatedAt"`
	EventDate     *time.Time `bson:"eventDate"`     // should I drop 'event_'?
	EventLocation string     `bson:"eventLocation"` // should I drop 'event_'?

	// Customer information
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
	Email string `bson:"email"`
}

// NewEventsRepo returns the EventRepository implementation
func NewEventsRepo(db string, m *mongo.Client) repository.EventRepository {
	return &events{
		db,
		m,
	}
}

func (e *events) Create(event model.Event) (model.Event, error) {
	collection := e.mongo.Database(e.db).Collection(Collection)

	// Step 0. Let's create a UUID
	uuid := guuid.New().String()
	event.ID = uuid

	// Step 0.1. Now!
	now := time.Now()
	event.CreatedAt = &now
	event.UpdatedAt = &now

	insert := &mongoEvent{
		ID:            event.ID,
		Description:   event.Description,
		Type:          event.Type,
		Status:        event.Status,
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
		EventDate:     event.EventDate,
		EventLocation: event.EventLocation,
		Name:          event.Name,
		Phone:         event.Phone,
		Email:         event.Email,
	}

	_, err := collection.InsertOne(context.Background(), insert)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (e *events) GetByID(id string) (model.Event, error) {
	var response mongoEvent

	collection := e.mongo.Database(e.db).Collection(Collection)

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.Background(), filter).Decode(&response)

	if err != nil {
		return model.Event{}, err
	}

	return response.toModel(), nil
}

func (e *events) GetAll() ([]model.Event, error) {
	collection := e.mongo.Database(e.db).Collection(collection)

	// bson.D{{}} gets all
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	var response []model.Event

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.Background()) {
		// create a value into which the single document can be decoded
		var item mongoEvent
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

func (e *events) Update(event model.Event) (model.Event, error) {
	return event, fmt.Errorf("not implemented")
}

func (e *events) Delete(id string) error {
	return fmt.Errorf("not implemented")
}

func (e mongoEvent) toModel() model.Event {
	return model.Event{
		ID:            e.ID,
		Description:   e.Description,
		Type:          e.Type,
		Status:        e.Status,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		EventDate:     e.EventDate,
		EventLocation: e.EventLocation,

		// Customer information
		Name:  e.Name,
		Phone: e.Phone,
		Email: e.Email,
	}
}
