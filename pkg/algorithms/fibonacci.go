package algorithms

func GetFibonacciImperative(num int) int {
	nums := make([]int, num+1)

	nums[0] = 0
	nums[1] = 1

	for i := 2; i <= num; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums[num]
}

func GetFibonacciFunctional(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return GetFibonacciFunctional(num-1) + GetFibonacciFunctional(num-2)
}
