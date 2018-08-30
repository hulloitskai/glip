package glip

import (
	"github.com/steven-xie/glip/portal"
	"io"
)

// dualBoard is a Board that is composed of two inner dynPortals; one of which
// that is responsible for reading to the clipboard, and another that is
// responsible for writing to the clipboard.
type dualBoard struct {
	Writer *dynPortal
	Reader *dynPortal
}

// newDualBoard makes a new DualBoard instance.
func newDualBoard(writer, reader *dynPortal) *dualBoard {
	return &dualBoard{Writer: writer, Reader: reader}
}

// Write writes data into the system clipboard.
func (db *dualBoard) Write(data []byte) (n int, err error) {
	return db.Writer.Write(data)
}

// ReadFrom reads data from an io.Reader into the system clipboard.
func (db *dualBoard) ReadFrom(r io.Reader) (n int64, err error) {
	return db.Writer.ReadFrom(r)
}

// WriteString writes a string into the system clipboard.
func (db *dualBoard) WriteString(s string) (n int, err error) {
	return db.Writer.WriteString(s)
}

// Read reads len(dst) bytes from the system clipboard into dst.
func (db *dualBoard) Read(dst []byte) (n int, err error) {
	return db.Reader.Read(dst)
}

// WriteTo writes data from the system clipboard into an io.Writer.
func (db *dualBoard) WriteTo(w io.Writer) (n int64, err error) {
	return db.Reader.WriteTo(w)
}

// ReadString reads data from the system clipboard as a string.
func (db *dualBoard) ReadString() (s string, err error) {
	return db.Reader.ReadString()
}

// ReadPortal exposes a portal.Portal used for reading from the clipboard.
func (db *dualBoard) ReadPortal() *portal.Portal {
	return db.Reader.Portal
}

// WritePortal exposes a portal.Portal used for writing to the clipboard.
func (db *dualBoard) WritePortal() *portal.Portal {
	return db.Writer.Portal
}
