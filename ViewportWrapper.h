#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern void iggViewportSetPlatformHandle(IggViewport handle, void * window);
extern int iggViewportGetPositionX(IggViewport handle);
extern int iggViewportGetPositionY(IggViewport handle);

#ifdef __cplusplus
}
#endif
