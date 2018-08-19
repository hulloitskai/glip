package glip

import (
	"fmt"
	"io"
	"strings"
)

// Read reads data from the system clipboard into "p".
func (b *Board) Read(p []byte) (n int, err error) {
	if !b.IsReadable() {
		return 0, ErrNotReadable
	}

	if n, err = b.PastePortal.Read(p); err != nil {
		return n, fmt.Errorf("glip: could not read from PastePortal: %v", err)
	}
	return n, nil
}

// ReadString reads the contents of the system clipboard into a string.
func (b *Board) ReadString() (s string, err error) {
	if !b.IsReadable() {
		return "", ErrNotReadable
	}

	builder := new(strings.Builder)
	if _, err = b.WriteTo(builder); err != nil {
		return "", fmt.Errorf("glip: could not write to strings.Builder: %v", err)
	}

	return builder.String(), nil
}

// WriteTo writes data from the system clipboard into the provided io.Writer.
func (b *Board) WriteTo(w io.Writer) (n int64, err error) {
	if !b.IsReadable() {
		return 0, ErrNotReadable
	}

	if n, err = b.PastePortal.WriteTo(w); err != nil {
		return n, fmt.Errorf("glip: could not write to PastePortal: %v", err)
	}
	return n, nil
}
