package Windows

import (
	abstractWindow "github.com/Erickype/GoGameEngine/Window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type data struct {
	title         string
	width         int
	height        int
	eventCallback abstractWindow.EventCallBackFn
	vSync         bool
}

type Window struct {
	glfwWindow glfw.Window
	data       data
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

func (w *Window) Create(props *abstractWindow.Properties) abstractWindow.IWindow {
	//TODO implement me
	panic("implement me")
}

func (w *Window) Init() {
	//TODO implement me
	panic("implement me")
}
