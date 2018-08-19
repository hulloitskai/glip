package glip

import (
	"fmt"
	"io"
)

// Write copies data from "p" into the system clipboard.
func (b *Board) Write(p []byte) (n int, err error) {
	if !b.IsWriteable() {
		return 0, ErrNotWriteable
	}

	if n, err = b.CopyPortal.Write(p); err != nil {
		return n, copyWriteErr(err)
	}
	return n, nil
}

// WriteString writes the provided string into the system clipboard.
func (b *Board) WriteString(s string) (n int, err error) {
	if !b.IsWriteable() {
		return 0, ErrNotWriteable
	}

	if n, err = b.CopyPortal.Write([]byte(s)); err != nil {
		return n, copyWriteErr(err)
	}
	return n, nil
}

// ReadFrom reads data from the provided io.Reader into the system clipboard.
func (b *Board) ReadFrom(r io.Reader) (n int64, err error) {
	if !b.IsWriteable() {
		return 0, ErrNotWriteable
	}

	if n, err = b.CopyPortal.ReadFrom(r); err != nil {
		return n, fmt.Errorf("glip: could not read from CopyPortal: %v", err)
	}
	return n, nil
}
