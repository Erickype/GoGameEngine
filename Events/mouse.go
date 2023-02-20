package Events

import "fmt"

// IMouseMovedEvent interface to implement MouseMovedEvent
type IMouseMovedEvent interface {
	GetX() float64
	GetY() float64
}

type MouseMovedEvent struct {
	*Event
	mouseX float64
	mouseY float64
}

func (m *MouseMovedEvent) GetX() float64 {
	return m.mouseX
}

func (m *MouseMovedEvent) GetY() float64 {
	return m.mouseY
}

func (m *MouseMovedEvent) Init() {
	m.Event.eventType = MouseMoved
	m.Event.eventCategory = Mouse | Input
}

// IMouseScrolledEvent interface to implement MouseScrolledEvent
type IMouseScrolledEvent interface {
	GetXOffset() float64
	GetYOffset() float64
}

type MouseScrolledEvent struct {
	*Event
	xOffset float64
	yOffset float64
}

func (m *MouseScrolledEvent) GetXOffset() float64 {
	return m.xOffset
}

func (m *MouseScrolledEvent) GetYOffset() float64 {
	return m.yOffset
}

func (m *MouseScrolledEvent) ToString() string {
	return fmt.Sprintf("MouseScrolledEvent: %f, %f", m.xOffset, m.yOffset)
}

func (m *MouseScrolledEvent) Init() {
	m.eventType = MouseScrolled
	m.eventCategory = Mouse | Input
}
