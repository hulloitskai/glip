// +build !windows

package glip

import "github.com/steven-xie/glip/portal"

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

// xselapp is the base structure for an application that interacts with X
// selections.
type xselapp struct {
	*portal.Portal
	sel XSelection
}
