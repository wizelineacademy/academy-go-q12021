package domain

// AggregateRoot is a group of entities and value objects which can be used as unit of work inside a domain-driven system.
type AggregateRoot interface {
	// PullEvents returns all the domain events that had happened inside the current aggregate root
	PullEvents() []Event
}
