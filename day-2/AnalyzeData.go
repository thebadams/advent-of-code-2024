package day2

func AnalyzeData(data []int) bool {
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
			return false
		}
		if diff > 3 || diff < -3 {

			return false
		}
		if diff == 0 {
			return false
		}
	}
	return true
}
