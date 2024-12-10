package day2

func ProblemDampener(data []int) bool {
	isSafe, _ := AnalyzeData(data)
	// if the data is safe without the dampener, return true immediately.
	if isSafe {
		return true
	}

	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		isValid, _ := AnalyzeData(data1)
		if isValid == true {
			return true
		}
	}

	return false

}
