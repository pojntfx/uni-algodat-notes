package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := []int{5, 6, 9, 15, 31, 43545, 234234234}

	for _, input := range inputs {
		outputImperative := algorithms.GetDigitSumImperative(input)
		outputFunctional := algorithms.GetDigitSumFunctional(input, 0)

		fmt.Printf("in=%-3v out=%v=%v\n", input, outputImperative, outputFunctional)
	}
}
