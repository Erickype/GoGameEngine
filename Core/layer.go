package Core

import "github.com/Erickype/GoGameEngine/Events"

type iLayer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnEvent(event *Events.IEvent)
	GetName() string
	Construct(debugName string)
}
