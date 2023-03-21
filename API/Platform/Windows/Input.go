package Windows

import (
	abstractInput "github.com/Erickype/GoGameEngine/API/Input"
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

func init() {
	input := &Input{}
	iInput := abstractInput.IInput(input)
	abstractInput.SetInputInstance(&iInput)
}

type Input struct{}

func (i *Input) IsMouseButtonPressed(button int, windowPointer unsafe.Pointer) bool {
	window := (*glfw.Window)(windowPointer)
	state := window.GetMouseButton(glfw.MouseButton(button))
	return state == glfw.Press
}

func (i *Input) GetMousePosition(windowPointer unsafe.Pointer) (float64, float64) {
	window := (*glfw.Window)(windowPointer)
	return window.GetCursorPos()
}

func (i *Input) GetMouseX(windowPointer unsafe.Pointer) float64 {
	window := (*glfw.Window)(windowPointer)
	xPos, _ := window.GetCursorPos()

	return xPos
}

func (i *Input) GetMouseY(windowPointer unsafe.Pointer) float64 {
	window := (*glfw.Window)(windowPointer)
	_, yPos := window.GetCursorPos()

	return yPos
}

func (i *Input) IsKeyPressed(keyCode int, windowPointer unsafe.Pointer) bool {
	window := (*glfw.Window)(windowPointer)
	state := window.GetKey(glfw.Key(keyCode))

	return state == glfw.Press || state == glfw.Repeat
}
