package Layers

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Internal/renderers/gl/v3.2-core/gl"
	"github.com/Erickype/GoGameEngine/Core"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

var (
	clearColor = [3]float32{0.1, 0.5, 0.8}
)

type Layer struct {
	Core.Layer
}

func (l *Layer) OnUpdate() {
	(*Core.ApplicationInstance.GetPlatform()).NewFrame()
	imgui.NewFrame()

	demo := true
	imgui.ShowDemoWindow(&demo)

	imgui.Render()
	(*Core.ApplicationInstance.GetRenderer()).PreRender(clearColor)
	(*Core.ApplicationInstance.GetRenderer()).Render(
		(*Core.ApplicationInstance.GetPlatform()).DisplaySize(),
		(*Core.ApplicationInstance.GetPlatform()).FramebufferSize(),
		imgui.RenderedDrawData())
}

func (l *Layer) OnEvent(event *Events.IEvent) {
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
	io.SetMousePosition(imgui.Vec2{
		X: float32(event.GetX()),
		Y: float32(event.GetY()),
	})
	return false
}

func (l *Layer) OnMouseScrolledEvent(event *Events.MouseScrolledEvent) bool {
	io := imgui.CurrentIO()
	io.AddMouseWheelDelta(float32(event.GetXOffset()), float32(event.GetYOffset()))
	return false
}

func (l *Layer) OnKeyPressedEvent(event *Events.KeyPressedEvent) bool {
	io := imgui.CurrentIO()
	io.KeyPress(event.GetKeyCode())
	keyModifiers()
	return false
}

func (l *Layer) OnKeyReleasedEvent(event *Events.KeyReleasedEvent) bool {
	io := imgui.CurrentIO()
	io.KeyRelease(event.GetKeyCode())
	keyModifiers()
	return false
}

func keyModifiers() {
	io := imgui.CurrentIO()
	io.KeyCtrl(int(glfw.KeyLeftControl), int(glfw.KeyRightControl))
	io.KeyShift(int(glfw.KeyLeftShift), int(glfw.KeyRightShift))
	io.KeyAlt(int(glfw.KeyLeftAlt), int(glfw.KeyRightAlt))
	io.KeySuper(int(glfw.KeyLeftSuper), int(glfw.KeyRightSuper))
}

func (l *Layer) OnKeyTypedEvent(event *Events.KeyTypedEvent) bool {
	io := imgui.CurrentIO()
	io.AddInputCharacters(string(rune(event.GetKeyCode())))
	return false
}

func (l *Layer) OnWindowResizeEvent(event *Events.WindowResizeEvent) bool {
	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{
		X: float32(event.GetWidth()),
		Y: float32(event.GetHeight()),
	})
	io.DisplayFrameBufferScale()
	io.SetDisplayFrameBufferScale(imgui.Vec2{X: 1, Y: 1})
	gl.Viewport(0, 0, int32(event.GetWidth()), int32(event.GetHeight()))
	return false
}
func NewImGui() *Layer {
	layer := Layer{}
	layer.Construct("ImGui")
	return &layer
}
