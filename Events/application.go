package Events

import "fmt"

// IWindowResizeEvent interface to implement WindowResizeEvent
type IWindowResizeEvent interface {
	GetWidth() int
	GetHeight() int
}

type WindowResizeEvent struct {
	*Event
	width  int
	height int
}

func (w *WindowResizeEvent) GetWidth() int {
	return w.width
}

func (w *WindowResizeEvent) GetHeight() int {
	return w.height
}

func (w *WindowResizeEvent) ToString() string {
	return fmt.Sprintf("WindowResizeEvent: %d, %d", w.width, w.height)
}

func (w *WindowResizeEvent) Init() {
	w.eventType = WindowResize
	w.eventCategory = Application
}
