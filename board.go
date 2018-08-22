package glip

import (
	"fmt"
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

// makeBoard makes a Board, given a set of cmdBuilders for making copy and paste
// exec.Cmds. It will use the first cmdBuilder that is valid.
func makeBoardFromPossibleCBs(copyCBs, pasteCBs []cmdBuilder) (
	b *Board, err error,
) {
	var (
		copyCmd, pasteCmd *exec.Cmd
		index             int
	)

	if index, err = autoselectCB(copyCBs); index != -1 {
		copyCmd = copyCBs[index].build()
	}
	if index, err = autoselectCB(pasteCBs); index != -1 {
		pasteCmd = pasteCBs[index].build()
	}
	if err != nil {
		return nil, fmt.Errorf("glip: error while selecting command: %v", err)
	}

	return MakeBoard(copyCmd, pasteCmd), nil
}
