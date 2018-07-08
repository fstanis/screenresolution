// Package screenresolution is used to retrieve the current screen resolution.
package screenresolution

// Get returns the current screen resolution of the primary display.
func Get() string {
	return get()
}
