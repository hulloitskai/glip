package glip

import (
	"fmt"
	"github.com/steven-xie/glip/portal"
	"io"
	"os/exec"
	"strings"
)

// Board is capable of both reading and writing to the system clipboard.
//
// It implements io.Reader, io.Writer, and io.WriterTo.
type Board struct {
	CopyPortal  *portal.Portal
	PastePortal *portal.Portal
}

// MakeBoard creates a Board by wrapping "copyCmd" and "pastCmd" into
// portal.Portals.
func MakeBoard(copyCmd, pasteCmd *exec.Cmd) *Board {
	return &Board{
		CopyPortal:  portal.MakeFrom(copyCmd),
		PastePortal: portal.MakeFrom(pasteCmd),
	}
}

//////////////////////////
// Error handling
//////////////////////////

// IsReadable determines if clipboard data can be read using Board.
func (b *Board) IsReadable() bool {
	return b.PastePortal != nil
}

// IsWriteable determines if clipboard data can be written using board.
func (b *Board) IsWriteable() bool {
	return b.CopyPortal != nil
}

//////////////////////////
// Reading from Board
//////////////////////////

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
		return "", fmt.Errorf("glip: could not write to buffer: %v", err)
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

//////////////////////////
// Writing to Board
//////////////////////////

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
