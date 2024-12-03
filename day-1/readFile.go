package day1

import (
	"bufio"
	"log"
	"os"
)

func readFile(path string) (group1 []int, group2 []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums := GetNums(scanner.Text())
		group1 = append(group1, nums[0])

		group2 = append(group2, nums[1])

	}

	if err := scanner.Err(); err != nil {
		log.Panicf(`Error scanning file: %v`, err)
	}
	return
}
