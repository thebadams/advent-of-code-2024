package day1

import "sort"

func Solve(input string) int {
	group1, group2 := readFile(input)
	sort.Ints(group1)
	sort.Ints(group2)
	return 11
}
