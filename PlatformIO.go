package imgui

// #include "PlatformIOWrapper.h"
import "C"

import (
	"unsafe"
)

//WindowCallback functions all take a Viewport parameter
type WindowCallback func(v *Viewport)

//GetWindowBoolCallback function signature
type GetWindowBoolCallback func(v *Viewport) bool

//SetWindowVecCallback function signature
type SetWindowVecCallback func(v *Viewport, xpos, ypos float64)

//GetWindowVecCallback function signature
type GetWindowVecCallback func(v *Viewport) (float64, float64)

//SetWindowTitleCallback function signature
type SetWindowTitleCallback func(v *Viewport, title string)

// PlatformIO is
type PlatformIO struct {
	handle C.IggPlatformIO

	CreateWindow       WindowCallback
	DestroyWindow      WindowCallback
	ShowWindow         WindowCallback
	SetWindowPos       SetWindowVecCallback
	GetWindowPos       GetWindowVecCallback
	SetWindowSize      SetWindowVecCallback
	GetWindowSize      GetWindowVecCallback
	SetWindowFocus     WindowCallback
	GetWindowFocus     GetWindowBoolCallback
	GetWindowMinimized GetWindowBoolCallback
	SetWindowTitle     SetWindowTitleCallback
	RenderWindow       WindowCallback
	SwapBuffers        WindowCallback
}

var platform PlatformIO

// CurrentPlatformIO returns access to the ImGui communication struct for the currently active context.
func CurrentPlatformIO() PlatformIO {
	platform = PlatformIO{handle: C.iggGetCurrentPlatformIO()}
	return platform
}

// TODO: Store PlatformIO in pointer and restore here to call the registered callback
	//viewport := pointer.Restore(v).(*Viewport)
	// fmt.Println("Create Window Trigger hit")

//export goCreateWindowCallback
func goCreateWindowCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)
	// Forwarding callback into the registered Go function
	platform.CreateWindow(viewport)
}

//SetCallbackCreateWindow sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackCreateWindow(cbfun WindowCallback) WindowCallback {
	previous := io.CreateWindow

	io.CreateWindow = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackCreateWindow(io.handle)
	} else {
		C.iggSetCallbackCreateWindow(io.handle)
	}

	return previous
}

//export goDestroyWindowCallback
func goDestroyWindowCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.DestroyWindow(viewport)
}

//SetCallbackDestroyWindow sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackDestroyWindow(cbfun WindowCallback) WindowCallback {
	previous := io.DestroyWindow

	io.DestroyWindow = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackDestroyWindow(io.handle)
	} else {
		C.iggSetCallbackDestroyWindow(io.handle)
	}

	return previous
}

//export goShowWindowCallback
func goShowWindowCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.ShowWindow(viewport)
}

//SetCallbackShowWindow sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackShowWindow(cbfun WindowCallback) WindowCallback {
	previous := io.ShowWindow

	io.ShowWindow = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackShowWindow(io.handle)
	} else {
		C.iggSetCallbackShowWindow(io.handle)
	}

	return previous
}

//export goSetWindowPosCallback
func goSetWindowPosCallback(v unsafe.Pointer, xpos, ypos C.float) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.SetWindowPos(viewport, float64(xpos), float64(ypos))
}

//SetCallbackSetWindowPos sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackSetWindowPos(cbfun SetWindowVecCallback) SetWindowVecCallback {
	previous := io.SetWindowPos

	io.SetWindowPos = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackSetWindowPos(io.handle)
	} else {
		C.iggSetCallbackSetWindowPos(io.handle)
	}

	return previous
}

//export goGetWindowPosCallback
func goGetWindowPosCallback(v unsafe.Pointer, xpos, ypos *C.float) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	x, y := platform.GetWindowPos(viewport)
	*xpos = (C.float)(x)
	*ypos = (C.float)(y) 

}

//SetCallbackGetWindowPos sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackGetWindowPos(cbfun GetWindowVecCallback) GetWindowVecCallback {
	previous := io.GetWindowPos

	io.GetWindowPos = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackGetWindowPos(io.handle)
	} else {
		C.iggSetCallbackGetWindowPos(io.handle)
	}

	return previous
}

//export goSetWindowSizeCallback
func goSetWindowSizeCallback(v unsafe.Pointer, xpos, ypos C.float) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.SetWindowSize(viewport, float64(xpos), float64(ypos))
}

//SetCallbackSetWindowSize sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackSetWindowSize(cbfun SetWindowVecCallback) SetWindowVecCallback {
	previous := io.SetWindowSize

	io.SetWindowSize = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackSetWindowSize(io.handle)
	} else {
		C.iggSetCallbackSetWindowSize(io.handle)
	}

	return previous
}

//export goGetWindowSizeCallback
func goGetWindowSizeCallback(v unsafe.Pointer, xpos, ypos *C.float) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	x, y := platform.GetWindowSize(viewport)
	*xpos = (C.float)(x)
	*ypos = (C.float)(y)
}

//SetCallbackGetWindowSize sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackGetWindowSize(cbfun GetWindowVecCallback) GetWindowVecCallback {
	previous := io.GetWindowSize

	io.GetWindowSize = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackGetWindowSize(io.handle)
	} else {
		C.iggSetCallbackGetWindowSize(io.handle)
	}

	return previous
}

//export goSetWindowFocusCallback
func goSetWindowFocusCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.SetWindowFocus(viewport)
}

//SetCallbackSetWindowFocus sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackSetWindowFocus(cbfun WindowCallback) WindowCallback {
	previous := io.SetWindowFocus

	io.SetWindowFocus = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackSetWindowFocus(io.handle)
	} else {
		C.iggSetCallbackSetWindowFocus(io.handle)
	}

	return previous
}

//export goGetWindowFocusCallback
func goGetWindowFocusCallback(v unsafe.Pointer) C.int {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	focus := platform.GetWindowFocus(viewport)
	if focus {
		return 1
	}
	return 0
	//return (C.int)(focus)
}

//SetCallbackGetWindowFocus sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackGetWindowFocus(cbfun GetWindowBoolCallback) GetWindowBoolCallback {
	previous := io.GetWindowFocus

	io.GetWindowFocus = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackGetWindowFocus(io.handle)
	} else {
		C.iggSetCallbackGetWindowFocus(io.handle)
	}

	return previous
}

//export goGetWindowMinimizedCallback
func goGetWindowMinimizedCallback(v unsafe.Pointer) C.int {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	minimized := platform.GetWindowMinimized(viewport)
	if minimized {
		return 1
	}
	return 0
	//return (C.int)(minimized)
}

//SetCallbackGetWindowMinimized sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackGetWindowMinimized(cbfun GetWindowBoolCallback) GetWindowBoolCallback {
	previous := io.GetWindowMinimized

	io.GetWindowMinimized = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackGetWindowMinimized(io.handle)
	} else {
		C.iggSetCallbackGetWindowMinimized(io.handle)
	}

	return previous
}

//export goSetWindowTitleCallback
func goSetWindowTitleCallback(v unsafe.Pointer, title *C.char) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.SetWindowTitle(viewport, C.GoString(title))
}

//SetCallbackSetWindowTitle sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackSetWindowTitle(cbfun SetWindowTitleCallback) SetWindowTitleCallback {
	previous := io.SetWindowTitle

	io.SetWindowTitle = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackSetWindowTitle(io.handle)
	} else {
		C.iggSetCallbackSetWindowTitle(io.handle)
	}

	return previous
}

//export goRenderWindowCallback
func goRenderWindowCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.RenderWindow(viewport)
}

//SetCallbackRenderWindow sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackRenderWindow(cbfun WindowCallback) WindowCallback {
	previous := io.RenderWindow

	io.RenderWindow = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackRenderWindow(io.handle)
	} else {
		C.iggSetCallbackRenderWindow(io.handle)
	}

	return previous
}

//export goSwapBuffersCallback
func goSwapBuffersCallback(v unsafe.Pointer) {
	viewport := GetViewport(v)

	// Forwarding callback into the registered Go function
	platform.SwapBuffers(viewport)
}

//SetCallbackSwapBuffers sets the callback function in the PlatformIO struct and triggers
//the C setup to route the callback back into Go
func (io PlatformIO) SetCallbackSwapBuffers(cbfun WindowCallback) WindowCallback {
	previous := io.SwapBuffers

	io.SwapBuffers = cbfun

	// if the function is null, clear the callback
	if cbfun == nil {
		C.iggRemoveCallbackSwapBuffers(io.handle)
	} else {
		C.iggSetCallbackSwapBuffers(io.handle)
	}

	return previous
}
