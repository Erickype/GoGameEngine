package Windows

import (
	common "github.com/Erickype/GoGameEngine/Common"
	abstractWindow "github.com/Erickype/GoGameEngine/Window"
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

var glfwInitialized = false

type data struct {
	title         string
	width         int
	height        int
	eventCallback *abstractWindow.EventCallBackFn
	vSync         bool
}

type Window struct {
	glfwWindow *glfw.Window
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
	for !w.glfwWindow.ShouldClose() {
		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *Window) Shutdown() {
	w.glfwWindow.Destroy()
}

func (w *Window) Init() {

	common.CoreLogger.Info("Creating window", w.data.title, w.data.width, w.data.height)

	initGlfw()

	window, err := glfw.CreateWindow(w.data.width, w.data.height, w.data.title, nil, nil)
	if err != nil {
		common.CoreLogger.Fatal(err)
	}
	w.glfwWindow = window
	w.glfwWindow.MakeContextCurrent()
	w.glfwWindow.SetUserPointer(unsafe.Pointer(w.data))
}

func initGlfw() {
	if !glfwInitialized {
		if err := glfw.Init(); err != nil {
			common.CoreLogger.Fatal(err)
		}
		common.CoreLogger.Info("GLFW initialized")
		glfwInitialized = true
	}
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
