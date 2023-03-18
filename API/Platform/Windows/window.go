package Windows

import (
	common "github.com/Erickype/GoGameEngine/API/Common"
	"github.com/Erickype/GoGameEngine/API/Internal"
	"github.com/Erickype/GoGameEngine/API/Internal/platforms"
	"github.com/Erickype/GoGameEngine/API/Internal/renderers"
	abstractWindow "github.com/Erickype/GoGameEngine/API/Window"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
	"os"
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
	data     *data
	Platform Internal.IPlatform
	Renderer Internal.IRenderer
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
	w.Platform.PostRender()
	w.Platform.ProcessEvents()
}

func (w *Window) Shutdown() {
	w.Platform.(*platforms.GLFW).Dispose()
	w.Renderer.(*renderers.OpenGL3).Dispose()
}

func (w *Window) Init() {

	common.CoreLogger.Info("Creating window", w.data.title, w.data.width, w.data.height)

	//This creates the imGui context and IO for platform creation
	imgui.CreateContext(nil)
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3, w.data.width, w.data.height, w.data.title)
	if err != nil {
		common.CoreLogger.Fatal("Failing creating platform: ", os.Stderr)
	}
	w.Platform = platform

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		common.CoreLogger.Fatal("Failing creating renderer: ", os.Stderr)
	}
	w.Renderer = renderer

	w.Platform.(*platforms.GLFW).SetUserPointer(unsafe.Pointer(w.data))

	declareCallbacks(w)
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
