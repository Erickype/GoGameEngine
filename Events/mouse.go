package Events

import "fmt"

// IMouseMovedEvent interface to implement MouseMovedEvent
type IMouseMovedEvent interface {
	GetX() float64
	GetY() float64
}

type MouseMovedEvent struct {
	*Event
	MouseX float64
	MouseY float64
}

func (m *MouseMovedEvent) GetX() float64 {
	return m.MouseX
}

func (m *MouseMovedEvent) GetY() float64 {
	return m.MouseY
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
	XOffset float64
	YOffset float64
}

func (m *MouseScrolledEvent) GetXOffset() float64 {
	return m.XOffset
}

func (m *MouseScrolledEvent) GetYOffset() float64 {
	return m.YOffset
}

func (m *MouseScrolledEvent) ToString() string {
	return fmt.Sprintf("MouseScrolledEvent: %f, %f", m.XOffset, m.YOffset)
}

func (m *MouseScrolledEvent) Init() {
	m.eventType = MouseScrolled
	m.eventCategory = Mouse | Input
}

// IMouseButtonEvent common interface to MouseButtonEvents
type IMouseButtonEvent interface {
	GetMouseButton() int
}

type MouseButtonEvent struct {
	*Event
	Button int
}

func (m *MouseButtonEvent) GetMouseButton() int {
	return m.Button
}

// IMouseButtonPressedEvent interface to implement MouseButtonPressedEvent
type IMouseButtonPressedEvent interface{}

type MouseButtonPressedEvent struct {
	*MouseButtonEvent
}

func (m *MouseButtonPressedEvent) ToString() string {
	return fmt.Sprintf("MouseButtonPressedEvent: %d", m.Button)
}

func (m *MouseButtonPressedEvent) Init() {
	m.eventType = MouseButtonPressed
	m.eventCategory = Mouse | Input
}

// IMouseButtonReleaseEvent interface to implement MouseButtonReleaseEvent
type IMouseButtonReleaseEvent interface{}

type MouseButtonReleaseEvent struct {
	*MouseButtonEvent
}

func (m *MouseButtonReleaseEvent) ToString() string {
	return fmt.Sprintf("MouseButtonReleasedEvent: %d", m.Button)
}

func (m *MouseButtonReleaseEvent) Init() {
	m.eventType = MouseButtonReleased
	m.eventCategory = Mouse | Input
}
