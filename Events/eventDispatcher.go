package Events

var dispatcherInstance *EventDispatcher

func init() {
	dispatcherInstance = &EventDispatcher{}
}

type IEventDispatcher interface {
	Dispatch(event IEvent) bool
}

type EventDispatcher struct{}

func GetEventDispatcherInstance() *EventDispatcher {
	return dispatcherInstance
}

func (d *EventDispatcher) Dispatch(event IEvent) bool {
	response := false
	switch event := event.(type) {
	case *MouseMovedEvent:
		event.handled = true
		response = event.handled
	case *MouseScrolledEvent:
		event.handled = true
		response = event.handled
	case *MouseButtonPressedEvent:
		event.handled = true
		response = event.handled
	case *MouseButtonReleaseEvent:
		event.handled = true
		response = event.handled
	case *KeyPressedEvent:
		event.handled = true
		response = event.handled
	case *KeyReleasedEvent:
		event.handled = true
		response = event.handled
	case *WindowResizeEvent:
		event.handled = true
		response = event.handled
	case *WindowCloseEvent:
		event.handled = true
		response = event.handled
	}
	return response
}
