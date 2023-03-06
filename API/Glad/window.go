package Glad

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

type WinOption func()

func NewOGLWindow(width, height int, title string, opts ...WinOption) *glfw.Window {
	// Initialize OpenGL
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize GLFW", err)
	}
	for _, opt := range opts {
		opt()
	}
	win, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln("Failed to create window", err)
	}
	win.MakeContextCurrent()

	return win
}

func CoreProfile(v bool) WinOption {
	return func() {
		if v {
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		} else {
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCompatProfile)
		}
	}
}

func Resizable(v bool) WinOption {
	return func() {
		glfw.WindowHint(glfw.Resizable, glfwTF(v))
	}
}

func glfwTF(v bool) int {
	if v {
		return glfw.True
	}
	return glfw.False
}
func ContextVersion(maj, min int) WinOption {
	return func() {
		glfw.WindowHint(glfw.ContextVersionMajor, maj)
		glfw.WindowHint(glfw.ContextVersionMinor, min)
	}
}
