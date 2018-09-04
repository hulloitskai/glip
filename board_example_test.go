package glip_test

import (
	"fmt"
	"github.com/steven-xie/glip"
)

func makeSafeBoard() glip.Board {
	b, _ := glip.NewBoard()

	// If b is a glip.Board that interacts with the X server, ensure that all
	// write operations wait for the X server to finish processing the data
	// before continuing.
	if xb, ok := b.(glip.XBoard); ok {
		xb.Opts().SafeWrites = true
	}

	return b
}

func ExampleBoard() {
	// Make a new glip.Board instance.
	b := makeSafeBoard()

	// Write a string into the clipboard.
	b.WriteString("example string")

	// Read the string back from the clipboard. We expect it to be the same as
	// the input string.
	out, _ := b.ReadString()

	fmt.Println(out)
	// Output: example string
}
