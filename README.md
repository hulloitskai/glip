# glip

_A clipboard interface for Go, compatible with Windows, Mac OS X, and Linux._

[![godoc: reference][godoc-img]][godoc]
[![codecov: coverage][codecov-img]][codecov]
[![travis: build][travis-img]][travis]
[![appveyor: build][appveyor-img]][appveyor]

## Usage

Install `glip` like you would any other Go package:

```bash
go get github.com/steven-xie/glip
```

Since a `glip.Board` (a clipboard interface) implements `io.Writer`,
`io.Reader`, `io.WriterTo`, it can be used just about anywhere:

```go
import (
  "fmt"
  "github.com/steven-xie/glip"
)

func main() {
  // Save "snip snip" into the system clipboard.
  glip.WriteString("snip snip")
  // And we're done!

  // Read clipboard contents into a string (ignoring the error).
  out, _ := glip.ReadString()
  fmt.Println(out)
  // Output: snip snip
}
```

## Compatibility

### Windows:

`glip` uses the `clip` commands on Windows to write to the system clipboard.
This is available _starting from Windows 7, and onwards_.

No native paste command is availble on Windows, but if
[this third-party `paste` command](https://www.c3scripts.com/tutorials/msdos/paste.html)
is installed, it will be used to read from the system clipboard.

### Mac OS X:

`glip` uses `pbcopy` and `pbpaste` commands on OS X; these commands have been
available since 2005, so no compatibility worries here.

### Linux:

`glip` requires the _installation of either `xclip` or `xsel`_ to function on
Linux (since there's no built-in clipboard interface). `glip` will choose one
of those two programs automatically, unless you build a custom board with the
`NewLinuxBoard` function.

_`glip` has not been tested yet on Linux!_

<br />

## glipboard

For an example of an application that uses `glip`, check out
[`glipboard`](https://github.com/steven-xie/glipboard), a universal clipboard
command-line accessor.

[godoc]: https://godoc.org/github.com/steven-xie/glip
[godoc-img]: https://godoc.org/github.com/steven-xie/glip?status.svg
[travis]: https://travis-ci.org/steven-xie/glip
[travis-img]: https://travis-ci.org/steven-xie/glip.svg?branch=master
[codecov]: https://codecov.io/gh/steven-xie/glip
[codecov-img]: https://codecov.io/gh/steven-xie/glip/branch/master/graph/badge.svg
[appveyor]: https://ci.appveyor.com/project/StevenXie/glip
[appveyor-img]: https://ci.appveyor.com/api/projects/status/ntdxh30vlbo55da7?svg=true
