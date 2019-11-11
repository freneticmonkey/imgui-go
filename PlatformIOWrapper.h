#pragma once

#include "imguiWrapperTypes.h"


#ifdef __cplusplus
extern "C"
{
#endif

extern IggPlatformIO iggGetCurrentPlatformIO(void);

extern void iggCreateWindowCallback(IggViewport *v);
extern void iggRemoveCallbackCreateWindow(IggPlatformIO handle);
extern void iggSetCallbackCreateWindow(IggPlatformIO handle);

extern void iggDestroyWindowCallback(IggViewport *v);
extern void iggRemoveCallbackDestroyWindow(IggPlatformIO handle);
extern void iggSetCallbackDestroyWindow(IggPlatformIO handle);

extern void iggShowWindowCallback(IggViewport *v);
extern void iggRemoveCallbackShowWindow(IggPlatformIO handle);
extern void iggSetCallbackShowWindow(IggPlatformIO handle);

extern void iggSetWindowPosCallback(IggViewport *v, IggVec2 pos);
extern void iggRemoveCallbackSetWindowPos(IggPlatformIO handle);
extern void iggSetCallbackSetWindowPos(IggPlatformIO handle);

extern IggVec2 iggGetWindowPosCallback(IggViewport *v);
extern void iggRemoveCallbackGetWindowPos(IggPlatformIO handle);
extern void iggSetCallbackGetWindowPos(IggPlatformIO handle);

extern void iggSetWindowSizeCallback(IggViewport *v, IggVec2 size);
extern void iggRemoveCallbackSetWindowSize(IggPlatformIO handle);
extern void iggSetCallbackSetWindowSize(IggPlatformIO handle);

extern IggVec2 iggGetWindowSizeCallback(IggViewport *v);
extern void iggRemoveCallbackGetWindowSize(IggPlatformIO handle);
extern void iggSetCallbackGetWindowSize(IggPlatformIO handle);

extern void iggSetWindowFocusCallback(IggViewport *v);
extern void iggRemoveCallbackSetWindowFocus(IggPlatformIO handle);
extern void iggSetCallbackSetWindowFocus(IggPlatformIO handle);

extern IggBool iggGetWindowFocusCallback(IggViewport *v);
extern void iggRemoveCallbackGetWindowFocus(IggPlatformIO handle);
extern void iggSetCallbackGetWindowFocus(IggPlatformIO handle);

extern IggBool iggGetWindowMinimizedCallback(IggViewport *v);
extern void iggRemoveCallbackGetWindowMinimized(IggPlatformIO handle);
extern void iggSetCallbackGetWindowMinimized(IggPlatformIO handle);

extern void iggSetWindowTitleCallback(IggViewport *v, const char * title);
extern void iggRemoveCallbackSetWindowTitle(IggPlatformIO handle);
extern void iggSetCallbackSetWindowTitle(IggPlatformIO handle);

extern void iggRenderWindowCallback(IggViewport *v);
extern void iggRemoveCallbackRenderWindow(IggPlatformIO handle);
extern void iggSetCallbackRenderWindow(IggPlatformIO handle);

extern void iggSwapBuffersCallback(IggViewport *v);
extern void iggRemoveCallbackSwapBuffers(IggPlatformIO handle);
extern void iggSetCallbackSwapBuffers(IggPlatformIO handle);

extern void iggClearMonitors(IggPlatformIO handle);
extern void iggCreateMonitor(IggPlatformIO handle, int posX, int posY, int sizeX, int sizeY);

#ifdef __cplusplus
}
#endif