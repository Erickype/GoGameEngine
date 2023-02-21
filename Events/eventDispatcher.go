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
	case *MouseButtonPressedEvent:
		concreteEvent := event.(*MouseButtonPressedEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *MouseButtonReleaseEvent:
		concreteEvent := event.(*MouseButtonReleaseEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *KeyPressedEvent:
		concreteEvent := event.(*KeyPressedEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *KeyReleasedEvent:
		concreteEvent := event.(*KeyReleasedEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *WindowResizeEvent:
		concreteEvent := event.(*WindowResizeEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	case *WindowCloseEvent:
		concreteEvent := event.(*WindowCloseEvent)
		concreteEvent.handled = true
		response = concreteEvent.handled
	}
	return response
}
