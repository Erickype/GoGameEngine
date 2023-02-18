package Events

type IMouseMovedEvent interface {
	GetX() float64
	GetY() float64
	Init()
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
