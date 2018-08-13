package main

import (
	"fmt"
	"github.com/steven-xie/glip"
	"os"
)

func main() {
	var err error

	info, err := os.Stdin.Stat()
	if err != nil {
		errln("Failed to read standard input information:", err)
	}

	b, err := glip.NewBoard()
	if err != nil {
		errln("Failed to open system clipboard:", err)
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		if _, err := b.ReadFrom(os.Stdin); err != nil {
			errln("Failed to write to system clipboard:", err)
		}

		return
	}

	// No input was available...
	if _, err = b.WriteTo(os.Stdout); err != nil {
		errln("Failed to write clipboard contents to standard output:", err)
	}
}

func errln(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}
