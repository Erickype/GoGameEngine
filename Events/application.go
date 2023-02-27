package Events

import "fmt"

// IWindowResizeEvent interface to implement WindowResizeEvent
type IWindowResizeEvent interface {
	GetWidth() int
	GetHeight() int
}

type WindowResizeEvent struct {
	*Event
	Width  int
	Height int
}

func (w *WindowResizeEvent) GetWidth() int {
	return w.Width
}

func (w *WindowResizeEvent) GetHeight() int {
	return w.Height
}

func (w *WindowResizeEvent) ToString() string {
	return fmt.Sprintf("WindowResizeEvent: %d, %d", w.Width, w.Height)
}

func (w *WindowResizeEvent) Init() {
	w.eventType = WindowResize
	w.eventCategory = Application
}

// WindowCloseEvent struct to implement the event
type WindowCloseEvent struct {
	*Event
}

func (w *WindowCloseEvent) Init() {
	w.eventType = WindowClose
	w.eventCategory = Application
}
