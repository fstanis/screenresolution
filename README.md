# screenresolution

screenresolution is a simple, cross-platform Go package library used to detect
the current screen resolution. It follows the UNIX philosophy of doing one thing
and doing it well.

It supports Windows, Linux and MaxOS.

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/fstanis/screenresolution"
)

func main() {
	resolution := screenresolution.GetPrimary()
	if resolution == "" {
		fmt.Println("failed to get screen resolution")
		os.Exit(1)
	}
	fmt.Println(resolution)
}
```

## Limitations

screenresolution can only get the resolution from the default (primary) display.
A future update will add support for getting the resolutions of all attached
displays.

## Dependencies

On Linux, `libx11-dev` is required.

```bash
sudo apt-get install libx11-dev
```
