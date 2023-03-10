package ImGui

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/Core"
)

type Layer struct {
	Core.Layer
}

func (l *Layer) OnAttach() {
}

func (l *Layer) OnDetach() {

}

func (l *Layer) OnUpdate() {

}

func (l *Layer) OnEvent(_ *Events.IEvent) {

}

func New() *Layer {
	layer := Layer{}
	layer.Construct("ImGui")
	return &layer
}
