package Opengl

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Context struct {
	windowHandle *glfw.Window
}

func (c *Context) Init() {
	c.windowHandle.MakeContextCurrent()
	glfw.SwapInterval(1)
}

func (c *Context) SwapBuffers() {
	c.windowHandle.SwapBuffers()
}

func NewOpenglContext(windowHandle *glfw.Window) *Context {
	context := Context{}
	context.windowHandle = windowHandle
	return &context
}
