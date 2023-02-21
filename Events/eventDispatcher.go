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
		mouseMovedEvent := event.(*MouseMovedEvent)
		mouseMovedEvent.handled = true
		response = mouseMovedEvent.handled
	}
	return response
}
