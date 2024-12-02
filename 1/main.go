package main

import (
	"flag"
)

func main() {
	part := flag.Bool("next", false, "second part")

	flag.Parse()

	if *part {
		B()
	} else {
		A()
	}

}
