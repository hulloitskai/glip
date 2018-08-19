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

// New creates a new Portal using exec.Command internally.
func New(name string, args ...string) *Portal {
	return MakeFrom(exec.Command(name, args...))
}

// MakeFrom creates a Portal by wrapping exec.Cmd.
//
// If "cmd" is nil, MakeFrom will return nil.
func MakeFrom(cmd *exec.Cmd) *Portal {
	if cmd == nil {
		return nil
	}
	return &Portal{blueprint: cmd}
}

// Restore loads Portal's exec.Cmd instance from an internal reference,
// preparing it to perform a comamnd.
func (p *Portal) Restore() {
	p.Cmd = *p.blueprint
}
