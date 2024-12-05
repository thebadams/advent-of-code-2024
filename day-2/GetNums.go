package day2

import (
	"strconv"
	"strings"
)

func GetNums(data string) ([]int, error) {
	nums := []int{}
	subStr := strings.Split(data, " ")
	for _, v := range subStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nums, err
		}
		nums = append(nums, num)
	}
	return nums, nil

}
