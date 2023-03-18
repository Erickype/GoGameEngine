package Events

var factoryInstance *EventFactory

func init() {
	factoryInstance = NewEventFactory()
}

type EventFactory struct {
	events map[EventType]IEvent
}

func NewEventFactory() *EventFactory {
	factory := &EventFactory{
		events: make(map[EventType]IEvent),
	}

	event := Event{
		handled:       false,
		eventCategory: 0,
		eventType:     0,
	}

	factory.events[MouseMoved] = &MouseMovedEvent{Event: &event}
	factory.events[MouseScrolled] = &MouseScrolledEvent{Event: &event}

	mouseEvent := MouseButtonEvent{
		Event:  &event,
		Button: 0,
	}
	factory.events[MouseButtonPressed] = &MouseButtonPressedEvent{MouseButtonEvent: &mouseEvent}
	factory.events[MouseButtonReleased] = &MouseButtonReleaseEvent{MouseButtonEvent: &mouseEvent}

	keyEvent := keyEvent{
		Event:   &event,
		KeyCode: 0,
	}
	factory.events[KeyPressed] = &KeyPressedEvent{keyEvent: &keyEvent}
	factory.events[KeyReleased] = &KeyReleasedEvent{keyEvent: &keyEvent}
	factory.events[KeyTyped] = &KeyTypedEvent{keyEvent: &keyEvent}

	factory.events[WindowResize] = &WindowResizeEvent{Event: &event}
	factory.events[WindowClose] = &WindowCloseEvent{Event: &event}

	return factory
}

func (factory *EventFactory) CreateEvent(eventType EventType) IEvent {
	if event, ok := factory.events[eventType]; ok {
		event.Init()
		return event
	}

	return nil
}

func GetEventFactoryInstance() *EventFactory {
	return factoryInstance
}
