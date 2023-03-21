package Core

import (
	"github.com/Erickype/GoGameEngine/API"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Internal"
	"github.com/Erickype/GoGameEngine/API/Log"
	"github.com/Erickype/GoGameEngine/API/Window"
)

type IApplication interface {
	run()
	destroy()
	init(layer *ILayer, window *Window.IWindow)
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

func (a *Application) run() {
	for !(*(*a.window).GetPlatform()).ShouldStop() {
		if *a.layerStack.layerInsert != 0 {
			for _, layer := range *a.layerStack.layers {
				(*layer).OnUpdate()
			}
		}
		Log.GetCoreInstance().Debug((*API.GetInputInstance()).IsKeyPressed(32, (*ApplicationInstance.GetPlatform()).GetWindowPtr()))
		(*a.window).OnUpdate()
	}
}

func (a *Application) destroy() {
	a.running = false
	(*a.window).Shutdown()
}

func (a *Application) init(layer *ILayer, window *Window.IWindow) {
	Log.GetCoreInstance().Info("Starting engine!!")
	a.running = true
	a.window = window

	eventCallbackFn := Window.EventCallBackFn(func(event *Events.IEvent) {
		a.onEvent(event)
	})
	(*a.window).SetEventCallback(&eventCallbackFn)
	a.layerStack = &LayerStack{}
	a.layerStack.Construct()
	a.PushLayer(layer)
}

func (a *Application) onEvent(event *Events.IEvent) {
	Log.GetCoreInstance().Trace((*event).ToString())

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

// CreateApplication This is the entry point to create an application
func CreateApplication(layer *ILayer, window *Window.IWindow) {
	application := &Application{}
	ApplicationInstance = application
	application.init(layer, window)
	application.run()
	application.destroy()
}
