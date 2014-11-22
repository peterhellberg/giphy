package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/peterhellberg/giphy"
)

func main() {
	g := giphy.DefaultClient

	if len(os.Args) < 2 {
		fmt.Println(strings.Join([]string{
			"Commands:",
			"search, s           [args]",
			"gif, id             [args]",
			"random, rand, r     [args]",
			"translate, trans, t [args]",
			"trending, trend, tr [args]",
		}, "\n\t"))

		return
	}

	args := os.Args[1:]

	switch args[0] {
	default:
		search(g, args)
	case "search", "s":
		search(g, args[1:])
	case "gif", "id":
		gif(g, args[1:])
	case "random", "rand", "r":
		random(g, args[1:])
	case "translate", "trans", "t":
		translate(g, args[1:])
	case "trending", "trend", "tr":
		trending(g, args[1:])
	}
}

func search(c *giphy.Client, args []string) {
	res, err := c.Search(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range res.Data {
		fmt.Println(d.Images.Original.URL)
	}
}

func gif(c *giphy.Client, args []string) {
	if len(args) == 0 {
		fmt.Println("missing Giphy id")
		os.Exit(1)
	}

	res, err := c.GIF(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.Images.Original.URL)
}

func random(c *giphy.Client, args []string) {
	res, err := c.Random(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.ImageOriginalURL)
}

func translate(c *giphy.Client, args []string) {
	res, err := c.Translate(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(res.Data.Images.Original.URL)
}

func trending(c *giphy.Client, args []string) {
	res, err := c.Trending(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, d := range res.Data {
		fmt.Println(d.Images.Original.URL)
	}
}
