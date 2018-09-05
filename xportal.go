package glip

import (
	"io"
	"time"
)

// Xopts are a set of common options related to interacting with X selections.
type Xopts struct {
	// Selection is target X selection to be used.
	Selection XSelection

	// SafeWrites causes write methods to wait a short while for the
	// underlying X server to finish processing the write operation before
	// continuing.
	SafeWrites bool
}

// xPortal is a dynPortal that can interact with X server selections.
type xPortal struct {
	Xopts
	*dynPortal
	pauseLength int
}

// Write writes len(p) bytes into the xPortal.
//
// If xp.SafeWrites is enabled, Write will cause the current goroutine to sleep
// for xp.pauseLength seconds after returning.
func (xp *xPortal) Write(p []byte) (n int, err error) {
	if xp.SafeWrites {
		defer xp.sleep()
	}
	return xp.dynPortal.Write(p)
}

// WriteString writes a string s into the xPortal.
//
// If xp.SafeWrites is enabled, WriteString will cause the current goroutine to
// sleep for xp.pauseLength seconds after returning.
func (xp *xPortal) WriteString(s string) (n int, err error) {
	if xp.SafeWrites {
		defer xp.sleep()
	}
	return xp.dynPortal.WriteString(s)
}

// ReadFrom reads data from r into the xPortal.
//
// If xp.SafeWrites is enabled, ReadFrom will cause the current goroutine to
// sleep for xp.pauseLength seconds after returning.
func (xp *xPortal) ReadFrom(r io.Reader) (n int64, err error) {
	if xp.SafeWrites {
		defer xp.sleep()
	}
	return xp.dynPortal.ReadFrom(r)
}

// Opts exposes xPortal's X-server-related options.
func (xp *xPortal) Opts() *Xopts {
	return &xp.Xopts
}

func (xp *xPortal) sleep() {
	time.Sleep(time.Duration(xp.pauseLength) * time.Millisecond)
}
