package day1

import "sort"

func Solve(input string) (partOne int, partTwo int) {
	// read file and get group data
	group1, group2 := readFile(input)
	simScores := GetSimilarityScore(group1, group2)
	//Sort data in place
	sort.Ints(group1)
	sort.Ints(group2)
	//get difference between data
	distances := GetDifference(group1, group2)
	// Sum up the data
	partOne = SumNums(distances)
	partTwo = SumNums(simScores)
	return
}
