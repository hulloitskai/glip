package glip

// XSelection is a string that represents a particular X selection (a
// clipboard data source that can be written to / read from).
type XSelection string

// These XSelections represent the available X selections where data from
// "xclip" and "xsel" can be stored.
const (
	XSPrimary   XSelection = "primary"   // also known as XA_PRIMARY
	XSSecondary            = "secondary" // also known as XA_SECONDARY
	XSClipboard            = "clipboard" // also known as XA_CLIPBOARD
)

// XBoard is a Board that interacts with an X server (and thus uses either
// Xclip or Xsel under-the-hood).
type XBoard interface {
	// Opts exposes an X-server-related options struct.
	Opts() *Xopts
	Board
}
