package sum

func Sum(nums []int) int {
	var accum int
	for i := 0; i < len(nums); i++ {
		accum = accum + nums[i]
	}
	return accum
}

func SumRecursion(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumRecursion(nums[1:])
}

func SumWithAccumulator(nums []int, accum int) int {
	if len(nums) == 0 {
		return accum
	}
	return SumWithAccumulator(nums[1:], accum+nums[0])
}
