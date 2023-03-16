package Events

type IEventManager interface {
	CreateEvent(eventType EventType) IEvent
}

type EventManager struct{}

func (e *EventManager) CreateEvent(eventType EventType) IEvent {
	return factoryInstance.CreateEvent(eventType)
}
