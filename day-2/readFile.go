package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(path string) [][]int {
	data := [][]int{}
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nums, err := GetNums(scanner.Text())
		if err != nil {
			fmt.Printf("Error Parsing %v", scanner.Text())
		}

		data = append(data, nums)

	}

	if err := scanner.Err(); err != nil {
		log.Panicf(`Error scanning file: %v`, err)
	}
	return data
}
