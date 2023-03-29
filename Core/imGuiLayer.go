package Core

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	clearColor = [3]float32{0.1, 0.5, 0.8}
)

type ImGuiLayer struct {
	Layer
}

func (l *ImGuiLayer) Begin() {
	(*ApplicationInstance.GetPlatform()).NewFrame()
	imgui.NewFrame()
}

func (l *ImGuiLayer) OnImGuiRender() {
	imgui.ShowDemoWindow()
}

func (l *ImGuiLayer) End() {
	imgui.Render()
	(*ApplicationInstance.GetRenderer()).PreRender(clearColor)
	(*ApplicationInstance.GetRenderer()).Render(
		(*ApplicationInstance.GetPlatform()).DisplaySize(),
		(*ApplicationInstance.GetPlatform()).FramebufferSize(),
		imgui.GetDrawData())
}

func (l *ImGuiLayer) OnEvent(event *Events.IEvent) {
	dispatcher := Events.CreateDispatcher(event)

	dispatcher.Dispatch(l.OnMouseButtonPressedEvent)
	dispatcher.Dispatch(l.OnMouseButtonReleasedEvent)
	dispatcher.Dispatch(l.OnMouseMovedEvent)
	dispatcher.Dispatch(l.OnMouseScrolledEvent)
	dispatcher.Dispatch(l.OnKeyPressedEvent)
	dispatcher.Dispatch(l.OnKeyReleasedEvent)
	dispatcher.Dispatch(l.OnKeyTypedEvent)
	dispatcher.Dispatch(l.OnWindowResizeEvent)
}

func (l *ImGuiLayer) OnMouseButtonPressedEvent(event *Events.MouseButtonPressedEvent) bool {
	io := imgui.GetIO()
	io.SetMouseButtonDown(event.GetMouseButton(), true)
	return false
}

func (l *ImGuiLayer) OnMouseButtonReleasedEvent(event *Events.MouseButtonReleaseEvent) bool {
	io := imgui.GetIO()
	io.SetMouseButtonDown(event.GetMouseButton(), false)
	return false
}

func (l *ImGuiLayer) OnMouseMovedEvent(event *Events.MouseMovedEvent) bool {
	io := imgui.GetIO()
	io.SetMousePos(imgui.ImVec2{
		X: float32(event.GetX()),
		Y: float32(event.GetY()),
	})
	return false
}

func (l *ImGuiLayer) OnMouseScrolledEvent(event *Events.MouseScrolledEvent) bool {
	io := imgui.GetIO()
	io.AddMouseWheelDelta(float32(event.GetXOffset()), float32(event.GetYOffset()))
	return false
}

func (l *ImGuiLayer) OnKeyPressedEvent(event *Events.KeyPressedEvent) bool {

	io := imgui.GetIO()
	iKeyEvent := Events.IKeyEvent(event)
	io.AddKeyEvent(keyMapped(&iKeyEvent), true)
	return false
}

func (l *ImGuiLayer) OnKeyReleasedEvent(event *Events.KeyReleasedEvent) bool {
	io := imgui.GetIO()
	iKeyEvent := Events.IKeyEvent(event)
	io.AddKeyEvent(keyMapped(&iKeyEvent), false)
	return false
}

func keyMapped(event *Events.IKeyEvent) imgui.ImGuiKey {
	platform := *ApplicationInstance.GetPlatform()

	imKey := imgui.ImGuiKey((*event).GetKeyCode())

	if mapped, ok := platform.GetKeyMap()[glfw.Key((*event).GetKeyCode())]; ok {
		imKey = mapped
	}
	return imKey
}

func (l *ImGuiLayer) OnKeyTypedEvent(event *Events.KeyTypedEvent) bool {
	io := imgui.GetIO()
	io.AddInputCharacter(uint32(event.GetKeyCode()))
	return false
}

func (l *ImGuiLayer) OnWindowResizeEvent(event *Events.WindowResizeEvent) bool {
	io := imgui.GetIO()
	io.SetDisplaySize(imgui.ImVec2{
		X: float32(event.GetWidth()),
		Y: float32(event.GetHeight()),
	})
	io.GetDisplayFramebufferScale()
	io.SetDisplayFramebufferScale(imgui.ImVec2{X: 1, Y: 1})
	gl.Viewport(0, 0, int32(event.GetWidth()), int32(event.GetHeight()))
	return false
}
func NewImGui() *ImGuiLayer {
	layer := ImGuiLayer{}
	layer.Construct("ImGui")
	return &layer
}
