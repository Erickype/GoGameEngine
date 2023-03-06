package Windows

import (
	common "github.com/Erickype/GoGameEngine/Common"
	abstractWindow "github.com/Erickype/GoGameEngine/Window"
	"github.com/akiross/go-glad"
	"github.com/go-gl/glfw/v3.2/glfw"
	"unsafe"
)

type data struct {
	title         string
	width         int
	height        int
	eventCallback *abstractWindow.EventCallBackFn
	vSync         bool
}

type Window struct {
	GlfwWindow *glfw.Window
	data       *data
}

func (w *Window) GetWidth() int {
	return w.data.width
}

func (w *Window) GetHeight() int {
	return w.data.height
}

func (w *Window) SetEventCallback(callback *abstractWindow.EventCallBackFn) {
	w.data.eventCallback = callback
}

func (w *Window) SetVsync(enabled bool) {
	if enabled {
		glfw.SwapInterval(1)
	} else {
		glfw.SwapInterval(0)
	}
	w.data.vSync = enabled
}

func (w *Window) IsVSync() bool {
	return w.data.vSync
}

func (w *Window) OnUpdate() {
	w.GlfwWindow.SwapBuffers()
	glfw.PollEvents()
}

func (w *Window) Shutdown() {
	w.GlfwWindow.Destroy()
}

func (w *Window) Init() {

	common.CoreLogger.Info("Creating window", w.data.title, w.data.width, w.data.height)

	w.GlfwWindow = initGlfw(w)

	w.GlfwWindow.SetUserPointer(unsafe.Pointer(w.data))

	declareCallbacks(w)
}

func initGlfw(w *Window) *glfw.Window {
	return glad.NewOGLWindow(w.data.width, w.data.height, w.data.title,
		glad.CoreProfile(true),
		glad.Resizable(true),
		glad.ContextVersion(4, 4))
}

func Create(props *abstractWindow.Properties) *Window {
	window := &Window{
		data: &data{
			title:  props.Title,
			width:  props.Width,
			height: props.Height,
		},
	}
	window.Init()
	return window
}
