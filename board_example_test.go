package glip_test

import (
	"fmt"
	"github.com/steven-xie/glip"
)

func ExampleBoard() {
	// Make a new glip.Board instance.
	b, _ := glip.NewBoard()

	// Write a string into the clipboard.
	b.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// the input string.
	out, _ := b.ReadString()

	fmt.Println(out)
	// Output: example string
}
