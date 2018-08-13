// Package portal provides a Portal, which wraps and improves the handling of
// IO-heavy commands.
package portal

import (
	"fmt"
	"io"
	"os/exec"
)

// Portal is a struct used to conveniently read and write data to a named
// external program opened by exec.Cmd, by implementing functions required
// by io.Writer, io.Reader, and io.WriterTo.
type Portal struct {
	*exec.Cmd
	Name string
}

// New creates a new Portal.
func New(c *exec.Cmd, name string) *Portal {
	return &Portal{Cmd: c, Name: name}
}

// StdoutPipe returns a pipe that will be connected to the command's standard
// output when the command starts.
func (p *Portal) StdoutPipe() (out io.ReadCloser, err error) {
	if out, err = p.Cmd.StdoutPipe(); err != nil {
		return out, fmt.Errorf("failed to open %s command output: %v", p.Name, err)
	}
	return out, nil
}

// StdinPipe returns a pipe that will be connected to the command's standard
// input when the command starts.
func (p *Portal) StdinPipe() (in io.WriteCloser, err error) {
	if in, err = p.Cmd.StdinPipe(); err != nil {
		return in, fmt.Errorf("failed to open %s command input: %v", p.Name, err)
	}
	return in, nil
}

// Start starts the specified command but does not wait for it to complete.
func (p *Portal) Start() error {
	if err := p.Cmd.Start(); err != nil {
		return fmt.Errorf("failed to start %s command: %v", p.Name, err)
	}
	return nil
}

// Wait waits for the command to exit and waits for any copying to stdin or
// copying from stdout or stderr to complete.
func (p *Portal) Wait() error {
	if err := p.Cmd.Wait(); err != nil {
		return fmt.Errorf("%s command failed to stop correctly: %v", p.Name, err)
	}
	return nil
}

// Read allows for the reading of data from a command's standard output.
func (p *Portal) Read(dst []byte) (n int, err error) {
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, err
	}

	if err = p.Start(); err != nil {
		return 0, err
	}

	if n, err = out.Read(dst); err != nil {
		return n, fmt.Errorf("failed to read output from %s command: %v", p.Name,
			err)
	}

	err = p.Wait()
	return n, err
}

// ReadFrom allows for the piping of data from a io.Writer into a command's
// standard output.
func (p *Portal) ReadFrom(r io.Reader) (n int64, err error) {
	in, err := p.StdinPipe()
	if err != nil {
		return 0, err
	}

	if err = p.Start(); err != nil {
		return 0, err
	}

	if n, err = io.Copy(in, r); err != nil {
		return n, fmt.Errorf("failed to write to %s command input: %v", p.Name, err)
	}

	if err = in.Close(); err != nil {
		return n, fmt.Errorf("failed to close input to %s command: %v", p.Name, err)
	}

	err = p.Wait()
	return n, err
}

// Write allows for the writing of data into a command's standard input.
func (p *Portal) Write(src []byte) (n int, err error) {
	in, err := p.StdinPipe()
	if err != nil {
		return 0, err
	}

	if err = p.Start(); err != nil {
		return 0, err
	}

	if n, err = in.Write(src); err != nil {
		return n, fmt.Errorf("failed to write to %s command input: %v", p.Name, err)
	}

	if err = in.Close(); err != nil {
		return n, fmt.Errorf("failed to close input to %s command: %v", p.Name, err)
	}

	err = p.Wait()
	return n, err
}

// WriteTo allows for the piping of data from a command's standard output into
// an io.Writer.
func (p *Portal) WriteTo(w io.Writer) (n int64, err error) {
	out, err := p.StdoutPipe()
	if err != nil {
		return 0, err
	}

	if err = p.Start(); err != nil {
		return 0, err
	}

	if n, err = io.Copy(w, out); err != nil {
		return n, fmt.Errorf("failed to copy output from %s command: %v", p.Name,
			err)
	}

	err = p.Wait()
	return n, err
}
