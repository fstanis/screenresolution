package screenresolution

import "fmt"

func ExampleGetPrimary() {
	fmt.Println(GetPrimary().String())
	// Output: 1024x768
}
