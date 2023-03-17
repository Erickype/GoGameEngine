package Events

import (
	"reflect"
)

type EventDispatcher struct {
	event *IEvent
}

func (d *EventDispatcher) Dispatch(fn interface{}) bool {
	argumentType := reflect.TypeOf(fn).In(0)

	if reflect.TypeOf(*d.event).AssignableTo(argumentType) {
		(*d.event).SetHandled(reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(*d.event)})[0].Bool())
		return true
	}
	return false
}

func CreateDispatcher(event *IEvent) *EventDispatcher {
	dispatcher := &EventDispatcher{event: event}
	return dispatcher
}
