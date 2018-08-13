package main

import (
	"io"
	"fmt"
	"bufio"
	"github.com/steven-xie/glip"
	"os"
)


func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		errln("Failed to read standard input information:", err)
	}

	b, err := glip.NewBoard()
	if err != nil {
		errln("Failed to open system clipboard:", err)
	}

	if (info.Mode() & os.ModeCharDevice) == 0 {
		r := bufio.NewReader(os.Stdin)

		if _, err = r.WriteTo(b); err != nil {
			errln("Failed to write to system clipboard:", err)
		}
	}
	s := bufio.NewScanner(os.Stdin)

	if !s.Scan() &&

	// See if there's any data to be read from os.Stdin.
	data, err := r.Peek(1)
	if err = bufio.ErrFinalToken
}

func errln(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}
