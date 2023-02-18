package Events

var instance *EventDispatcher

func init() {
	instance = &EventDispatcher{}
}

type IEventDispatcher interface {
	GetInstance() *EventDispatcher
}

type EventDispatcher struct {
}

func (d *EventDispatcher) GetInstance() *EventDispatcher {
	return instance
}
