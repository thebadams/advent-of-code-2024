package day2

func Solve(input string) int {
	data := readFile(input)
	safeData := []bool{}
	for _, v := range data {
		safe := AnalyzeData(v)
		safeData = append(safeData, safe)
	}

	return CountSafe(safeData)
}
