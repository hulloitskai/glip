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

// New creates a new Portal.
//
// It uses exec.Command internally to generate the command.
func New(name string, args ...string) *Portal {
	return NewFrom(exec.Command(name, args...))
}

// NewFrom creates a new Portal by wrapping an exec.Cmd.
//
// If cmd is nil, NewFrom will return nil.
func NewFrom(cmd *exec.Cmd) *Portal {
	if cmd == nil {
		return nil
	}
	return &Portal{Cmd: *cmd, blueprint: cmd}
}

// Reload resets p.Cmd, preparing it for a new execution.
//
// Reload must be called before each execution of p.Cmd.
func (p *Portal) Reload() {
	p.Cmd = *p.blueprint
}

// IsReady determines if the Portal is ready to be executed.
//
// If IsReady returns false, it means that p.Cmd has already finished executing,
// and will fail if started again.
func (p *Portal) IsReady() bool {
	if p.ProcessState == nil {
		return true
	}
	return !p.ProcessState.Exited()
}

// PersistentArgs gets the arguments of the Portal's exec.Cmd blueprint.
//
// These arguments will persist between Portal command executions.
func (p *Portal) PersistentArgs() []string {
	return p.blueprint.Args
}

// SetPersistentArgs sets the arguments of Portal's exec.Cmd blueprint.
//
// These arguments will persist between Portal command executions.
func (p *Portal) SetPersistentArgs(args []string) {
	p.blueprint.Args = args
}
