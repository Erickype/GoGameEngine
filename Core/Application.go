package Core

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Internal"
	"github.com/Erickype/GoGameEngine/API/Log"
	"github.com/Erickype/GoGameEngine/API/Window"
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

var ApplicationInstance *Application

type Application struct {
	window     *Window.IWindow
	running    bool
	layerStack *LayerStack
}

func (a *Application) Run() {
	for !(*(*a.window).GetPlatform()).ShouldStop() {
		if *a.layerStack.layerInsert != 0 {
			for _, layer := range *a.layerStack.layers {
				(*layer).OnUpdate()
			}
		}
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
