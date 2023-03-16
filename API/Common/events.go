package Common

import (
	Events2 "github.com/Erickype/GoGameEngine/API/Events"
)

var EventFactory *Events2.EventFactory

func initEventsFactory() {
	EventFactory = Events2.GetEventFactoryInstance()
}
