package screenresolution

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <AppKit/NSScreen.h>

CGSize getPrimaryScreenSize() {
    NSArray* screens = [NSScreen screens];
    if ([screens count] < 1) {
        CGSize empty;
        return empty;
    }
	return [[screens objectAtIndex: 0] frame].size;
}
*/
import "C"

func getPrimary() *Resolution {
	rect := C.getPrimaryScreenSize()
	if rect.width == 0 || rect.height == 0 {
		return nil
	}
	return &Resolution{int(rect.width), int(rect.height)}
}
