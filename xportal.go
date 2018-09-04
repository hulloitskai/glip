package glip

import (
	"io"
	"time"
)

// XPortal is a dynPortal that can interact with X server selections.
type XPortal struct {
	// Selection is XPortal's target X selection.
	Selection XSelection

	// SafeWrites causes XPortal write methods to wait a short while for the
	// underlying X server to finish processing the write operation before
	// continuing.
	SafeWrites bool

	*dynPortal
}

// Write writes len(p) bytes into the XPortal.
func (xp *XPortal) Write(p []byte) (n int, err error) {
	if xp.SafeWrites {
		defer xp.breathe()
	}
	return xp.dynPortal.Write(p)
}

// WriteString writes a string s into the XPortal.
func (xp *XPortal) WriteString(s string) (n int, err error) {
	if xp.SafeWrites {
		defer xp.breathe()
	}
	return xp.dynPortal.WriteString(s)
}

// ReadFrom reads data from r into the XPortal.
func (xp *XPortal) ReadFrom(r io.Reader) (n int64, err error) {
	if xp.SafeWrites {
		defer xp.breathe()
	}
	return xp.dynPortal.ReadFrom(r)
}

// breathe pauses the current goroutine for a short while in order to give
// the X server some time to process a clipboard write operation.
func (*XPortal) breathe() {
	time.Sleep(4 * time.Millisecond)
}
