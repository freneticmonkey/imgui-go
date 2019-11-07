package imgui

// TODO: Define a Viewport Data identifier that's not platform (glfw) specific
// Wrap core ImGuiViewport functions in the wrapper
// create Viewport struct and expose the wrapper functions

// #include "ViewportWrapper.h"
import "C"

// Viewport is a new construct with the imgui 1.74 viewport feature
type Viewport struct {
	handle C.IggGuiViewport
}
