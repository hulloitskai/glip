package glip

import (
	"bytes"
	"github.com/steven-xie/glip/portal"
	"io"
	"os/exec"
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
		CopyPortal:  portal.New(copyCmd, "copy"),
		PastePortal: portal.New(pasteCmd, "paste"),
	}
}

// Write copies data from "p" into the system clipboard.
func (b *Board) Write(p []byte) (n int, err error) {
	return b.CopyPortal.Write(p)
}

// WriteString writes the provided string into the system clipboard.
func (b *Board) WriteString(s string) (n int, err error) {
	return b.CopyPortal.Write([]byte(s))
}

// Read reads data from the system clipboard into "p".
func (b *Board) Read(p []byte) (n int, err error) {
	return b.PastePortal.Read(p)
}

// ReadString reads the contents of the system clipboard into a string.
func (b *Board) ReadString() (s string, err error) {
	buf := new(bytes.Buffer)
	if _, err = b.WriteTo(buf); err != nil {
		return "", err
	}
	return buf.String(), err
}

// WriteTo writes data from the system clipboard into the provided io.Writer.
func (b *Board) WriteTo(w io.Writer) (n int64, err error) {
	return b.PastePortal.WriteTo(w)
}
