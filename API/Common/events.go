package Common

import (
	Events2 "github.com/Erickype/GoGameEngine/API/Events"
)

var EventFactory *Events2.EventFactory
var EventDispatcher *Events2.EventDispatcher

func initEventsFactory() {
	EventFactory = Events2.GetEventFactoryInstance()
}

func initEventDispatcher() {
	EventDispatcher = Events2.GetEventDispatcherInstance()
}
