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
	out, _ := glip.ReadString()

	fmt.Println(out)
}
