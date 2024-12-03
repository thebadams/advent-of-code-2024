package day1

func GetDifference(group1 []int, group2 []int) []int {
	result := []int{}
	for i := range group1 {
		var difference int
		if group1[i] >= group2[i] {
			difference = group1[i] - group2[i]
		} else if group2[i] >= group1[i] {
			difference = group2[i] - group1[i]
		}
		result = append(result, difference)
	}
	return result

}
