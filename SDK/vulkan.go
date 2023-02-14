package SDK

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

func InitializeVulkan() {
	vulkan.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())

	err := vulkan.Init()
	if err != nil {
		panic(err)
	}
}
