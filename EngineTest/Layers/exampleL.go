package Layers

import (
	"fmt"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Input"
	"github.com/Erickype/GoGameEngine/API/Log"
	core "github.com/Erickype/GoGameEngine/Core"
)

type ExampleLayer struct {
	core.Layer
}

func (e *ExampleLayer) OnUpdate() {
	if (*Input.GetInputInstance()).IsKeyPressed((int)(Input.KeySpace), (*(*core.ApplicationInstance).GetPlatform()).GetWindowPtr()) {
		Log.GetClientInstance().Warn("Key pressed (poll)")
	}
}

func (e *ExampleLayer) OnEvent(event *Events.IEvent) {
	if keyPressedEvent, ok := (*event).(*Events.KeyPressedEvent); ok {
		character := fmt.Sprintf("%c", rune(keyPressedEvent.GetKeyCode()))
		Log.GetClientInstance().Info(e.GetName(), "layer event (event):", character)

		if keyPressedEvent.GetKeyCode() == (int)(Input.KeySpace) {
			Log.GetClientInstance().Trace("Key space pressed (event)")
		}
	}
}
