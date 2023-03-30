package Core

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Internal"
	"github.com/Erickype/GoGameEngine/API/Log"
	"github.com/Erickype/GoGameEngine/API/Window"
	"github.com/go-gl/gl/v4.1-compatibility/gl"
	"unsafe"
)

type IApplication interface {
	Run()
	Destroy()
	init(window *Window.IWindow)
	onEvent(event *Events.IEvent)
	PushLayer(layer *ILayer)
	PushOverlay(overlay *ILayer)
	GetPlatform() *Internal.IPlatform
	GetRenderer() *Internal.IRenderer
}

var (
	ApplicationInstance *Application
	vertexArray         uint32
	vertexBuffer        uint32
	indexBuffer         uint32
)

type Application struct {
	window     *Window.IWindow
	running    bool
	layerStack *LayerStack
	imGuiLayer *ImGuiLayer
}

func (a *Application) Run() {
	for !(*(*a.window).GetPlatform()).ShouldStop() {
		(*ApplicationInstance.GetRenderer()).PreRender(clearColor)

		gl.BindVertexArray(vertexArray)
		gl.DrawElements(gl.TRIANGLES, 3, gl.UNSIGNED_INT, nil)

		if *a.layerStack.layerInsert != 0 {
			for _, layer := range *a.layerStack.layers {
				(*layer).OnUpdate()
			}
		}

		a.imGuiLayer.Begin()
		for _, layer := range *a.layerStack.layers {
			(*layer).OnImGuiRender()
		}
		a.imGuiLayer.End()

		(*a.window).OnUpdate()
	}
}

func (a *Application) Destroy() {
	a.running = false
	(*a.window).Shutdown()
}

func (a *Application) init(window *Window.IWindow) {
	Log.GetCoreInstance().Info("Starting engine!!")
	a.running = true
	a.window = window

	eventCallbackFn := Window.EventCallBackFn(func(event *Events.IEvent) {
		a.onEvent(event)
	})
	(*a.window).SetEventCallback(&eventCallbackFn)
	a.layerStack = &LayerStack{}
	a.layerStack.Construct()
	//ImguiLayer Init
	a.imGuiLayer = NewImGui()
	iImGui := ILayer(a.imGuiLayer)
	a.layerStack.PushOverlay(&iImGui)
	gl.GenVertexArrays(1, &vertexArray)
	gl.BindVertexArray(vertexArray)

	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)

	vertices := [][]float32{
		{-.5, -.5, 0},
		{.5, .5, 0},
		{0, .5, 0}}

	gl.BufferData(gl.ARRAY_BUFFER, len(vertices), unsafe.Pointer(&vertices), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)

	gl.GenBuffers(1, &indexBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer)

	indices := []int{0, 1, 2}

	gl.BufferData(gl.ARRAY_BUFFER, len(indices), unsafe.Pointer(&indices), gl.STATIC_DRAW)
}

func (a *Application) onEvent(event *Events.IEvent) {
	if *a.layerStack.layerInsert != 0 {
		for i := len(*a.layerStack.layers) - 1; i >= 0; i-- {
			layer := (*a.layerStack.layers)[i]
			(*layer).OnEvent(event)
			if (*event).WasHandled() {
				break
			}
		}
	}
}

func (a *Application) PushLayer(layer *ILayer) {
	(*layer).OnAttach()
	a.layerStack.PushLayer(layer)
}

func (a *Application) PushOverlay(overlay *ILayer) {
	a.layerStack.PushOverlay(overlay)
}

func (a *Application) GetPlatform() *Internal.IPlatform {
	return (*a.window).GetPlatform()
}

func (a *Application) GetRenderer() *Internal.IRenderer {
	return (*a.window).GetRenderer()
}

func (a *Application) Construct(window *Window.IWindow) {
	ApplicationInstance = a
	a.init(window)
}
