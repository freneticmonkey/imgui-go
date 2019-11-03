package imgui

// TODO: Define a Viewport Data identifier that's not platform (glfw) specific
// Wrap core ImGuiViewport functions in the wrapper
// create Viewport struct and expose the wrapper functions

import "C"

type Viewport struct {
	handle C.IggGuiViewport
}
