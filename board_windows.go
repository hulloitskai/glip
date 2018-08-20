// +build windows

package glip

// NewBoard creates a new Board.
func NewBoard() (b *Board, err error) {
	var (
		copyCBs = []cmdBuilder{
			newCmdBuilder("clip"),
			newCmdBuilder("powershell", "-c", "Set-Clipboard"),
		}
		pasteCBs = []cmdBuilder{
			newCmdBuilder("powershell", "-c", "Get-Clipboard"),
			newCmdBuilder("paste"),
		}
	)
	return makeBoardFromPossibleCBs(copyCBs, pasteCBs)
}
