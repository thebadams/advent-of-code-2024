package day1

func SumNums(nums []int) int {
	sum := 0
	for i := range nums {
		sum = sum + nums[i]
	}
	return sum
}
