package glip

import (
	"errors"
	"fmt"
)

var (
	// ErrNotWriteable is an error that occurs when a system does not have the
	// ability to write data to the clipboard.
	ErrNotWriteable = errors.New("glip: no writeable copying interface is " +
		"available")

	// ErrNotReadable is an error that occurs when a system does not have the
	// ability to read data from the clipboard.
	ErrNotReadable = errors.New("glip: no readable pasting interface is " +
		"available")
)

func copyWriteErr(err error) error {
	return fmt.Errorf("glip: could not write to CopyPortal: %v", err)
}
