package screenresolution

import (
	"github.com/lxn/win"
)

func getPrimary() *Resolution {
	width := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	height := int(win.GetSystemMetrics(win.SM_CYSCREEN))
	if width == 0 || height == 0 {
		return nil
	}
	return &Resolution{width, height}
}
