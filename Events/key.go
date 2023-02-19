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
