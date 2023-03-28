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
	keyMap  map[glfw.Key]imgui.ImGuiKey
}

func (g *GLFW) GetKeyMap() map[glfw.Key]imgui.ImGuiKey {
	return g.keyMap
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

func (g *GLFW) setKeyMapping() {
	g.keyMap = map[glfw.Key]imgui.ImGuiKey{
		glfw.KeyTab:       imgui.ImGuiKey_Tab,
		glfw.KeyLeft:      imgui.ImGuiKey_LeftArrow,
		glfw.KeyRight:     imgui.ImGuiKey_RightArrow,
		glfw.KeyUp:        imgui.ImGuiKey_UpArrow,
		glfw.KeyDown:      imgui.ImGuiKey_DownArrow,
		glfw.KeyPageUp:    imgui.ImGuiKey_PageUp,
		glfw.KeyPageDown:  imgui.ImGuiKey_PageDown,
		glfw.KeyHome:      imgui.ImGuiKey_Home,
		glfw.KeyEnd:       imgui.ImGuiKey_End,
		glfw.KeyInsert:    imgui.ImGuiKey_Insert,
		glfw.KeyDelete:    imgui.ImGuiKey_Delete,
		glfw.KeyBackspace: imgui.ImGuiKey_Backspace,
		glfw.KeySpace:     imgui.ImGuiKey_Space,
		glfw.KeyEnter:     imgui.ImGuiKey_Enter,
		glfw.KeyEscape:    imgui.ImGuiKey_Escape,
		glfw.KeyA:         imgui.ImGuiKey_A,
		glfw.KeyC:         imgui.ImGuiKey_C,
		glfw.KeyV:         imgui.ImGuiKey_V,
		glfw.KeyX:         imgui.ImGuiKey_X,
		glfw.KeyY:         imgui.ImGuiKey_Y,
		glfw.KeyZ:         imgui.ImGuiKey_Z,

		glfw.KeyLeftControl:  imgui.ImGuiKey_ModCtrl,
		glfw.KeyRightControl: imgui.ImGuiKey_ModCtrl,
		glfw.KeyLeftAlt:      imgui.ImGuiKey_ModAlt,
		glfw.KeyRightAlt:     imgui.ImGuiKey_ModAlt,
		glfw.KeyLeftSuper:    imgui.ImGuiKey_ModSuper,
		glfw.KeyRightSuper:   imgui.ImGuiKey_ModSuper,
		glfw.KeyLeftShift:    imgui.ImGuiKey_ModShift,
		glfw.KeyRightShift:   imgui.ImGuiKey_ModShift,
	}
}

// ClipboardText returns the current clipboard text, if available.
func (g *GLFW) ClipboardText() (string, error) {
	return g.window.GetClipboardString(), nil
}

// SetClipboardText sets the text as the current clipboard text.
func (g *GLFW) SetClipboardText(text string) {
	g.window.SetClipboardString(text)
}
