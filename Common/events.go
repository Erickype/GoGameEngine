package Common

import "github.com/Erickype/GoGameEngine/Events"

var EventFactory *Events.EventFactory
var _ *Events.EventDispatcher

func initEventsFactory() {
	EventFactory = Events.GetEventFactoryInstance()
}

func initEventDispatcher() {
	_ = Events.GetEventDispatcherInstance()
}
