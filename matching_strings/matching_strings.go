package matching_strings

func MatchingStrings(strings []string, queries []string) []int32 {
	matches := make([]int32, len(queries))
	for i, query := range queries {
		for _, str := range strings {
			if query == str {
				matches[i]++
			}
		}
	}
	return matches
}
