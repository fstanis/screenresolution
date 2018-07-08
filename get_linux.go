package screenresolution

/*
#cgo LDFLAGS: -lX11

#include <X11/Xlib.h>
*/
import "C"
import "fmt"

func get() string {
	dy := C.XOpenDisplay(nil)
	if dy == nil {
		return ""
	}
	defer C.XCloseDisplay(dy)

	screen := C.XDefaultScreenOfDisplay(dy)
	width := int(C.XWidthOfScreen(screen))
	height := int(C.XHeightOfScreen(screen))
	if width == 0 || height == 0 {
		return ""
	}
	return fmt.Sprintf("%dx%d", width, height)
}
