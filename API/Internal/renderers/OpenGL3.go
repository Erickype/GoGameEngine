package renderers

import (
	_ "embed" // using embed for the shader sources
	"fmt"
	"github.com/AllenDang/cimgui-go"

	"github.com/Erickype/GoGameEngine/API/Internal/renderers/gl/v3.2-core/gl"
)

//go:embed gl-shader/main.vert
var unVersionedVertexShader string

//go:embed gl-shader/main.frag
var unVersionedFragmentShader string

// OpenGL3 implements a renderer based on github.com/go-gl/gl (v3.2-core).
type OpenGL3 struct {
	imGuiIO imgui.IO

	glslVersion            string
	fontTexture            uint32
	shaderHandle           uint32
	vertHandle             uint32
	fragHandle             uint32
	attribLocationTex      int32
	attribLocationProjMtx  int32
	attribLocationPosition int32
	attribLocationUV       int32
	attribLocationColor    int32
	vboHandle              uint32
	elementsHandle         uint32
}

// NewOpenGL3 attempts to initialize a renderer.
// An OpenGL context has to be established before calling this function.
func NewOpenGL3(io imgui.IO) (*OpenGL3, error) {
	err := gl.Init()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize OpenGL: %w", err)
	}

	renderer := &OpenGL3{
		imGuiIO:     io,
		glslVersion: "#version 150",
	}
	//renderer.createDeviceObjects()

	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasVtxOffset)

	return renderer, nil
}

// Dispose cleans up the resources.
func (renderer *OpenGL3) Dispose() {
	//renderer.invalidateDeviceObjects()
}
