package day2

func Solve(input string) (int, int) {
	data := readFile(input)
	safeData := []bool{}
	stillSafeData := []bool{}
	for _, v := range data {
		safe, _ := AnalyzeData(v)
		safeData = append(safeData, safe)
		stillSafe := ProblemDampener(v)
		stillSafeData = append(stillSafeData, stillSafe)
	}

	return CountSafe(safeData), CountSafe(stillSafeData)
}
