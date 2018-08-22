// +build windows

package glip

// NewBoard creates a new Board.
func NewBoard() (b *Board, err error) {
	var (
		copyCBs = []cmdBuilder{
			newCmdBuilder("clip"),
			newCmdBuilder("PowerShell", "-Command", "Set-Clipboard"),
		}
		pasteCBs = []cmdBuilder{
			newCmdBuilder("PowerShell", "-Command", "Get-Clipboard", "-Format",
				"Text", "-Raw", "|", "Write-Host", "-NoNewline"),
			newCmdBuilder("paste"),
		}
	)
	return makeBoardFromPossibleCBs(copyCBs, pasteCBs)
}
