package imgui

// TODO: Define a Viewport Data identifier that's not platform (glfw) specific
// Wrap core ImGuiViewport functions in the wrapper
// create Viewport struct and expose the wrapper functions

// #include "ViewportWrapper.h"
import "C"

import (
	"unsafe"
)

// Viewport is a new construct with the imgui 1.74 viewport feature
type Viewport struct {
	handle C.IggViewport
	data unsafe.Pointer
}

// Store for viewports
var viewports = map[unsafe.Pointer]*Viewport{}

//NewViewport creates a viewport using the pointer parameter
func NewViewport(v unsafe.Pointer) *Viewport {
	viewport := &Viewport{
		handle: (C.IggViewport)(v),
	}
	viewports[v] = viewport
	return viewport
}

//GetViewport returns a Viewport wrapper given the *ImguiViewport
func GetViewport(v unsafe.Pointer) *Viewport {
	viewport, ok := viewports[v]
	if ok {
		return viewport
	}
	// TODO: Error out here
	return nil
}

//SetPlaformHandle provides the ability to set a implementation specific window pointer
func (v Viewport) SetPlaformHandle(h unsafe.Pointer) {
	C.iggViewportSetPlatformHandle(v.handle, h)
}

//GetPosition returns the position of the ImguiViewport
func (v Viewport) GetPosition() (int, int) {
	x := C.iggViewportGetPositionX(v.handle)
	y := C.iggViewportGetPositionY(v.handle)
	return int(x), int(y)
}

//SetData allows an arbitrary data pointer to be stored alongside the Viewport
func (v *Viewport) SetData(d unsafe.Pointer) {
	v.data = d
}

//GetData returns the data pointer stored inside Viewport
func (v Viewport) GetData() unsafe.Pointer {
	return v.data
}
