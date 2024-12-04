package main

import (
	"flag"
)

func main() {
	next := flag.Bool("next", false, "is second part")
	sample := flag.Bool("sample", false, "is sample data")

	flag.Parse()

	if *next {
		B(*sample)
	} else {
		A(*sample)
	}
}
