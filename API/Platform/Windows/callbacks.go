package Windows

import (
	common "github.com/Erickype/GoGameEngine/API/Common"
	Events2 "github.com/Erickype/GoGameEngine/API/Events"
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
}

func (w *Window) setSizeCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetSizeCallback(func(window *glfw.Window, width int, height int) {
		data := (*data)(window.GetUserPointer())
		data.width = width
		data.height = height

		event := common.EventFactory.CreateEvent(Events2.WindowResize)

		if resizeEvent, ok := event.(*Events2.WindowResizeEvent); ok {
			resizeEvent.Width = width
			resizeEvent.Height = height
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCloseCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetCloseCallback(func(window *glfw.Window) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events2.WindowClose)
		if data.eventCallback != nil {
			(*data.eventCallback)(&event)
		}
	})
}

func (w *Window) setKeyCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := common.EventFactory.CreateEvent(Events2.KeyPressed)
			if keyPressedEvent, ok := event.(*Events2.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 0
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := common.EventFactory.CreateEvent(Events2.KeyReleased)
			if keyPressedEvent, ok := event.(*Events2.KeyReleasedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Repeat:
			event := common.EventFactory.CreateEvent(Events2.KeyPressed)
			if keyPressedEvent, ok := event.(*Events2.KeyPressedEvent); ok {
				keyPressedEvent.KeyCode = (int)(key)
				keyPressedEvent.RepeatCount = 1
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setMouseCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetMouseButtonCallback(func(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		data := (*data)(window.GetUserPointer())
		switch action {
		case glfw.Press:
			event := common.EventFactory.CreateEvent(Events2.MouseButtonPressed)
			if keyPressedEvent, ok := event.(*Events2.MouseButtonPressedEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		case glfw.Release:
			event := common.EventFactory.CreateEvent(Events2.MouseButtonReleased)
			if keyPressedEvent, ok := event.(*Events2.MouseButtonReleaseEvent); ok {
				keyPressedEvent.Button = (int)(button)
			}
			(*data.eventCallback)(&event)
			break
		}
	})
}

func (w *Window) setScrollCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetScrollCallback(func(window *glfw.Window, xOff float64, yOff float64) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events2.MouseScrolled)
		if mouseScrolledEvent, ok := event.(*Events2.MouseScrolledEvent); ok {
			mouseScrolledEvent.XOffset = xOff
			mouseScrolledEvent.YOffset = yOff
		}
		(*data.eventCallback)(&event)
	})
}

func (w *Window) setCursorPosCallback() {
	w.Platform.(*platforms.GLFW).GetWindow().SetCursorPosCallback(func(window *glfw.Window, xPos float64, yPos float64) {
		data := (*data)(window.GetUserPointer())
		event := common.EventFactory.CreateEvent(Events2.MouseMoved)
		if mouseMovedEvent, ok := event.(*Events2.MouseMovedEvent); ok {
			mouseMovedEvent.MouseX = xPos
			mouseMovedEvent.MouseY = yPos
		}
		(*data.eventCallback)(&event)
	})
}
