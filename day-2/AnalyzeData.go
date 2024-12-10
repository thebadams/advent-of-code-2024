package day2

func AnalyzeData(data []int) (bool, int) {
	decreases := 0
	increases := 0
	for i := 1; i < len(data); i++ {
		// compare to previous value
		diff := data[i] - data[i-1]

		if diff > 0 {
			increases = increases + 1
		}
		if diff < 0 {
			decreases = decreases + 1
		}
		if increases > 0 && decreases > 0 {
			return false, i
		}
		if diff > 3 || diff < -3 {

			return false, i
		}
		if diff == 0 {
			return false, i
		}
	}
	return true, 0
}
