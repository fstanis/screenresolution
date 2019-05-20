// Package screenresolution is used to retrieve the current screen resolution.
package screenresolution

// Resolution represents the resolution of a single screen.
type Resolution struct {
	Width, Height int
}

// Get returns the current screen resolution of the primary display.
func Get() *Resolution {
	return get()
}
