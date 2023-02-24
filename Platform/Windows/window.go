package Windows

import (
	"github.com/Erickype/GoGameEngine/Log"
	abstractWindow "github.com/Erickype/GoGameEngine/Window"
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

var GLFWInitialized bool = false

type data struct {
	title         string
	width         int
	height        int
	eventCallback abstractWindow.EventCallBackFn
	vSync         bool
}

type Window struct {
	glfwWindow *glfw.Window
	data       *data
}

func (w *Window) GetWidth() int {
	//TODO implement me
	panic("implement me")
}

func (w *Window) GetHeight() int {
	//TODO implement me
	panic("implement me")
}

func (w *Window) SetEventCallback(callback *abstractWindow.EventCallBackFn) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (w *Window) OnUpdate() {
	//TODO implement me
	panic("implement me")
}

func (w *Window) Shutdown() {
	w.glfwWindow.Destroy()
}

func (w *Window) Init() {

	Log.InstanceCoreLogger.Info("Creating window", w.data.title, w.data.width, w.data.height)

	if !GLFWInitialized {
		err := glfw.Init()
		if err != nil {
			Log.InstanceCoreLogger.Fatal(err)
		}
		defer glfw.Terminate()
		GLFWInitialized = true
	}

	window, err := glfw.CreateWindow(w.data.width, w.data.height, w.data.title, nil, nil)
	if err != nil {
		Log.InstanceCoreLogger.Fatal(err)
	}
	w.glfwWindow = window
	w.glfwWindow.MakeContextCurrent()
	w.glfwWindow.SetUserPointer(unsafe.Pointer(&w.data))
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
