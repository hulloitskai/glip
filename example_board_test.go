package glip_test

import (
	"fmt"
	"github.com/steven-xie/glip"
)

func ExampleBoard() {
	// Create a new glip.Board instance.
	b, _ := glip.NewBoard()

	// Write a string into the clipboard.
	b.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// the input string.
	out, _ := b.ReadString()
	fmt.Print(out)
	// Output: example string
}

func Example() {
	// Write a string into the clipboard.
	glip.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// teh input string.
	out, _ := glip.ReadString()
	fmt.Print(out)
	// Output: example string
}
