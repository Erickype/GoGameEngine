package Event

type Types int

const (
	None Types = iota
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

func (e Types) String() string {
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
