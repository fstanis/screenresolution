package screenresolution

/*
#cgo LDFLAGS: -lX11

#include <X11/Xlib.h>
*/
import "C"

func getPrimary() *Resolution {
	dy := C.XOpenDisplay(nil)
	if dy == nil {
		return nil
	}
	defer C.XCloseDisplay(dy)

	screen := C.XDefaultScreenOfDisplay(dy)
	width := int(C.XWidthOfScreen(screen))
	height := int(C.XHeightOfScreen(screen))
	if width == 0 || height == 0 {
		return nil
	}
	return &Resolution{width, height}
}
