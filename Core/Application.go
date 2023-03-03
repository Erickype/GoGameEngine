package Core

import (
	common "github.com/Erickype/GoGameEngine/Common"
	"github.com/Erickype/GoGameEngine/Events"
	"github.com/Erickype/GoGameEngine/Platform/Windows"
	"github.com/Erickype/GoGameEngine/Window"
)

type IApplication interface {
	run()
	destroy()
	init()
	onEvent(event *Events.IEvent)
	PushLayer(layer *ILayer)
	PushOverlay(overlay *ILayer)
}

type Application struct {
	window     *Windows.Window
	running    bool
	layerStack *LayerStack
}

func (a *Application) run() {
	for !a.window.GlfwWindow.ShouldClose() {
		for _, layer := range *a.layerStack.layers {
			(*layer).OnUpdate()
		}
		a.window.OnUpdate()
	}
}

func (a *Application) destroy() {
	a.running = false
}

func (a *Application) init() {
	common.CoreLogger.Info("Starting engine!!")
	a.running = true
	a.window = Windows.Create(&Window.Properties{
		Title:  "GoGameEngine",
		Width:  1280,
		Height: 720,
	})

	eventCallbackFn := Window.EventCallBackFn(func(event *Events.IEvent) {
		a.onEvent(event)
	})
	a.window.SetEventCallback(&eventCallbackFn)
	a.layerStack.Construct()
}

func (a *Application) onEvent(event *Events.IEvent) {
	common.EventDispatcher.Dispatch(*event)
	common.CoreLogger.Trace((*event).ToString())

	for i := len(*a.layerStack.layers) - 1; i >= 0; i-- {
		layer := (*a.layerStack.layers)[i]
		(*layer).OnEvent(event)
		if (*event).WasHandled() {
			break
		}
	}

}

func (a *Application) PushLayer(layer *ILayer) {
	a.layerStack.PushLayer(layer)
}

func (a *Application) PushOverlay(overlay *ILayer) {
	a.layerStack.PushOverlay(overlay)
}

// CreateApplication This is the entry point to create an application
func CreateApplication() {
	application := &Application{}
	application.init()
	application.run()
	application.destroy()
}
