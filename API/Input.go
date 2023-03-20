package API

type IInput interface {
	IsKeyPressed(keyCode int) bool
	isKeyPressed(keyCode int) bool
}

var inputInstance *IInput

type Input struct {
}

func (i *Input) IsKeyPressed(keyCode int) bool {
	return (*inputInstance).isKeyPressed(keyCode)
}

func (i *Input) isKeyPressed(_ int) bool {
	return true
}

func SetInputInstance(instance *IInput) {
	inputInstance = instance
}

func GetInputInstance() *IInput {
	return inputInstance
}
