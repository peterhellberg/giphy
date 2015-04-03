giphy
=====

Go library for the [Giphy API](https://github.com/Giphy/GiphyAPI)

[![GoDoc](https://godoc.org/github.com/peterhellberg/giphy?status.svg)](https://godoc.org/github.com/peterhellberg/giphy)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/giphy#license-mit)

## Command line tool

### Installation

```bash
go get -u github.com/peterhellberg/giphy/cmd/giphy
```

### Usage


```bash
Commands:
	search, s           [args]
	gif, id             [args]
	random, rand, r     [args]
	translate, trans, t [args]
	trending, trend, tr [args]
```

```bash
GIPHY_LIMIT=4 giphy trending
http://media0.giphy.com/media/kMM7XbRvSgpAA/giphy.gif
http://media2.giphy.com/media/lgcaIKboeo8ZW/giphy.gif
http://media2.giphy.com/media/vcyroBgx2nrby/giphy.gif
http://media2.giphy.com/media/UVHUzM00JWsDu/giphy.gif
```

## Examples of using the library

### translate.go

```go
package main

import (
  "fmt"
  "os"

  "github.com/peterhellberg/giphy"
)

func main() {
  if len(os.Args) < 2 {
    return
  }

  c := giphy.DefaultClient

  res, err := c.Translate(os.Args[1:])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(res.Data.URL)
}

```

### trending.go

```go
package main

import (
  "fmt"

  "github.com/peterhellberg/giphy"
)

func main() {
  g := giphy.DefaultClient

  if trending, err := g.Trending(); err == nil {
    for i, d := range trending.Data {
      fmt.Println(i, "-", d.URL)
    }
  }
}
```

### Run an example

```bash
GIPHY_API_KEY=dc6zaTOxFJmzC GIPHY_RATING=pg-13 go run example.go
```

## License (MIT)

Copyright (c) 2015 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

![Computer](http://media.giphy.com/media/MzX5hCfR5nP20/giphy.gif)
