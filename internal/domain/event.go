package domain

// Event is a representation of aggregate's mutations inside the system in order to propagate side-effects to the same or external system(s).
type Event interface {
	// Kind returns the specific type of the event using the Async API naming specification
	Kind() string
	// AggregateID returns the unique identifier of the aggregate whos state has changed
	AggregateID() string
}
