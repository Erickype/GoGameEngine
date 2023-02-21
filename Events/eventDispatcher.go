package Events

var instance *EventDispatcher

func init() {
	instance = &EventDispatcher{}
}

type IEventDispatcher interface {
	GetInstance() *EventDispatcher
	Dispatch(event IEvent) bool
}

type EventDispatcher struct{}

func (d *EventDispatcher) GetInstance() *EventDispatcher {
	return instance
}

func (d *EventDispatcher) Dispatch(event IEvent) bool {
	response := false
	switch event.(type) {
	case *MouseMovedEvent:
		concreteEvent := event.(*MouseMovedEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *MouseScrolledEvent:
		concreteEvent := event.(*MouseScrolledEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	}
	return response
}
