package Events

type IEventManager interface {
	Dispatch() bool
	CreateEvent(eventType EventType) IEvent
}

type EventManager struct{}

func (e *EventManager) Dispatch(event IEvent) bool {
	return dispatcherInstance.Dispatch(event)
}

func (e *EventManager) CreateEvent(eventType EventType) IEvent {
	return factoryInstance.CreateEvent(eventType)
}
