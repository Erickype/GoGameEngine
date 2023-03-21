package Windows

import (
	"github.com/Erickype/GoGameEngine/API/Events"
	"github.com/Erickype/GoGameEngine/API/Internal/platforms"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func declareCallbacks(w *Window) {
	w.setSizeCallback()
	w.setCloseCallback()
	w.setKeyCallback()
	w.setMouseCallback()
	w.setScrollCallback()
	w.setCursorPosCallback()
	w.setCharCallback()
}

func (w *Window) setSizeCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetSizeCallback(func(window *glfw.Window, width int, height int) {
		data := (*data)(window.GetUserPointer())
		data.width = width
		data.height = height

		event := Events.GetEventFactoryInstance().CreateEvent(Events.WindowResize)

		if resizeEvent, ok := event.(*Events.WindowResizeEvent); ok {
			resizeEvent.Width = width
			resizeEvent.Height = height
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCloseCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetCloseCallback(func(window *glfw.Window) {
		data := (*data)(window.GetUserPointer())
		event := Events.GetEventFactoryInstance().CreateEvent(Events.WindowClose)
		if data.eventCallback != nil {
			(*data.eventCallback)(&event)
		}
	})
}

func (w *Window) setKeyCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := Events.GetEventFactoryInstance().CreateEvent(Events.KeyPressed)
			if keyPressedEvent, ok := event.(*Events.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 0
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := Events.GetEventFactoryInstance().CreateEvent(Events.KeyReleased)
			if keyPressedEvent, ok := event.(*Events.KeyReleasedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Repeat:
			event := Events.GetEventFactoryInstance().CreateEvent(Events.KeyPressed)
			if keyPressedEvent, ok := event.(*Events.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 1
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setCharCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetCharCallback(func(window *glfw.Window, char rune) {
		data := (*data)(window.GetUserPointer())
		event := Events.GetEventFactoryInstance().CreateEvent(Events.KeyTyped)
		if keyTypedEvent, ok := event.(*Events.KeyTypedEvent); ok {
			keyTypedEvent.KeyCode = int(char)
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setMouseCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetMouseButtonCallback(func(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := Events.GetEventFactoryInstance().CreateEvent(Events.MouseButtonPressed)
			if keyPressedEvent, ok := event.(*Events.MouseButtonPressedEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := Events.GetEventFactoryInstance().CreateEvent(Events.MouseButtonReleased)
			if keyPressedEvent, ok := event.(*Events.MouseButtonReleaseEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setScrollCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetScrollCallback(func(window *glfw.Window, xOff float64, yOff float64) {
		data := (*data)(window.GetUserPointer())
		event := Events.GetEventFactoryInstance().CreateEvent(Events.MouseScrolled)
		if mouseScrolledEvent, ok := event.(*Events.MouseScrolledEvent); ok {
			mouseScrolledEvent.XOffset = xOff
			mouseScrolledEvent.YOffset = yOff
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCursorPosCallback() {
	(*w.GetPlatform()).(*platforms.GLFW).GetWindow().SetCursorPosCallback(func(window *glfw.Window, xPos float64, yPos float64) {
		data := (*data)(window.GetUserPointer())
		event := Events.GetEventFactoryInstance().CreateEvent(Events.MouseMoved)
		if mouseMovedEvent, ok := event.(*Events.MouseMovedEvent); ok {
			mouseMovedEvent.MouseX = xPos
			mouseMovedEvent.MouseY = yPos
		}
		(*data.eventCallback)(&event)
	})
}
