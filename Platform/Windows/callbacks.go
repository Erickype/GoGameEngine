package Windows

import (
	common "github.com/Erickype/GoGameEngine/Common"
	"github.com/Erickype/GoGameEngine/Events"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func declareCallbacks(w *Window) {
	w.setSizeCallback()
	w.setCloseCallback()
	w.setKeyCallback()
}

func (w *Window) setSizeCallback() {
	w.glfwWindow.SetSizeCallback(func(window *glfw.Window, width int, height int) {
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
	w.glfwWindow.SetCloseCallback(func(window *glfw.Window) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events.WindowClose)
		if data.eventCallback != nil {
			(*data.eventCallback)(&event)
		}
	})
}

func (w *Window) setKeyCallback() {
	w.glfwWindow.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
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
