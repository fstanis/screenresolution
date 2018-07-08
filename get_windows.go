package screenresolution

import (
	"fmt"

	"github.com/lxn/win"
)

func get() string {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	if width == 0 || height == 0 {
		return ""
	}
	return fmt.Sprintf("%dx%d", width, height)
}
