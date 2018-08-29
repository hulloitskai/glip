package portal_test

import (
	"github.com/steven-xie/glip/portal"
	"os"
)

const TestPhrase = "Hello Portal!"

func ExamplePortal() {
	p := portal.New("echo", TestPhrase)
	p.WriteTo(os.Stdout)
	// Output: Hello Portal!
}
