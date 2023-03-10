package Core

import (
	"github.com/Erickype/GoGameEngine/API/Common"
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Platform/Windows"
	"github.com/Erickype/GoGameEngine/API/Window"
)

type IApplication interface {
	run()
	destroy()
	init(layer *ILayer)
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
		if *a.layerStack.layerInsert != 0 {
			for _, layer := range *a.layerStack.layers {
				(*layer).OnUpdate()
			}
		}
		a.window.OnUpdate()
	}
}

func (a *Application) destroy() {
	a.running = false
}

func (a *Application) init(layer *ILayer) {
	Common.CoreLogger.Info("Starting engine!!")
	a.running = true
	a.window = Windows.Create(&Window.Properties{
		Title:  "GoGameEngine",
		Width:  800,
		Height: 600,
	})

	eventCallbackFn := Window.EventCallBackFn(func(event *Events.IEvent) {
		a.onEvent(event)
	})
	a.window.SetEventCallback(&eventCallbackFn)
	a.layerStack = &LayerStack{}
	a.layerStack.Construct()
	a.PushLayer(layer)
}

func (a *Application) onEvent(event *Events.IEvent) {
	Common.EventDispatcher.Dispatch(*event)
	Common.CoreLogger.Trace((*event).ToString())

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
	a.layerStack.PushLayer(layer)
}

func (a *Application) PushOverlay(overlay *ILayer) {
	a.layerStack.PushOverlay(overlay)
}

// CreateApplication This is the entry point to create an application
func CreateApplication(layer *ILayer) {
	application := &Application{}
	application.init(layer)
	application.run()
	application.destroy()
}
