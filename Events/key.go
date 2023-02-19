package Events

import "fmt"

type IKeyEvent interface {
	GetKeyCode() int
}

// keyEvent common struct to key events
type keyEvent struct {
	*Event
	keyCode int
}

func (k *keyEvent) GetKeyCode() int {
	return k.keyCode
}

type IKeyPressedEvent interface {
	GetRepeatCount() int
	Init()
}

// KeyPressedEvent is the struct that implements the event, have the reference to the common keyEvent
type KeyPressedEvent struct {
	*keyEvent
	repeatCount int
}

func (k *KeyPressedEvent) GetRepeatCount() int {
	return k.repeatCount
}

func (k *KeyPressedEvent) ToString() string {
	return fmt.Sprintf("KeyPressedEvent: %d ( %d repeats)", k.keyCode, k.repeatCount)
}

func (k *KeyPressedEvent) Init() {
	k.eventType = KeyPressed
	k.eventCategory = Keyboard | Input
}

// IKeyReleasedEvent interface to implement KeyReleasedEvent
type IKeyReleasedEvent interface {
	Init()
}

type KeyReleasedEvent struct {
	*keyEvent
}

func (k *KeyReleasedEvent) ToString() string {
	return fmt.Sprintf("KeyReleasedEvent: %d", k.keyCode)
}

func (k *KeyReleasedEvent) Init() {
	k.eventType = KeyReleased
	k.eventCategory = Keyboard | Input
}
