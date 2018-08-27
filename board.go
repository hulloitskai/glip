package glip

import (
	"github.com/steven-xie/glip/portal"
	"io"
)

// Board is capable of both reading and writing to the system clipboard.
//
// It implements a lot of interfaces for reading and writing, but also exposes
// a method for accessing the underlying portal.Portal that represents its
// underlying program(s).
type Board interface {
	ReadPortal() *portal.Portal
	WritePortal() *portal.Portal
	ReadString() (s string, err error)
	WriteString(s string) (n int, err error)
	io.Reader
	io.Writer
	io.ReaderFrom
	io.WriterTo
}
