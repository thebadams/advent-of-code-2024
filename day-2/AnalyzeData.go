package day2

func AnalyzeData(data []int) bool {

	increases := 0
	decreases := 0
	// count numbers of increase, decrease, no change between numbers
	for i, v := range data {
		if i < len(data)-1 {
			if v > data[i+1] {
				// if difference is greater than 3, unsafe data return false
				if v-data[i+1] > 3 {
					return false
				}
				decreases = decreases + 1
			} else if v < data[i+1] {
				// if difference is greater than 3, unsafe data return false
				if data[i+1]-v > 3 {
					return false
				}
				increases = increases + 1
			} else {
				// if there was neither an increase nor decrease, immediately return false. Unsafe data
				return false
			}

		}
	}
	// return false if there were both increases and decreases
	if increases > 0 && decreases > 0 {
		return false
	}

	return true
}
