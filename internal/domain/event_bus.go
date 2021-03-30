package domain

// EventBus used to push events to the whole system so anybody can subscribe to many specific events in order to process events
// asynchronously.
type EventBus interface {
	// PublishEvents pushes the given events into the system
	PublishEvents(events ...Event) error
}
