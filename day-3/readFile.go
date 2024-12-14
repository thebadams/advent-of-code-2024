package day3

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readFile(path string) (data strings.Builder) {

	// Open the file
	file, err := os.Open(path) // Replace "filename.txt" with the actual file name
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Loop through the file, reading each line
	for scanner.Scan() {
		line := scanner.Text()
		data.WriteString(line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data

}
