package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAKE SOME NOISE")
	initSpeaker()

	var track = example_drums()

	runSampler(&track)
}
