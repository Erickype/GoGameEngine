package Window

import "github.com/Erickype/GoGameEngine/Events"

type EventCallBackFn func(event *Events.IEvent)

type Properties struct {
	title  string
	width  int
	height int
}

var windowProps *Properties

func init() {
	windowProps = &Properties{
		title:  "GoGameEngine",
		width:  1280,
		height: 720,
	}
}

// SetWindowProps will be used to set new title, width, height
func SetWindowProps(title string, width int, height int) {
	windowProps.title = title
	windowProps.width = width
	windowProps.height = height
}

type IWindow interface {
	GetWidth() int
	GetHeight() int
	SetEventCallback(callback *EventCallBackFn)
	SetVsync(enabled bool)
	IsVSync() bool
	OnUpdate()
	Create(props *Properties) IWindow
	Init()
}
