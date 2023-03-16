package Layers

import (
	"github.com/AllenDang/cimgui-go"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/Core"
	"unsafe"
)

var (
	clearColor = [3]float32{0.0, 0.0, 0.0}
)

type Layer struct {
	Core.Layer
}

func (l *Layer) OnAttach() {
	var data int32 = 42
	imgui.CurrentIO().SetClipboardUserData(unsafe.Pointer(&data))
}

func (l *Layer) OnUpdate() {
	Core.ApplicationInstance.GetPlatform().NewFrame()
	imgui.NewFrame()

	imgui.ShowDemoWindow()

	imgui.Render()
	Core.ApplicationInstance.GetRenderer().PreRender(clearColor)
	Core.ApplicationInstance.GetRenderer().Render(Core.ApplicationInstance.GetPlatform().DisplaySize(), Core.ApplicationInstance.GetPlatform().FramebufferSize(), imgui.CurrentDrawData())
}

func (l *Layer) OnEvent(_ *Events.IEvent) {

}

func NewImGui() *Layer {
	layer := Layer{}
	layer.Construct("ImGui")
	return &layer
}
