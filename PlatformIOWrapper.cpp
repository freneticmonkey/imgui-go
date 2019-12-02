#include "PlatformIOWrapper.h"
#include "imguiWrappedHeader.h"
#include "WrapperConverter.h"

#include "_cgo_export.h"

IggPlatformIO iggGetCurrentPlatformIO()
{
   return reinterpret_cast<IggPlatformIO>(&ImGui::GetPlatformIO());
}

void iggCreateWindowCallback(ImGuiViewport* v)
{
    // ImGuiViewport *view = reinterpret_cast<ImGuiViewport *>(v);
    if (v == NULL)
        return;
    goCreateWindowCallback(v);
}

void iggRemoveCallbackCreateWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_CreateWindow = nullptr;
}

void iggSetCallbackCreateWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_CreateWindow = iggCreateWindowCallback;
}

void iggDestroyWindowCallback(ImGuiViewport *v)
{
    goDestroyWindowCallback(v);
}

void iggRemoveCallbackDestroyWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_DestroyWindow = nullptr;
}

void iggSetCallbackDestroyWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_DestroyWindow = iggDestroyWindowCallback;
}

void iggShowWindowCallback(ImGuiViewport *v)
{
    goShowWindowCallback(v);
}

void iggRemoveCallbackShowWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_ShowWindow = nullptr;
}

void iggSetCallbackShowWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_ShowWindow = iggShowWindowCallback;
}

void iggSetWindowPosCallback(ImGuiViewport *v, ImVec2 pos)
{
    goSetWindowPosCallback(v, pos.x, pos.y);
}

void iggRemoveCallbackSetWindowPos(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowPos = nullptr;
}

void iggSetCallbackSetWindowPos(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowPos = iggSetWindowPosCallback;
}

ImVec2 iggGetWindowPosCallback(ImGuiViewport *v)
{
    float x, y;
    goGetWindowPosCallback(v, &x, &y);
    return ImVec2(x,y);
}

void iggRemoveCallbackGetWindowPos(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowPos = nullptr;
}

void iggSetCallbackGetWindowPos(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowPos = iggGetWindowPosCallback;
}

void iggSetWindowSizeCallback(ImGuiViewport *v, ImVec2 size)
{
    goSetWindowSizeCallback(v, size.x, size.y);
}

void iggRemoveCallbackSetWindowSize(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowSize = nullptr;
}

void iggSetCallbackSetWindowSize(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowSize = iggSetWindowSizeCallback;
}

ImVec2 iggGetWindowSizeCallback(ImGuiViewport *v)
{
    float x, y;
    goGetWindowSizeCallback(v, &x, &y);
    return ImVec2(x, y);
}

void iggRemoveCallbackGetWindowSize(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowSize = nullptr;
}

void iggSetCallbackGetWindowSize(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowSize = iggGetWindowSizeCallback;
}

void iggSetWindowFocusCallback(ImGuiViewport *v)
{
    goSetWindowFocusCallback(v);
}

void iggRemoveCallbackSetWindowFocus(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowFocus = nullptr;
}

void iggSetCallbackSetWindowFocus(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowFocus = iggSetWindowFocusCallback;
}

bool iggGetWindowFocusCallback(ImGuiViewport *v)
{
    return goGetWindowFocusCallback(v);
}

void iggRemoveCallbackGetWindowFocus(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowFocus = nullptr;
}

void iggSetCallbackGetWindowFocus(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowFocus = iggGetWindowFocusCallback;
}

bool iggGetWindowMinimizedCallback(ImGuiViewport *v)
{
    return goGetWindowMinimizedCallback(v);
}

void iggRemoveCallbackGetWindowMinimized(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowMinimized = nullptr;
}

void iggSetCallbackGetWindowMinimized(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_GetWindowMinimized = iggGetWindowMinimizedCallback;
}

void iggSetWindowTitleCallback(ImGuiViewport *v, const char * title)
{
    goSetWindowTitleCallback(v, const_cast<char *>(title));
}

void iggRemoveCallbackSetWindowTitle(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowTitle = nullptr;
}

void iggSetCallbackSetWindowTitle(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SetWindowTitle = iggSetWindowTitleCallback;
}

void iggRenderWindowCallback(ImGuiViewport *v, void*)
{
    goRenderWindowCallback(v);
}

void iggRemoveCallbackRenderWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_RenderWindow = nullptr;
}

void iggSetCallbackRenderWindow(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_RenderWindow = iggRenderWindowCallback;
}

void iggSwapBuffersCallback(ImGuiViewport *v, void*)
{
    goSwapBuffersCallback(v);
}

void iggRemoveCallbackSwapBuffers(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SwapBuffers = nullptr;
}

void iggSetCallbackSwapBuffers(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Platform_SwapBuffers = iggSwapBuffersCallback;
}

void iggClearMonitors(IggPlatformIO handle)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    io->Monitors.resize(0);
}

void iggCreateMonitor(IggPlatformIO handle, int posX, int posY, int sizeW, int sizeH)
{
    ImGuiPlatformIO *io = reinterpret_cast<ImGuiPlatformIO *>(handle);
    ImGuiPlatformMonitor monitor;
    monitor.MainPos = monitor.WorkPos = ImVec2((float)posX, (float)posY);
    monitor.MainSize = monitor.WorkSize = ImVec2((float)sizeW, (float)sizeH);
    io->Monitors.push_back(monitor);
}
