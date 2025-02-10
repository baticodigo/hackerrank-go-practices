package breaking_the_records

func BreakingRecords(scores []int32) []int32 {
	if len(scores) == 0 {
		return []int32{0, 0}
	}
	minScore, maxScore := scores[0], scores[0]
	minBreaks, maxBreaks := 0, 0
	for _, score := range scores[1:] {
		switch {
		case score > maxScore:
			maxScore = score
			maxBreaks++
		case score < minScore:
			minScore = score
			minBreaks++
		}
	}
	return []int32{int32(maxBreaks), int32(minBreaks)}
}

/*

Input (stdin)
9
10 5 20 20 4 5 2 25 1
Your Output (stdout)
2 4
Expected Output
2 4

*/
