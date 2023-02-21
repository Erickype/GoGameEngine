package Events

import "fmt"

type EventType int

const (
	None EventType = iota
	WindowClose
	WindowResize
	WindowFocus
	WindowLostFocus
	WindowMoved
	AppTick
	AppUpdate
	AppRender
	KeyPressed
	KeyReleased
	MouseButtonPressed
	MouseButtonReleased
	MouseMoved
	MouseScrolled
)

func (e EventType) String() string {
	switch e {
	case None:
		return "None"
	case WindowClose:
		return "WindowClose"
	case WindowResize:
		return "WindowResize"
	case WindowFocus:
		return "WindowFocus"
	case WindowLostFocus:
		return "WindowLostFocus"
	case WindowMoved:
		return "WindowMoved"
	case AppTick:
		return "AppTick"
	case AppUpdate:
		return "AppUpdate"
	case AppRender:
		return "AppRender"
	case KeyPressed:
		return "KeyPressed"
	case KeyReleased:
		return "KeyReleased"
	case MouseButtonPressed:
		return "MouseButtonPressed"
	case MouseButtonReleased:
		return "MouseButtonReleased"
	case MouseMoved:
		return "MouseMoved"
	case MouseScrolled:
		return "MouseScrolled"
	default:
		return "Unknown"
	}
}

type Category int

const (
	Application Category = iota
	Input
	Keyboard
	Mouse
	MouseButton
)

func (c Category) String() string {
	switch c {
	case Application:
		return "Application"
	case Input:
		return "Input"
	case Keyboard:
		return "Keyboard"
	case Mouse:
		return "Mouse"
	case MouseButton:
		return "MouseButton"
	default:
		return "Unknown"
	}
}

// IEvent interface to implement an Event
type IEvent interface {
	GetEventType() EventType
	GetName() string
	GetCategoryFlags() int
	IsInCategory() bool
	ToString() string
	WasHandled() bool
	Init()
}

// Event struct that implement IEvent, has a reference to EventDispatcher and eventCategory, eventType and handled fields
type Event struct {
	handled       bool
	eventCategory Category
	eventType     EventType
}

func (e *Event) GetEventType() EventType {
	return e.eventType
}

func (e *Event) GetName() string {
	return e.eventType.String()
}

func (e *Event) GetCategoryFlags() int {
	return int(e.eventCategory)
}

func (e *Event) IsInCategory() bool {
	return e.GetCategoryFlags() == int(e.eventCategory)
}

func (e *Event) WasHandled() bool {
	return e.handled
}

func (e *Event) ToString() string {
	return fmt.Sprintf("Category: %s, EventType: %s", e.eventCategory, e.eventType)
}

func (e *Event) Init() {}
