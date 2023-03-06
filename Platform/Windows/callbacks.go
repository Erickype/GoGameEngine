package Windows

import (
	common "github.com/Erickype/GoGameEngine/Common"
	"github.com/Erickype/GoGameEngine/Events"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func declareCallbacks(w *Window) {
	w.setSizeCallback()
	w.setCloseCallback()
	w.setKeyCallback()
	w.setMouseCallback()
	w.setScrollCallback()
	w.setCursorPosCallback()
}

func (w *Window) setSizeCallback() {
	w.GlfwWindow.SetSizeCallback(func(window *glfw.Window, width int, height int) {
		data := (*data)(window.GetUserPointer())
		data.width = width
		data.height = height

		event := common.EventFactory.CreateEvent(Events.WindowResize)

		if resizeEvent, ok := event.(*Events.WindowResizeEvent); ok {
			resizeEvent.Width = width
			resizeEvent.Height = height
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCloseCallback() {
	w.GlfwWindow.SetCloseCallback(func(window *glfw.Window) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events.WindowClose)
		if data.eventCallback != nil {
			(*data.eventCallback)(&event)
		}
	})
}

func (w *Window) setKeyCallback() {
	w.GlfwWindow.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := common.EventFactory.CreateEvent(Events.KeyPressed)
			if keyPressedEvent, ok := event.(*Events.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 0
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := common.EventFactory.CreateEvent(Events.KeyReleased)
			if keyPressedEvent, ok := event.(*Events.KeyReleasedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Repeat:
			event := common.EventFactory.CreateEvent(Events.KeyPressed)
			if keyPressedEvent, ok := event.(*Events.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 1
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setMouseCallback() {
	w.GlfwWindow.SetMouseButtonCallback(func(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := common.EventFactory.CreateEvent(Events.MouseButtonPressed)
			if keyPressedEvent, ok := event.(*Events.MouseButtonPressedEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := common.EventFactory.CreateEvent(Events.MouseButtonReleased)
			if keyPressedEvent, ok := event.(*Events.MouseButtonReleaseEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setScrollCallback() {
	w.GlfwWindow.SetScrollCallback(func(window *glfw.Window, xOff float64, yOff float64) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events.MouseScrolled)
		if mouseScrolledEvent, ok := event.(*Events.MouseScrolledEvent); ok {
			mouseScrolledEvent.XOffset = xOff
			mouseScrolledEvent.YOffset = yOff
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCursorPosCallback() {
	w.GlfwWindow.SetCursorPosCallback(func(window *glfw.Window, xPos float64, yPos float64) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events.MouseMoved)
		if mouseMovedEvent, ok := event.(*Events.MouseMovedEvent); ok {
			mouseMovedEvent.MouseX = xPos
			mouseMovedEvent.MouseY = yPos
		}
		(*data.eventCallback)(&event)
	})
}
