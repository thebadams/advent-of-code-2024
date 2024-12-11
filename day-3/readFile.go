package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(path string) []string {
	data := []string{}
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			fmt.Printf("Error Parsing %v", scanner.Text())
		}
		data = append(data, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Panicf(`Error scanning file: %v`, err)
	}
	return data
}
