package portal

import (
	"fmt"
	"io"
)

// iores represents the result of an IO function call.
type iores struct {
	// n is the number bytes read or written using the function.
	n int
	// err is the error that the function encountered.
	err error
}

// iores64 is like iores, but where n is an int64 instead of an int.
type iores64 struct {
	n   int64
	err error
}

// asyncCopy asynchronously copies data from usrc into dst, and reports the
// results to res.
func asyncCopy(dst io.Writer, src io.Reader, res chan iores64) {
	n, err := io.Copy(dst, src)

	if err != nil {
		res <- iores64{
			n:   0,
			err: fmt.Errorf("portal: error during asynchronous copy: %v", err),
		}
	}

	res <- iores64{n, nil}
}

func asyncWrite(w io.Writer, p []byte, res chan iores) {
	n, err := w.Write(p)

	if err != nil {
		res <- iores{
			n:   0,
			err: fmt.Errorf("portal: error during asyncrhonous write: %v", err),
		}
	}

	res <- iores{n, nil}
}

func asyncRead(r io.Reader, dst []byte, res chan iores) {
	n, err := r.Read(dst)

	if err != nil {
		res <- iores{
			n:   0,
			err: fmt.Errorf("portal: error during asyncrhonous read: %v", err),
		}
	}

	res <- iores{n, nil}
}
