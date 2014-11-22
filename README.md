giphy
=====

Go library for the [Giphy API](https://github.com/Giphy/GiphyAPI)

[![GoDoc](https://godoc.org/github.com/peterhellberg/giphy?status.svg)](https://godoc.org/github.com/peterhellberg/giphy)

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

## Examples

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
