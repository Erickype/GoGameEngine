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
	renderer.createDeviceObjects()

	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasVtxOffset)

	return renderer, nil
}

// Dispose cleans up the resources.
func (renderer *OpenGL3) Dispose() {
	//renderer.invalidateDeviceObjects()
}

func (renderer *OpenGL3) createDeviceObjects() {
	// Backup GL state
	var lastTexture int32
	var lastArrayBuffer int32
	var lastVertexArray int32
	gl.GetIntegerv(gl.TEXTURE_BINDING_2D, &lastTexture)
	gl.GetIntegerv(gl.ARRAY_BUFFER_BINDING, &lastArrayBuffer)
	gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &lastVertexArray)

	vertexShader := renderer.glslVersion + "\n" + unVersionedVertexShader

	fragmentShader := renderer.glslVersion + "\n" + unVersionedFragmentShader

	renderer.shaderHandle = gl.CreateProgram()
	renderer.vertHandle = gl.CreateShader(gl.VERTEX_SHADER)
	renderer.fragHandle = gl.CreateShader(gl.FRAGMENT_SHADER)

	glShaderSource := func(handle uint32, source string) {
		csource, free := gl.Strs(source + "\x00")
		defer free()

		gl.ShaderSource(handle, 1, csource, nil)
	}

	glShaderSource(renderer.vertHandle, vertexShader)
	glShaderSource(renderer.fragHandle, fragmentShader)
	gl.CompileShader(renderer.vertHandle)
	gl.CompileShader(renderer.fragHandle)
	gl.AttachShader(renderer.shaderHandle, renderer.vertHandle)
	gl.AttachShader(renderer.shaderHandle, renderer.fragHandle)
	gl.LinkProgram(renderer.shaderHandle)

	renderer.attribLocationTex = gl.GetUniformLocation(renderer.shaderHandle, gl.Str("Texture"+"\x00"))
	renderer.attribLocationProjMtx = gl.GetUniformLocation(renderer.shaderHandle, gl.Str("ProjMtx"+"\x00"))
	renderer.attribLocationPosition = gl.GetAttribLocation(renderer.shaderHandle, gl.Str("Position"+"\x00"))
	renderer.attribLocationUV = gl.GetAttribLocation(renderer.shaderHandle, gl.Str("UV"+"\x00"))
	renderer.attribLocationColor = gl.GetAttribLocation(renderer.shaderHandle, gl.Str("Color"+"\x00"))

	gl.GenBuffers(1, &renderer.vboHandle)
	gl.GenBuffers(1, &renderer.elementsHandle)

	//renderer.createFontsTexture()

	// Restore modified GL state
	gl.BindTexture(gl.TEXTURE_2D, uint32(lastTexture))
	gl.BindBuffer(gl.ARRAY_BUFFER, uint32(lastArrayBuffer))
	gl.BindVertexArray(uint32(lastVertexArray))
}
