// Package screenresolution is used to retrieve the current screen resolution.
package screenresolution

import "fmt"

// Resolution represents the resolution of a single screen.
type Resolution struct {
	Width, Height int
}

func (r *Resolution) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}

// GetPrimary returns the current screen resolution of the primary display.
func GetPrimary() *Resolution {
	return getPrimary()
}
