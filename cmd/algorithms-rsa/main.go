package main

import (
	"fmt"

	"github.com/pojntfx/uni-algodat-notes/pkg/algorithms"
)

func main() {
	inputs := [][]int{
		{22, 5, 7, 0}, // m p q "0 -> can't fail, 1 -> can fail"
		{2, 3, 7, 0},
		{69, 5, 7, 1},
	}

	for _, input := range inputs {
		{
			m, p, q := input[0], input[1], input[2]

			keys := algorithms.GetRSAKeysImperative(p, q)
			n, e, d := keys[0], keys[1], keys[2]

			c, err := algorithms.RSAEncrypt(m, e, n)
			if err != nil {
				if input[3] == 1 {
					fmt.Printf("imp: (p=%v q=%v) -> (n=%v e=%v d=%v) -> skipped because keyspace to short\n", p, q, n, e, d)

					continue
				}

				panic(err)
			}
			m2 := algorithms.RSADecrypt(c, d, n)

			fmt.Printf("imp: (p=%v q=%v) -> (n=%v e=%v d=%v) -> (m=%v) -> (c=%v) -> (m2=%v)\n", p, q, n, e, d, m, c, m2)
		}

		{
			m, p, q := input[0], input[1], input[2]

			keys := algorithms.GetRSAKeysFunctional(p, q, 1, 2)
			n, e, d := keys[0], keys[1], keys[2]

			c, err := algorithms.RSAEncrypt(m, e, n)
			if err != nil {
				if input[3] == 1 {
					fmt.Printf("imp: (p=%v q=%v) -> (n=%v e=%v d=%v) -> skipped because keyspace to short\n", p, q, n, e, d)

					continue
				}

				panic(err)
			}
			m2 := algorithms.RSADecrypt(c, d, n)

			fmt.Printf("fun: (p=%v q=%v) -> (n=%v e=%v d=%v) -> (m=%v) -> (c=%v) -> (m2=%v)\n", p, q, n, e, d, m, c, m2)
		}
	}
}
