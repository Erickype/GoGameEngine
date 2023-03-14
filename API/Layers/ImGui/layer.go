package Layers

import (
	"fmt"
	"github.com/AllenDang/cimgui-go"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/Core"
	"unsafe"
)

const (
	millisPerSecond = 1000
)

var (
	showDemoWindow    = false
	showGoDemoWindow  = false
	clearColor        = [3]float32{0.0, 0.0, 0.0}
	f                 = float32(0)
	counter           = 0
	showAnotherWindow = false
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

	// 1. Show a simple window.
	// Tip: if we don't call imGui.Begin()/imGui.End() the widgets automatically appears in a window called "Debug".
	{
		imgui.Text("ภาษาไทย测试조선말")                   // To display these, you'll need to register a compatible font
		imgui.Text("Hello, world!")                  // Display some text
		imgui.SliderFloat("float", &f, 0.0, 1.0)     // Edit 1 float using a slider from 0.0f to 1.0f
		imgui.ColorEdit3("clear color", &clearColor) // Edit 3 floats representing a color

		imgui.Checkbox("Demo Window", &showDemoWindow) // Edit booleans storing our window open/close state
		imgui.Checkbox("Go Demo Window", &showGoDemoWindow)
		imgui.Checkbox("Another Window", &showAnotherWindow)

		if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
			counter++
		}
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("counter = %d", counter))

		imgui.Text(fmt.Sprintf("Application average %.3f ms/frame (%.1f FPS)",
			millisPerSecond/imgui.CurrentIO().Framerate(), imgui.CurrentIO().Framerate()))
	}

	// 2. Show another simple window. In most cases you will use an explicit Beginning/End pair to name your windows.
	if showAnotherWindow {
		// Pass a pointer to our bool variable (the window will have a closing button that will clear the bool when clicked)
		imgui.BeginV("Another window", &showAnotherWindow, 0)
		imgui.Text("Hello from another window!")
		if imgui.Button("Close Me") {
			showAnotherWindow = false
		}
		imgui.End()
	}

	// 3. Show the ImGui demo window. Most of the sample code is in imGui.ShowDemoWindow().
	// Read its code to learn more about Dear ImGui!
	if showDemoWindow {
		// Normally user code doesn't need/want to call this because positions are saved in .ini file anyway.
		// Here we just want to make the demo initial state a bit more friendly!
		const demoX = 650
		const demoY = 20
		imgui.SetNextWindowPosV(imgui.Vec2{X: demoX, Y: demoY}, imgui.CondFirstUseEver, imgui.Vec2{})

		imgui.ShowDemoWindow()
	}

	// Rendering
	imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

	Core.ApplicationInstance.GetRenderer().PreRender(clearColor)
	// At this point, the application could perform its own rendering...
	// app.RenderScene()

	Core.ApplicationInstance.GetRenderer().Render(Core.ApplicationInstance.GetPlatform().DisplaySize(), Core.ApplicationInstance.GetPlatform().FramebufferSize(), imgui.CurrentDrawData())
}

func (l *Layer) OnEvent(_ *Events.IEvent) {

}

func NewImGui() *Layer {
	layer := Layer{}
	layer.Construct("ImGui")
	return &layer
}
