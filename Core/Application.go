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
}

type Application struct {
	window  *Windows.Window
	running bool
}

func (a *Application) run() {
	a.window.OnUpdate()
}

func (a *Application) destroy() {}

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
}

func (a *Application) onEvent(event *Events.IEvent) {
	common.EventDispatcher.Dispatch(*event)
	common.CoreLogger.Trace((*event).ToString())
}

// CreateApplication This is the entry point to create an application
func CreateApplication() {
	application := &Application{}
	application.init()
	application.run()
	application.destroy()
}
