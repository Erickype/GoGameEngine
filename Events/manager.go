package Events

type IEventManager interface {
	GetDispatcher() *EventDispatcher
	CreateEvent(eventType EventType) *IEvent
}

type EventManager struct {
	eventDispatcher *EventDispatcher
}

func (e *EventManager) GetDispatcher() *EventDispatcher {
	return e.eventDispatcher.GetInstance()
}

func (e *EventManager) CreateEvent(eventType EventType) IEvent {

	event := Event{
		handled:       false,
		eventCategory: 0,
		eventType:     0,
	}

	switch eventType {
	case MouseMoved:
		return &MouseMovedEvent{
			Event:  &event,
			mouseX: 0,
			mouseY: 0,
		}
	default:
		return nil
	}
}
