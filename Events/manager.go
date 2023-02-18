package Events

type IEventManager interface {
	GetDispatcher() *EventDispatcher
	CreateEvent(eventType Type) *MouseMovedEvent
}

type EventManager struct {
	eventDispatcher *EventDispatcher
}

func (e *EventManager) GetDispatcher() *EventDispatcher {
	return e.eventDispatcher.GetInstance()
}

func (e *EventManager) CreateEvent(eventType Type) *MouseMovedEvent {

	event := Event{
		dispatcher:    e.GetDispatcher(),
		handled:       false,
		eventCategory: 0,
		eventType:     0,
	}
	var newEvent MouseMovedEvent

	switch eventType {
	case MouseMoved:
		newEvent = MouseMovedEvent{
			Event:  &event,
			mouseX: 0,
			mouseY: 0,
		}
	}
	return &newEvent
}
