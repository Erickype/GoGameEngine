package Windows

import (
	"github.com/Erickype/GoGameEngine/API"
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

func init() {
	input := &Input{}
	iInput := API.IInput(input)
	API.SetInputInstance(&iInput)
}

type Input struct {
	*API.Input
}

func (i *Input) isKeyPressed(keyCode int, windowPtr unsafe.Pointer) bool {

	window := (*glfw.Window)(windowPtr)
	state := window.GetKey(glfw.Key(keyCode))

	return state == glfw.Press || state == glfw.Repeat
}
