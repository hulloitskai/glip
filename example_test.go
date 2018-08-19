// +build !windows

package glip_test

import (
	"fmt"
	"github.com/steven-xie/glip"
)

func Example() {
	// Write a string into the clipboard.
	glip.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// the input string.
	out, err := glip.ReadString()
	if err != nil {
		fmt.Printf("Encountered an error while reading from clipboard: %v", err)
	}

	fmt.Println(out)
	// Output: example string
}
