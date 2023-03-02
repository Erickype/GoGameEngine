package Common

import "github.com/Erickype/GoGameEngine/Events"

var EventFactory *Events.EventFactory
var EventDispatcher *Events.EventDispatcher

func initEventsFactory() {
	EventFactory = Events.GetEventFactoryInstance()
}

func initEventDispatcher() {
	EventDispatcher = Events.GetEventDispatcherInstance()
}
