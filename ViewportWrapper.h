#pragma once

#include "imguiWrapperTypes.h"

#ifdef __cplusplus
extern "C"
{
#endif

extern IggViewport iggGetMainViewport();

extern void iggViewportSetPlatformHandle(IggViewport handle, void * window);
extern int iggViewportGetPositionX(IggViewport handle);
extern int iggViewportGetPositionY(IggViewport handle);

extern int iggViewportGetSizeX(IggViewport handle);
extern int iggViewportGetSizeY(IggViewport handle);

extern int iggViewportGetFlags(IggViewport handle);

extern void iggViewportPlatformRequestClose(IggViewport handle);
extern void iggViewportPlatformRequestMove(IggViewport handle);
extern void iggViewportPlatformRequestResize(IggViewport handle);


#ifdef __cplusplus
}
#endif
