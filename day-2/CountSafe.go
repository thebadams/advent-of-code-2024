package day2

func CountSafe(input []bool) (safe int) {
	safe = 0
	for _, v := range input {
		if v == true {
			safe = safe + 1
		}
	}
	return

}
