// +build darwin

package glip

// NewBoard creates a new Board, if all the necessary system commands ar
// available.
func NewBoard() (b *Board, err error) {
	var (
		copyCBs  = []cmdBuilder{newCmdBuilder("pbcopy")}
		pasteCBs = []cmdBuilder{newCmdBuilder("pbpaste")}
	)
	return makeBoardFromPossibleCBs(copyCBs, pasteCBs)
}
