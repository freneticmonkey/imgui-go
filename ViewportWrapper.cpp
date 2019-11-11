#include "imguiWrappedHeader.h"
#include "ViewportWrapper.h"
#include "WrapperConverter.h"

void iggViewportSetPlatformHandle(IggViewport handle, void * window)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    v->PlatformHandle = window;
}

int iggViewportGetPositionX(IggViewport handle)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    return (int)v->Pos.x;
}

int iggViewportGetPositionY(IggViewport handle)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    return (int)v->Pos.y;
}