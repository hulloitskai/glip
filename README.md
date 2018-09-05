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

Since a `glip.Board` (a clipboard interface) includes `io.Writer`, `io.Reader`,
`io.WriterTo`, and `io.ReaderFrom`, it can be used just about anywhere:

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

### Advanced Usage:

`glip` provides API-wrapping structs specific to each clipboard-accessing
program; each struct implements the `Board` interface.

If you know which specific program you would like to use, you can create a new
instance of the associated wrapper struct (i.e. `PShellBoard`, `Clip`,
`DarwinBoard`, `Xclip`, or `Xsel`), and set program-specific flags / options
using the wrapper struct.

<br />

## Compatibility

### Windows:

`glip` uses the the PowerShell `Get-Clipboard` and `Set-Clipboard` cmdlets to
read and write to the Windows clipboard.

If PowerShell is not available, the `clip` command is used to write to the
Windows clipboard.

### macOS:

`glip` uses `pbcopy` and `pbpaste` commands on macOS.

### Linux:

`glip` requires the _installation of either `xclip` or `xsel`_ to function on
Linux (since there's no built-in clipboard interface). `glip` will choose
whichever program is available, with a preference for `xsel` if both are
avaiklable.

<br />

## glipboard

For an example of an application that uses `glip`, check out `glipboard`
(located at `/cmd/glipboard/`).

`glipboard` was developed to both showcase how `glip` can be used in a real
application, as well as to be a universal clipboard interface that external
programs can call in order to write to a system clipboard, if the underlying
commands are available.

[godoc]: https://godoc.org/github.com/steven-xie/glip
[godoc-img]: https://godoc.org/github.com/steven-xie/glip?status.svg
[travis]: https://travis-ci.org/steven-xie/glip
[travis-img]: https://travis-ci.org/steven-xie/glip.svg?branch=master
[codecov]: https://codecov.io/gh/steven-xie/glip
[codecov-img]: https://codecov.io/gh/steven-xie/glip/branch/master/graph/badge.svg
[appveyor]: https://ci.appveyor.com/project/StevenXie/glip
[appveyor-img]: https://ci.appveyor.com/api/projects/status/ntdxh30vlbo55da7/branch/master?svg=true
