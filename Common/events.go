package Common

import "github.com/Erickype/GoGameEngine/Events"

var EventsFactory *Events.EventFactory

func initEventsFactory() {
	EventsFactory = Events.NewEventFactory()
}
