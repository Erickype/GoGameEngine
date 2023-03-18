package Events

import "fmt"

type IKeyEvent interface {
	GetKeyCode() int
}

// keyEvent common struct to key events
type keyEvent struct {
	*Event
	KeyCode int
}

func (k *keyEvent) GetKeyCode() int {
	return k.KeyCode
}

type IKeyPressedEvent interface {
	GetRepeatCount() int
}

// KeyPressedEvent is the struct that implements the event, have the reference to the common keyEvent
type KeyPressedEvent struct {
	*keyEvent
	RepeatCount int
}

func (k *KeyPressedEvent) GetRepeatCount() int {
	return k.RepeatCount
}

func (k *KeyPressedEvent) ToString() string {
	return fmt.Sprintf("KeyPressedEvent: %d ( %d repeats)", k.KeyCode, k.RepeatCount)
}

func (k *KeyPressedEvent) Init() {
	k.eventType = KeyPressed
	k.eventCategory = Keyboard | Input
}

// IKeyReleasedEvent interface to implement KeyReleasedEvent
type IKeyReleasedEvent interface{}

type KeyReleasedEvent struct {
	*keyEvent
}

func (k *KeyReleasedEvent) ToString() string {
	return fmt.Sprintf("KeyReleasedEvent: %d", k.KeyCode)
}

func (k *KeyReleasedEvent) Init() {
	k.eventType = KeyReleased
	k.eventCategory = Keyboard | Input
}

type KeyTypedEvent struct {
	*keyEvent
}

func (k *KeyTypedEvent) ToString() string {
	return fmt.Sprintf("KeyTypedEvent: %d", k.KeyCode)
}

func (k *KeyTypedEvent) Init() {
	k.eventType = KeyTyped
	k.eventCategory = Keyboard | Input
}
