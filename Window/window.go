package Window

import "github.com/Erickype/GoGameEngine/Events"

type EventCallBackFn func(event *Events.IEvent)

type Properties struct {
	Title  string
	Width  int
	Height int
}

type IWindow interface {
	GetWidth() int
	GetHeight() int
	SetEventCallback(callback *EventCallBackFn)
	SetVsync(enabled bool)
	IsVSync() bool
	OnUpdate()
	Shutdown()
	Init()
}
