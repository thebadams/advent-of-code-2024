package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicf("Error reading file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Panicf(`Error scanning file: %v`, err)
	}
}
