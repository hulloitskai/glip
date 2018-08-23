// Package portal provides a Portal, which wraps and improves the handling of
// reusable, IO-heavy commands.
package portal

import "os/exec"

// Portal is used to conveniently read and write data to an external program
// opened by exec.Cmd, by implementing functions required IO interfaces such as
// io.Writer, io.WriterTo, io.Reader, and io.ReaderFrom.
type Portal struct {
	exec.Cmd
	blueprint *exec.Cmd
}

// New creates a new, unprepared Portal.
//
// It uses exec.Command internally to generate the command.
func New(name string, args ...string) *Portal {
	return MakeFrom(exec.Command(name, args...))
}

// MakeFrom creates a new, unprepared Portal by wrapping exec.Cmd.
//
// If "cmd" is nil, MakeFrom will return nil.
func MakeFrom(cmd *exec.Cmd) *Portal {
	if cmd == nil {
		return nil
	}
	return &Portal{blueprint: cmd}
}

// Prepare loads Portal's exec.Cmd instance from an internal reference,
// preparing it to perform a comamnd.
//
// Prepare must be called before each execution of Portal's internal Cmd.
func (p *Portal) Prepare() {
	p.Cmd = *p.blueprint
}

// IsReady determines if the Portal has been "prepared" to be executed.
//
// If IsReady returns false, that means Portal's internal Cmd field is empty
// and will fail if started.
func (p *Portal) IsReady() bool {
	return p.Cmd.Path != ""
}
