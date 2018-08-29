// +build windows

package glip

import (
	"errors"
	"github.com/steven-xie/glip/portal"
	"io"
)

// WinClip represents the "clip" Windows program. This program is only capable
// of writing data to the clipboard.
type WinClip struct {
	*dynPortal
}

// NewWinClip creates a new WinClip instance, if "clip" can be found in its
// system path.
func NewWinClip() (wc *WinClip, err error) {
	if err = ensureCmdExists("clip"); err != nil {
		return nil, err
	}

	return &WinClip{dynPortal: newDynPortal("clip")}, nil
}

// Read is a non-functioning method for WinClip, which is a write-only program.
//
// It reads zero bytes, and always returns an error.
func (wc *WinClip) Read(dst []byte) (n int, err error) {
	return 0, errors.New("glip: WinClip is unable to read data from the " +
		"clipboard.")
}

// WriteTo is a non-functioning method for WinClip, which is a write-only
// program.
//
// It writes zero bytes to the provided io.Writer, and always returns an error.
func (wc *WinClip) WriteTo(w io.Writer) (n int64, err error) {
	return 0, errors.New("glip: WinClip is unable to read data from the " +
		"clipboard.")
}

// ReadPortal always returns nil.
func (wc *WinClip) ReadPortal() *portal.Portal {
	return nil
}

// WritePortal retuns WinClip's underlying portal.Portal used for writing data
// to the clipboard.
func (wc *WinClip) WritePortal() *portal.Portal {
	return wc.Portal
}
