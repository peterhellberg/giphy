giphy
=====

Go library for the [Giphy API](https://developers.giphy.com/docs/api/)

[![Build status](https://github.com/peterhellberg/giphy/actions/workflows/test.yml/badge.svg)](https://github.com/peterhellberg/giphy/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/giphy)](https://goreportcard.com/report/github.com/peterhellberg/giphy)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/peterhellberg/giphy)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/giphy#license-mit)

> You will need an API key from Giphy, instructions to get one can be found here: <https://developers.giphy.com/docs/api#quick-start-guide>

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
GIPHY_API_KEY=[your-api-key] GIPHY_LIMIT=4 giphy search computer
http://media3.giphy.com/media/wvHC7zyCedEI0/giphy.gif
http://media3.giphy.com/media/gr8K2b72UefvO/giphy.gif
http://media2.giphy.com/media/L2u68v1MmZv5m/giphy.gif
http://media1.giphy.com/media/4WOs6Af0nOOeQ/giphy.gif
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

  g := giphy.DefaultClient

  res, err := g.Translate(os.Args[1:])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(res.Data.MediaURL())
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
      fmt.Println(i, "-", d.MediaURL())
    }
  }
}
```

### Run an example

```bash
GIPHY_API_KEY=[your-api-key] GIPHY_RATING=pg-13 go run example.go
```

## License (MIT)

Copyright (c) 2015-2021 [Peter Hellberg](https://c7.se)

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
