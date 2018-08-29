package glip

import (
	"fmt"
	"github.com/steven-xie/glip/portal"
	"io"
	"strings"
)

// dynPortal is a "dynamic" Portal that is a able to dynamically configure its
// arguments before using its IO methods.
//
// dynPortal's arguments are retrieved through GetArgs, and are set for the
// duration of Write, WriteTo, Read, and ReadFrom. They are reset to the
// arguments dynPortal was constructed with at the end of the aforementioned
// IO methods.
type dynPortal struct {
	*portal.Portal
	GetArgs func() []string
}

// newDynPortal constructs a new dynPortal from a program name and arguments.
// These arguments will persist after running dynPortal's IO methods.
func newDynPortal(name string, args ...string) *dynPortal {
	return newDynPortalFrom(portal.New(name, args...))
}

// newDynPortalFrom constructs a dynPortal from an existing portal.Portal.
func newDynPortalFrom(p *portal.Portal) *dynPortal {
	return &dynPortal{Portal: p}
}

// AppendArgs appends args to dynPortal's temporary program arguments.
func (dp *dynPortal) AppendArgs(args ...string) {
	dp.Args = append(dp.Args, args...)
}

// preflight is run before dynPortal's IO methods, to set its temporary
// arguments using GetArgs.
func (dp *dynPortal) preflight() {
	dp.AppendArgs(dp.GetArgs()...)
}

// Write writes len(p) bytes to the dynPortal.
func (dp *dynPortal) Write(p []byte) (n int, err error) {
	dp.preflight()
	return dp.Portal.Write(p)
}

// WriteString writes a string to the dynPortal.
func (dp *dynPortal) WriteString(s string) (n int, err error) {
	return dp.Write([]byte(s))
}

// ReadFrom reads data from an io.Reader into the dynPortal.
func (dp *dynPortal) ReadFrom(r io.Reader) (n int64, err error) {
	dp.preflight()
	return dp.Portal.ReadFrom(r)
}

// Read reads data from the dynPortal to dst.
func (dp *dynPortal) Read(dst []byte) (n int, err error) {
	dp.preflight()
	return dp.Portal.Read(dst)
}

// ReadString reads data from the dynPortal as a string.
func (dp *dynPortal) ReadString() (s string, err error) {
	builder := new(strings.Builder)
	if _, err = dp.WriteTo(builder); err != nil {
		return "", fmt.Errorf("glip: could not write to strings.Builder: %v", err)
	}
	return builder.String(), nil
}

// WriteTo writes data from the dynPortal into an io.Writer.
func (dp *dynPortal) WriteTo(w io.Writer) (n int64, err error) {
	dp.preflight()
	return dp.Portal.WriteTo(w)
}

// ReadPortal exposes a portal.Portal used for reading from the clipboard.
func (dp *dynPortal) ReadPortal() *portal.Portal {
	return dp.Portal
}

// WritePortal exposes a portal.Portal used for writing to the clipboard.
func (dp *dynPortal) WritePortal() *portal.Portal {
	return dp.Portal
}
