package API

import (
	"unsafe"
)

type IInput interface {
	IsKeyPressed(keyCode int, window unsafe.Pointer) bool
}

var inputInstance *IInput

func SetInputInstance(instance *IInput) {
	inputInstance = instance
}

func _() *IInput {
	return inputInstance
}
