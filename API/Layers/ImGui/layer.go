package Layers

import (
	"github.com/AllenDang/cimgui-go"
	"github.com/Erickype/GoGameEngine/API/Common"
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
	Core.ApplicationInstance.GetRenderer().Render(
		Core.ApplicationInstance.GetPlatform().DisplaySize(),
		Core.ApplicationInstance.GetPlatform().FramebufferSize(),
		imgui.CurrentDrawData())
}

func (l *Layer) OnEvent(event *Events.IEvent) {
	dispatcher := Events.CreateDispatcher(event)

	dispatcher.Dispatch(l.OnMouseButtonPressedEvent)
	dispatcher.Dispatch(l.OnMouseButtonReleasedEvent)
	dispatcher.Dispatch(l.OnMouseMovedEvent)
	dispatcher.Dispatch(l.OnMouseScrolledEvent)
	dispatcher.Dispatch(l.OnKeyPressedEvent)
	dispatcher.Dispatch(l.OnKeyReleasedEvent)
	dispatcher.Dispatch(l.OnWindowResizeEvent)
}

func (l *Layer) OnMouseButtonPressedEvent(event *Events.MouseButtonPressedEvent) bool {
	io := imgui.CurrentIO()
	io.SetMouseButtonDown(event.GetMouseButton(), true)

	return false
}

func (l *Layer) OnMouseButtonReleasedEvent(event *Events.MouseButtonReleaseEvent) bool {
	io := imgui.CurrentIO()
	io.SetMouseButtonDown(event.GetMouseButton(), false)
	return false
}

func (l *Layer) OnMouseMovedEvent(event *Events.MouseMovedEvent) bool {
	io := imgui.CurrentIO()
	io.SetMousePos(imgui.Vec2{
		X: float32(event.GetX()),
		Y: float32(event.GetY()),
	})
	return false
}

func (l *Layer) OnMouseScrolledEvent(event *Events.MouseScrolledEvent) bool {
	io := imgui.CurrentIO()
	io.SetMouseWheel(float32(event.GetYOffset()))
	io.SetMouseWheelH(float32(event.GetXOffset()))
	return false
}

func (l *Layer) OnKeyPressedEvent(event *Events.KeyPressedEvent) bool {
	Common.CoreLogger.Debug("OnKeyPressedEvent: ", event)
	return false
}

func (l *Layer) OnKeyReleasedEvent(event *Events.KeyReleasedEvent) bool {
	Common.CoreLogger.Debug("OnKeyReleasedEvent: ", event)
	return false
}
func (l *Layer) OnWindowResizeEvent(event *Events.WindowResizeEvent) bool {
	Common.CoreLogger.Debug("OnWindowResizeEvent: ", event)
	return false
}
func NewImGui() *Layer {
	layer := Layer{}
	layer.Construct("ImGui")
	return &layer
}
