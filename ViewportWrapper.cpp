#include "imguiWrappedHeader.h"
#include "ViewportWrapper.h"
#include "WrapperConverter.h"

IggViewport iggGetMainViewport()
{
    return ImGui::GetMainViewport();
}

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

int iggViewportGetSizeX(IggViewport handle)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    return (int)v->Size.x;
}

int iggViewportGetSizeY(IggViewport handle)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    return (int)v->Size.y;
}

int iggViewportGetFlags(IggViewport handle)
{
    ImGuiViewport *v = reinterpret_cast<ImGuiViewport *>(handle);
    return (int)v->Flags;
}

