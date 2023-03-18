package platforms

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
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
	imGuiIO imgui.IO
	window  *glfw.Window
	time    float64
}

func (g *GLFW) SetUserPointer(pointer unsafe.Pointer) {
	g.window.SetUserPointer(pointer)
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
}

// PostRender performs a buffer swap.
func (g *GLFW) PostRender() {
	g.window.SwapBuffers()
}

func (g *GLFW) setKeyMapping() {
	g.imGuiIO.KeyMap(imgui.KeyTab, int(glfw.KeyTab))
	g.imGuiIO.KeyMap(imgui.KeyLeftArrow, int(glfw.KeyLeft))
	g.imGuiIO.KeyMap(imgui.KeyRightArrow, int(glfw.KeyRight))
	g.imGuiIO.KeyMap(imgui.KeyUpArrow, int(glfw.KeyUp))
	g.imGuiIO.KeyMap(imgui.KeyDownArrow, int(glfw.KeyDown))
	g.imGuiIO.KeyMap(imgui.KeyPageUp, int(glfw.KeyPageUp))
	g.imGuiIO.KeyMap(imgui.KeyPageDown, int(glfw.KeyPageDown))
	g.imGuiIO.KeyMap(imgui.KeyHome, int(glfw.KeyHome))
	g.imGuiIO.KeyMap(imgui.KeyEnd, int(glfw.KeyEnd))
	g.imGuiIO.KeyMap(imgui.KeyInsert, int(glfw.KeyInsert))
	g.imGuiIO.KeyMap(imgui.KeyDelete, int(glfw.KeyDelete))
	g.imGuiIO.KeyMap(imgui.KeyBackspace, int(glfw.KeyBackspace))
	g.imGuiIO.KeyMap(imgui.KeySpace, int(glfw.KeySpace))
	g.imGuiIO.KeyMap(imgui.KeyEnter, int(glfw.KeyEnter))
	g.imGuiIO.KeyMap(imgui.KeyEscape, int(glfw.KeyEscape))
	g.imGuiIO.KeyMap(imgui.KeyA, int(glfw.KeyA))
	g.imGuiIO.KeyMap(imgui.KeyC, int(glfw.KeyC))
	g.imGuiIO.KeyMap(imgui.KeyV, int(glfw.KeyV))
	g.imGuiIO.KeyMap(imgui.KeyX, int(glfw.KeyX))
	g.imGuiIO.KeyMap(imgui.KeyY, int(glfw.KeyY))
	g.imGuiIO.KeyMap(imgui.KeyZ, int(glfw.KeyZ))
}

// ClipboardText returns the current clipboard text, if available.
func (g *GLFW) ClipboardText() (string, error) {
	return g.window.GetClipboardString(), nil
}

// SetClipboardText sets the text as the current clipboard text.
func (g *GLFW) SetClipboardText(text string) {
	g.window.SetClipboardString(text)
}
