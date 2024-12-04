package day1

func GetSimilarityScore(group1 []int, group2 []int) []int {
	var similarityScores []int
	for _, v := range group1 {
		score := 0
		for _, v2 := range group2 {
			if v == v2 {
				score = score + 1
			}
		}
		similarityScores = append(similarityScores, score*v)
	}
	return similarityScores
}
