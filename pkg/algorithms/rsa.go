package algorithms

import (
	"errors"
	"math"
)

func GetRSAKeysImperative(p int, q int) []int {
	n := p * q

	phi := (p - 1) * (q - 1)

	e := 2
	for {
		if phi%e != 0 {
			break
		}

		e++
	}

	k := 1
	for {
		if (k*phi+1)%e == 0 {
			break
		}

		k++
	}

	d := (k*phi + 1) / e

	return []int{n, e, d}
}

func RSAEncrypt(m int, e int, n int) (int, error) {
	if m > n {
		return -1, errors.New("keyspace to short")
	}

	return int(math.Pow(float64(m), float64(e))) % n, nil
}

func RSADecrypt(c int, d int, n int) int {
	return int(math.Pow(float64(c), float64(d))) % n
}

func GetRSAKeysFunctional(p int, q int, k int, e int) []int {
	/*
		To make this fully side-effect free, replace all references
		to the following with these expressions
	*/
	n := p * q
	phi := (p - 1) * (q - 1)

	if phi%e != 0 {
		if (k*phi+1)%e == 0 {
			return []int{n, e, (k*phi + 1) / e}
		}

		return GetRSAKeysFunctional(p, q, k+1, e)
	}

	return GetRSAKeysFunctional(p, q, k, e+1)
}
