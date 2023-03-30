package Core

import (
	"github.com/Erickype/GoGameEngine/API/Events"
)

type ILayer interface {
	OnAttach()
	OnDetach()
	OnUpdate()
	OnEvent(event *Events.IEvent)
	OnImGuiRender()
	GetName() string
	Construct(debugName string)
}
type Layer struct {
	debugName string
}

func (l *Layer) OnAttach() {}

func (l *Layer) OnDetach() {}

func (l *Layer) OnUpdate() {}

func (l *Layer) OnEvent(_ *Events.IEvent) {}

func (l *Layer) OnImGuiRender() {}

func (l *Layer) GetName() string {
	return l.debugName
}

func (l *Layer) Construct(debugName string) {
	l.debugName = debugName
}
