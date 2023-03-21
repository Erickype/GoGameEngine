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

func (i *Input) IsKeyPressed(keyCode int, windowPtr unsafe.Pointer) bool {

	window := (*glfw.Window)(windowPtr)
	state := window.GetKey(glfw.Key(keyCode))

	return state == glfw.Press || state == glfw.Repeat
}
