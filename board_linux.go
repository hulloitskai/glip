// +build linux openbsd netbsd solaris

package glip

// NewBoard creates a new Board, if all the necessary system commands ar
// available.
func NewBoard() (b *Board, err error) {
	var (
		copyCBs = []cmdBuilder{
			newCmdBuilder("xclip"),
			newCmdBuilder("xsel"),
		}
		pasteCBs = []cmdBuilder{
			newCmdBuilder("xclip", "-o"),
			newCmdBuilder("xsel"),
		}
	)
	return makeBoardFromPossibleCBs(copyCBs, pasteCBs)
}

// NewBoardWith creates a new board targeting a specific system clipboard.
func NewBoardWith(clipboard string) (b *Board, err error) {
	xselCB := newCmdBuilder("xsel", "--", clipboard)
	var (
		copyCBs = []cmdBuilder{
			newCmdBuilder("xclip", "-sel", clipboard),
			xselCB,
		}
		pasteCBs = []cmdBuilder{
			newCmdBuilder("xclip", "--sel", clipboard, "-o"),
			xselCB,
		}
	)
	return makeBoardFromPossibleCBs(copyCBs, pasteCBs)
}
