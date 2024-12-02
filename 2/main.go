package main

import (
	"flag"
)

func main() {
	next := flag.Bool("next", false, "is second part")

	flag.Parse()

	if *next {
		B()
	} else {
		A()
	}
}
