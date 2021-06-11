package algorithms

func GetRussianProductImperative(a int, b int) int {
	extra := 0

	for a != 1 {
		if a%2 != 0 {
			extra += b
		}

		a = a / 2
		b = b * 2
	}

	return b + extra
}

func GetRussianProductFunctional(a int, b int, extra int) int {
	if a == 1 {
		return b + extra
	}

	/*
		To make this fully side-effect free, replace all references
		to the following with these expressions
	*/
	x := a / 2
	y := b * 2

	if a%2 != 0 {
		return GetRussianProductFunctional(x, y, extra+b)
	}

	return GetRussianProductFunctional(x, y, extra)
}
