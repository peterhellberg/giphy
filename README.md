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
giphy trending

 0 - http://media0.giphy.com/media/kMM7XbRvSgpAA/giphy.gif
 1 - http://media2.giphy.com/media/lgcaIKboeo8ZW/giphy.gif
 2 - http://media2.giphy.com/media/vcyroBgx2nrby/giphy.gif
 3 - http://media2.giphy.com/media/UVHUzM00JWsDu/giphy.gif
 4 - http://media3.giphy.com/media/9qAgUlR16YgYo/giphy.gif
 5 - http://media1.giphy.com/media/gMlBLPI3zCsXC/giphy.gif
 6 - http://media3.giphy.com/media/B81XkL3dtnWTe/giphy.gif
 7 - http://media3.giphy.com/media/Iu7nC0GDWxu9y/giphy.gif
 8 - http://media0.giphy.com/media/B5oEcYBqZe22s/giphy.gif
 9 - http://media3.giphy.com/media/rjr9etfxrdP3i/giphy.gif
10 - http://media3.giphy.com/media/SYQpxVZGn56BG/giphy.gif
11 - http://media1.giphy.com/media/LJ5iCphkPZVPW/giphy.gif
12 - http://media4.giphy.com/media/d8YBM0IgPszle/giphy.gif
13 - http://media0.giphy.com/media/GKleSu1rgesSc/giphy.gif
14 - http://media0.giphy.com/media/U1YnBiy8rNUqI/giphy.gif
15 - http://media2.giphy.com/media/Ekla75geXUIYo/giphy.gif
16 - http://media2.giphy.com/media/FNfJO1GJHkbi8/giphy.gif
17 - http://media1.giphy.com/media/ODzsRb1t8nBv2/giphy.gif
18 - http://media0.giphy.com/media/13ln9K5TWkNTLa/giphy.gif
19 - http://media0.giphy.com/media/dmucaITbJPSY8/giphy.gif
20 - http://media0.giphy.com/media/zE0RFo8wgHJHW/giphy.gif
21 - http://media0.giphy.com/media/o0rzRKtJlO1na/giphy.gif
22 - http://media4.giphy.com/media/60RaCtUr2hXTa/giphy.gif
23 - http://media4.giphy.com/media/8Bbl0U61TN6DK/giphy.gif
24 - http://media4.giphy.com/media/aN1pVF1L997rO/giphy.gif
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
