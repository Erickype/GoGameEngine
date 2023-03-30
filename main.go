package main

import (
	"github.com/Erickype/GoGameEngine/API/Platform/Windows"
	"github.com/Erickype/GoGameEngine/Core"
	"github.com/Erickype/GoGameEngine/EngineTest/Layers"
)

func main() {
	window := Windows.CreateAbstractWindow("Test", 800, 600)

	app := Core.Application{}
	app.Construct(window)

	//Push the example layer
	exampleLayer := Layers.ExampleLayer{}
	exampleLayer.Construct("Example")
	iLayer := Core.ILayer(&exampleLayer)
	app.PushLayer(&iLayer)
	exampleLayer2 := Layers.ExampleLayer{}
	exampleLayer2.Construct("Example2")
	iLayer2 := Core.ILayer(&exampleLayer2)
	app.PushLayer(&iLayer2)

	app.Run()

	app.Destroy()
}
