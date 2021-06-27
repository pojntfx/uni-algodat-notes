package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []string{
		"test",
		"test asdf 9d2i",
		`   asdf dasdf
asdf`,
		"asdf	asd sad",
		"yay! ",
	}

	for _, input := range inputs {
		outputImperative := algorithms.TrimWhitespaceImperative(input)
		outputFunctional := algorithms.TrimWhitespaceFunctional(input, "", 0)

		fmt.Printf(`in="%v" out="%v"="%v"`+"\n", input, outputImperative, outputFunctional)
	}
}
