package Input

import (
	"unsafe"
)

type IInput interface {
	IsKeyPressed(keyCode int, windowPointer unsafe.Pointer) bool
	IsMouseButtonPressed(button int, windowPointer unsafe.Pointer) bool
	GetMousePosition(windowPointer unsafe.Pointer) (float64, float64)
	GetMouseX(windowPointer unsafe.Pointer) float64
	GetMouseY(windowPointer unsafe.Pointer) float64
}

var inputInstance *IInput

func SetInputInstance(instance *IInput) {
	inputInstance = instance
}

func GetInputInstance() *IInput {
	return inputInstance
}
