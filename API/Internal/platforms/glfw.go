package platforms

import (
	"fmt"
	"github.com/AllenDang/cimgui-go"
	"github.com/go-gl/glfw/v3.3/glfw"
	"math"
	"runtime"
)

type GLFWClientAPI string

// This is a list of GLFWClientAPI constants.
const (
	GLFWClientAPIOpenGL2 GLFWClientAPI = "OpenGL2"
	GLFWClientAPIOpenGL3 GLFWClientAPI = "OpenGL3"
)

// GLFW implements a platform based on github.com/go-gl/glfw (v3.3).
type GLFW struct {
	imGuiIO          imgui.IO
	window           *glfw.Window
	time             float64
	mouseJustPressed [3]bool
}

func (g *GLFW) GetWindow() *glfw.Window {
	return g.window
}

func NewGLFW(io imgui.IO, clientAPI GLFWClientAPI, width int, height int, title string) (*GLFW, error) {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize glfw: %w", err)
	}

	switch clientAPI {
	case GLFWClientAPIOpenGL2:
		glfw.WindowHint(glfw.ContextVersionMajor, 2)
		glfw.WindowHint(glfw.ContextVersionMinor, 1)
	case GLFWClientAPIOpenGL3:
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 2)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	default:
		glfw.Terminate()
		return nil, err
	}

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		glfw.Terminate()
		return nil, fmt.Errorf("failed to create window: %w", err)
	}
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	platform := &GLFW{
		imGuiIO: io,
		window:  window,
	}

	platform.setKeyMapping()
	platform.installCallbacks()

	return platform, nil
}

func (g *GLFW) Dispose() {
	g.window.Destroy()
	glfw.Terminate()
}

// ShouldStop returns true if the window is to be closed.
func (g *GLFW) ShouldStop() bool {
	return g.window.ShouldClose()
}

// ProcessEvents handles all pending window events.
func (g *GLFW) ProcessEvents() {
	glfw.PollEvents()
}

// DisplaySize returns the dimension of the display.
func (g *GLFW) DisplaySize() [2]float32 {
	w, h := g.window.GetSize()
	return [2]float32{float32(w), float32(h)}
}

// FramebufferSize returns the dimension of the framebuffer.
func (g *GLFW) FramebufferSize() [2]float32 {
	w, h := g.window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

// NewFrame marks the beginning of a render pass. It forwards all current state to imGui IO.
func (g *GLFW) NewFrame() {
	// Setup display size (every frame to accommodate for window resizing)
	displaySize := g.DisplaySize()
	g.imGuiIO.SetDisplaySize(imgui.Vec2{X: displaySize[0], Y: displaySize[1]})

	// Setup time step
	currentTime := glfw.GetTime()
	if g.time > 0 {
		g.imGuiIO.SetDeltaTime(float32(currentTime - g.time))
	}
	g.time = currentTime

	// Setup inputs
	if g.window.GetAttrib(glfw.Focused) != 0 {
		x, y := g.window.GetCursorPos()
		g.imGuiIO.SetMousePos(imgui.Vec2{X: float32(x), Y: float32(y)})
	} else {
		g.imGuiIO.SetMousePos(imgui.Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(g.mouseJustPressed); i++ {
		down := g.mouseJustPressed[i] || (g.window.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		g.imGuiIO.SetMouseButtonDown(i, down)
		g.mouseJustPressed[i] = false
	}
}

// PostRender performs a buffer swap.
func (g *GLFW) PostRender() {
	g.window.SwapBuffers()
}

func (g *GLFW) setKeyMapping() {
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyTab, int32(glfw.KeyTab), int32(glfw.GetKeyScancode(glfw.KeyTab)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyLeftArrow, int32(glfw.KeyLeft), int32(glfw.GetKeyScancode(glfw.KeyLeft)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyRightArrow, int32(glfw.KeyRight), int32(glfw.GetKeyScancode(glfw.KeyRight)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyUpArrow, int32(glfw.KeyUp), int32(glfw.GetKeyScancode(glfw.KeyUp)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyDownArrow, int32(glfw.KeyDown), int32(glfw.GetKeyScancode(glfw.KeyDown)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyPageDown, int32(glfw.KeyPageDown), int32(glfw.GetKeyScancode(glfw.KeyPageDown)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyPageUp, int32(glfw.KeyPageUp), int32(glfw.GetKeyScancode(glfw.KeyPageUp)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyHome, int32(glfw.KeyHome), int32(glfw.GetKeyScancode(glfw.KeyHome)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyEnd, int32(glfw.KeyEnd), int32(glfw.GetKeyScancode(glfw.KeyEnd)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyInsert, int32(glfw.KeyInsert), int32(glfw.GetKeyScancode(glfw.KeyInsert)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyDelete, int32(glfw.KeyDelete), int32(glfw.GetKeyScancode(glfw.KeyDelete)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyBackspace, int32(glfw.KeyBackspace), int32(glfw.GetKeyScancode(glfw.KeyBackspace)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeySpace, int32(glfw.KeySpace), int32(glfw.GetKeyScancode(glfw.KeySpace)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyEnter, int32(glfw.KeyEnter), int32(glfw.GetKeyScancode(glfw.KeyEnter)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyEscape, int32(glfw.KeyEscape), int32(glfw.GetKeyScancode(glfw.KeyEscape)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyA, int32(glfw.KeyA), int32(glfw.GetKeyScancode(glfw.KeyA)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyC, int32(glfw.KeyC), int32(glfw.GetKeyScancode(glfw.KeyC)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyV, int32(glfw.KeyV), int32(glfw.GetKeyScancode(glfw.KeyV)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyX, int32(glfw.KeyX), int32(glfw.GetKeyScancode(glfw.KeyX)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyY, int32(glfw.KeyY), int32(glfw.GetKeyScancode(glfw.KeyY)))
	g.imGuiIO.SetKeyEventNativeData(imgui.KeyZ, int32(glfw.KeyZ), int32(glfw.GetKeyScancode(glfw.KeyZ)))
}

func (g *GLFW) installCallbacks() {
	g.window.SetMouseButtonCallback(g.mouseButtonChange)
	g.window.SetScrollCallback(g.mouseScrollChange)
	g.window.SetKeyCallback(g.keyChange)
	g.window.SetCharCallback(g.charChange)
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: mouseButtonPrimary,
	glfw.MouseButton2: mouseButtonSecondary,
	glfw.MouseButton3: mouseButtonTertiary,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	mouseButtonPrimary:   glfw.MouseButton1,
	mouseButtonSecondary: glfw.MouseButton2,
	mouseButtonTertiary:  glfw.MouseButton3,
}

func (g *GLFW) mouseButtonChange(_ *glfw.Window, rawButton glfw.MouseButton, action glfw.Action, _ glfw.ModifierKey) {
	buttonIndex, known := glfwButtonIndexByID[rawButton]

	if known && (action == glfw.Press) {
		g.mouseJustPressed[buttonIndex] = true
	}
}

func (g *GLFW) mouseScrollChange(_ *glfw.Window, x, y float64) {
	g.imGuiIO.AddMouseWheelDelta(float32(x), float32(y))
}

func (g *GLFW) keyChange(_ *glfw.Window, _ glfw.Key, _ int, _ glfw.Action, _ glfw.ModifierKey) {

}

func (g *GLFW) charChange(_ *glfw.Window, char rune) {
	g.imGuiIO.AddInputCharacter(uint32(char))
}

// ClipboardText returns the current clipboard text, if available.
func (g *GLFW) ClipboardText() (string, error) {
	return g.window.GetClipboardString(), nil
}

// SetClipboardText sets the text as the current clipboard text.
func (g *GLFW) SetClipboardText(text string) {
	g.window.SetClipboardString(text)
}
