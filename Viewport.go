package imgui

// TODO: Define a Viewport Data identifier that's not platform (glfw) specific
// Wrap core ImGuiViewport functions in the wrapper
// create Viewport struct and expose the wrapper functions

// #include "ViewportWrapper.h"
import "C"

import (
	"unsafe"
)

// Flags stored in ImGuiViewport::Flags, giving indications to the platform back-ends.
const (
    ViewportFlags_None                     = 0
    ViewportFlags_NoDecoration             = 1 << 0   // Platform Window: Disable platform decorations: title bar, borders, etc. (generally set all windows, but if ImGuiConfigFlags_ViewportsDecoration is set we only set this on popups/tooltips)
    ViewportFlags_NoTaskBarIcon            = 1 << 1   // Platform Window: Disable platform task bar icon (generally set on popups/tooltips, or all windows if ImGuiConfigFlags_ViewportsNoTaskBarIcon is set)
    ViewportFlags_NoFocusOnAppearing       = 1 << 2   // Platform Window: Don't take focus when created.
    ViewportFlags_NoFocusOnClick           = 1 << 3   // Platform Window: Don't take focus when clicked on.
    ViewportFlags_NoInputs                 = 1 << 4   // Platform Window: Make mouse pass through so we can drag this window while peaking behind it.
    ViewportFlags_NoRendererClear          = 1 << 5   // Platform Window: Renderer doesn't need to clear the framebuffer ahead (because we will fill it entirely).
    ViewportFlags_TopMost                  = 1 << 6   // Platform Window: Display on top (for tooltips only).
    ViewportFlags_Minimized                = 1 << 7   // Platform Window: Window is minimized, can skip render. When minimized we tend to avoid using the viewport pos/size for clipping window or testing if they are contained in the viewport.
    ViewportFlags_NoAutoMerge              = 1 << 8   // Platform Window: Avoid merging this widow into another host window. This can only be set via ImGuiWindowClass viewport flags override (because we need to now ahead if we are going to create a viewport in the first place!).
    ViewportFlags_CanHostOtherWindows      = 1 << 9    // Main viewport: can host multiple imgui windows (secondary viewports are associated to a single window).
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

//GetMainViewport returns the main window / viewport that is owned by the application
func GetMainViewport() *Viewport {
	viewport := C.iggGetMainViewport()
	return NewViewport(unsafe.Pointer(viewport))
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

//SetPlatformHandle provides the ability to set a implementation specific window pointer
func (v Viewport) SetPlatformHandle(h unsafe.Pointer) {
	C.iggViewportSetPlatformHandle(v.handle, h)
}

//GetPosition returns the position of the ImguiViewport
func (v Viewport) GetPosition() (int, int) {
	x := C.iggViewportGetPositionX(v.handle)
	y := C.iggViewportGetPositionY(v.handle)
	return int(x), int(y)
}

//GetSize returns the position of the ImguiViewport
func (v Viewport) GetSize() (int, int) {
	x := C.iggViewportGetSizeX(v.handle)
	y := C.iggViewportGetSizeY(v.handle)
	return int(x), int(y)
}

//GetFlags from the backing Imgui viewport
func (v Viewport) GetFlags() int {
	return int(C.iggViewportGetFlags(v.handle))
}

//SetData allows an arbitrary data pointer to be stored alongside the Viewport
func (v *Viewport) SetData(d unsafe.Pointer) {
	v.data = d
}

//GetData returns the data pointer stored inside Viewport
func (v Viewport) GetData() unsafe.Pointer {
	return v.data
}
