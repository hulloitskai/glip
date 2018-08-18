package portal_test

import (
	"github.com/steven-xie/glip/portal"
	"os"
)

func ExamplePortal() {
	const in = "Hellooooooo"
	p := portal.New("echo", in)

	p.WriteTo(os.Stdout)
	// Output: Hellooooooo
}
