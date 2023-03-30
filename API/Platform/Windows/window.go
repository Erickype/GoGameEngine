package Windows

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/Erickype/GoGameEngine/API/Internal"
	"github.com/Erickype/GoGameEngine/API/Internal/platforms"
	"github.com/Erickype/GoGameEngine/API/Internal/renderers"
	"github.com/Erickype/GoGameEngine/API/Log"
	"github.com/Erickype/GoGameEngine/API/Platform/Opengl"
	abstractWindow "github.com/Erickype/GoGameEngine/API/Window"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	data            *data
	platform        *Internal.IPlatform
	renderer        *Internal.IRenderer
	graphicsContext *Opengl.Context
}

func (w *Window) GetPlatform() *Internal.IPlatform {
	return w.platform
}

func (w *Window) GetRenderer() *Internal.IRenderer {
	return w.renderer
}

func (w *Window) GetNativeWindow() unsafe.Pointer {
	return (*w.platform).GetWindowPtr()
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
	w.graphicsContext.SwapBuffers()
	(*w.platform).ProcessEvents()
}

func (w *Window) Shutdown() {
	(*w.platform).(*platforms.GLFW).Dispose()
	(*w.renderer).(*renderers.OpenGL3).Dispose()
}

func (w *Window) Init() {

	Log.GetCoreInstance().Info("Creating window", w.data.title, w.data.width, w.data.height)

	//This creates the imGui context and IO for platform creation
	imgui.CreateContext()
	io := imgui.GetIO()
	io.SetConfigFlags(io.GetConfigFlags() | imgui.ImGuiConfigFlags_NavEnableKeyboard)
	io.SetConfigFlags(io.GetConfigFlags() | imgui.ImGuiConfigFlags_NavEnableGamepad)
	io.SetConfigFlags(io.GetConfigFlags() | imgui.ImGuiConfigFlags_DockingEnable)
	io.SetConfigFlags(io.GetConfigFlags() | imgui.ImGuiConfigFlags_ViewportsEnable)

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3, w.data.width, w.data.height, w.data.title)
	if err != nil {
		Log.GetCoreInstance().Fatal("Failing creating platform: ", os.Stderr)
	}
	iPlatform := Internal.IPlatform(platform)
	w.platform = &iPlatform

	//Opengl context
	w.graphicsContext = Opengl.NewOpenglContext((*glfw.Window)((*w.GetPlatform()).GetWindowPtr()))
	w.graphicsContext.Init()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		Log.GetCoreInstance().Fatal("Failing creating renderer: ", os.Stderr)
	}
	iRenderer := Internal.IRenderer(renderer)
	w.renderer = &iRenderer

	(*w.platform).(*platforms.GLFW).SetUserPointer(unsafe.Pointer(w.data))

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

func CreateAbstractWindow(title string, width int, height int) *abstractWindow.IWindow {
	window := Create(&abstractWindow.Properties{
		Title:  title,
		Width:  width,
		Height: height,
	})

	iWindow := abstractWindow.IWindow(window)
	return &iWindow
}
