// +build darwin

package glip

import "fmt"

// NewBoard creates a new Board, using a program automatically selected based
// on the operating system and available system commands.
func NewBoard() (b Board, err error) {
	if b, err = NewDarwinBoard(); err != nil {
		return nil, fmt.Errorf("glip: could not create Board: %v", err)
	}
	return b, err
}
