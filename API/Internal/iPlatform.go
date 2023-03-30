package Internal

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

// IPlatform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type IPlatform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the beginning of a render pass. It must update the imGui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)

	GetWindowPtr() unsafe.Pointer

	GetKeyMap() map[glfw.Key]imgui.ImGuiKey
}

// IRenderer covers rendering imGui draw data.
type IRenderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imGui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.ImDrawData)
}
