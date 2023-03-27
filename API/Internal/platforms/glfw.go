package platforms

import (
	"fmt"
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/go-gl/glfw/v3.3/glfw"
	"runtime"
	"unsafe"
)

type GLFWClientAPI string

// This is a list of GLFWClientAPI constants.
const (
	GLFWClientAPIOpenGL2 GLFWClientAPI = "OpenGL2"
	GLFWClientAPIOpenGL3 GLFWClientAPI = "OpenGL3"
)

// GLFW implements a platform based on github.com/go-gl/glfw (v3.3).
type GLFW struct {
	imGuiIO imgui.ImGuiIO
	window  *glfw.Window
	time    float64
}

func (g *GLFW) SetUserPointer(pointer unsafe.Pointer) {
	g.window.SetUserPointer(pointer)
}

func (g *GLFW) GetWindowPtr() unsafe.Pointer {
	return unsafe.Pointer(g.window)
}

func (g *GLFW) GetWindow() *glfw.Window {
	return g.window
}

func NewGLFW(io imgui.ImGuiIO, clientAPI GLFWClientAPI, width int, height int, title string) (*GLFW, error) {
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
	g.imGuiIO.SetDisplaySize(imgui.ImVec2{X: displaySize[0], Y: displaySize[1]})

	// Setup time step
	currentTime := glfw.GetTime()
	if g.time > 0 {
		g.imGuiIO.SetDeltaTime(float32(currentTime - g.time))
	}
	g.time = currentTime
}

// PostRender performs a buffer swap.
func (g *GLFW) PostRender() {
	g.window.SwapBuffers()
}

func (g *GLFW) setKeyMapping() {
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Tab, int32(glfw.KeyTab), int32(glfw.GetKeyScancode(glfw.KeyTab)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_LeftArrow, int32(glfw.KeyLeft), int32(glfw.GetKeyScancode(glfw.KeyLeft)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_RightArrow, int32(glfw.KeyRight), int32(glfw.GetKeyScancode(glfw.KeyRight)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_UpArrow, int32(glfw.KeyUp), int32(glfw.GetKeyScancode(glfw.KeyUp)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_DownArrow, int32(glfw.KeyDown), int32(glfw.GetKeyScancode(glfw.KeyDown)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_PageUp, int32(glfw.KeyPageUp), int32(glfw.GetKeyScancode(glfw.KeyPageUp)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_PageDown, int32(glfw.KeyPageDown), int32(glfw.GetKeyScancode(glfw.KeyPageDown)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Home, int32(glfw.KeyHome), int32(glfw.GetKeyScancode(glfw.KeyHome)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_End, int32(glfw.KeyEnd), int32(glfw.GetKeyScancode(glfw.KeyEnd)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Insert, int32(glfw.KeyInsert), int32(glfw.GetKeyScancode(glfw.KeyInsert)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Delete, int32(glfw.KeyDelete), int32(glfw.GetKeyScancode(glfw.KeyDelete)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Backspace, int32(glfw.KeyBackspace), int32(glfw.GetKeyScancode(glfw.KeyBackspace)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Space, int32(glfw.KeySpace), int32(glfw.GetKeyScancode(glfw.KeySpace)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Enter, int32(glfw.KeyEnter), int32(glfw.GetKeyScancode(glfw.KeyEnter)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Escape, int32(glfw.KeyEscape), int32(glfw.GetKeyScancode(glfw.KeyEscape)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_A, int32(glfw.KeyA), int32(glfw.GetKeyScancode(glfw.KeyA)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_C, int32(glfw.KeyC), int32(glfw.GetKeyScancode(glfw.KeyC)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_V, int32(glfw.KeyV), int32(glfw.GetKeyScancode(glfw.KeyV)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_X, int32(glfw.KeyX), int32(glfw.GetKeyScancode(glfw.KeyX)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Y, int32(glfw.KeyY), int32(glfw.GetKeyScancode(glfw.KeyY)))
	g.imGuiIO.SetKeyEventNativeData(imgui.ImGuiKey_Z, int32(glfw.KeyZ), int32(glfw.GetKeyScancode(glfw.KeyZ)))
}

// ClipboardText returns the current clipboard text, if available.
func (g *GLFW) ClipboardText() (string, error) {
	return g.window.GetClipboardString(), nil
}

// SetClipboardText sets the text as the current clipboard text.
func (g *GLFW) SetClipboardText(text string) {
	g.window.SetClipboardString(text)
}
