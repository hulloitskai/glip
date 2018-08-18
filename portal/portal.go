// Package portal provides a Portal, which wraps and improves the handling of
// IO-heavy commands.
package portal

import (
	"fmt"
	"io"
	"os/exec"
)

// Portal is used to conveniently read and write data to an external program
// opened by exec.Cmd, by implementing functions required IO interfaces such as
// io.Writer, io.WriterTo, io.Reader, and io.ReaderFrom.
type Portal struct {
	*exec.Cmd
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
	return &Portal{cmd}
}

//////////////////////////
// Reading from Portal
//////////////////////////

// Read allows for the reading of data from a command's standard output.
func (p *Portal) Read(dst []byte) (n int, err error) {
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, stdoutPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = out.Read(dst); err != nil {
		return n, fmt.Errorf("portal: failed to read from Stdout: %v", err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, nil
}

// WriteTo allows for the piping of data from a command's standard output into
// an io.Writer.
func (p *Portal) WriteTo(w io.Writer) (n int64, err error) {
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, stdoutPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = io.Copy(w, out); err != nil {
		return n, fmt.Errorf("portal: could not to copy from Stdout: %v", err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, nil
}

//////////////////////////
// Writing to Portal
//////////////////////////

// Write allows for the writing of data into a command's standard input.
func (p *Portal) Write(src []byte) (n int, err error) {
	in, err := p.StdinPipe()
	if err != nil {
		return 0, stdinPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = in.Write(src); err != nil {
		return n, fmt.Errorf("portal: could not write to Stdin: %v", err)
	}
	if err = in.Close(); err != nil {
		return n, closeStdinErr(err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, err
}

// ReadFrom allows for the piping of data from a io.Writer into a command's
// standard output.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	in, err := p.StdinPipe()
	if err != nil {
		return 0, stdinPipeErr(err)
	}
	if err = p.Start(); err != nil {
		return 0, startErr(err)
	}
	if n, err = io.Copy(in, r); err != nil {
		return n, fmt.Errorf("portal: failed to write to Stdin: %v", err)
	}
	if err = in.Close(); err != nil {
		return n, closeStdinErr(err)
	}
	if err = p.Wait(); err != nil {
		return n, waitErr(err)
	}
	return n, err
}
