package day3

import (
	"regexp"
)

func UncorruptData(data string) (uncorruptedData []string) {
	re := regexp.MustCompile(`(do\(\))|(don\'t\(\))|(mul\([0-9]{1,3},[0-9]{1,3}\))`)
	parsed := re.FindAllString(data, -1)
	uncorruptedData = append(uncorruptedData, parsed...)
	return uncorruptedData
}
