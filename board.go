package glip

import (
	"github.com/steven-xie/glip/portal"
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
		CopyPortal:  portal.MakeFrom(copyCmd),
		PastePortal: portal.MakeFrom(pasteCmd),
	}
}

// IsReadable determines if clipboard data can be read using Board.
func (b *Board) IsReadable() bool {
	return b.PastePortal != nil
}

// IsWriteable determines if clipboard data can be written using board.
func (b *Board) IsWriteable() bool {
	return b.CopyPortal != nil
}
