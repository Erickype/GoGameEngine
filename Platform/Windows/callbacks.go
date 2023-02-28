package Windows

import (
	common "github.com/Erickype/GoGameEngine/Common"
	"github.com/Erickype/GoGameEngine/Events"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func declareCallbacks(w *Window) {
	w.setSizeCallback()
}

func (w *Window) setSizeCallback() {
	w.glfwWindow.SetSizeCallback(func(window *glfw.Window, width int, height int) {
		data := (*data)(window.GetUserPointer())
		data.width = width
		data.height = height

		event := common.EventFactory.CreateEvent(Events.WindowResize)

		// Type assertion to retrieve the WindowResizeEvent value from the IEvent interface type
		if resizeEvent, ok := event.(*Events.WindowResizeEvent); ok {
			// Use the resizeEvent variable of type *Events.WindowResizeEvent to access its properties
			resizeEvent.Width = width
			resizeEvent.Height = height
		}

		if data.eventCallback != nil {
			(*data.eventCallback)(&event)
		}
	})
}
