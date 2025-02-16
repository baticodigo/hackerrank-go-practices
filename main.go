package main

import (
	"fmt"
	"hackerrank-go-practices/breaking_the_records"
	"hackerrank-go-practices/camel_case_four"
	"hackerrank-go-practices/divisible_sum_pairs"
	"hackerrank-go-practices/matching_strings"
)

func main() {

	fmt.Println("Ejercicio Breaking Records")
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	result := breaking_the_records.BreakingRecords(scores)
	fmt.Println(result)
	fmt.Println("Ejercicio Camel Case Four")
	testCaseValue := []string{"S;M;plasticCup()", "C;V;mobile phone", "C;C;coffee machine", "S;C;LargeSoftwareBook", "C;M;white sheet of paper", "S;V;pictureFrame"}
	for _, testWord := range testCaseValue {
		camel_case_four.CamelCaseFour(testWord)
	}
	fmt.Println("Ejercicio Divisible Sum")
	validPair := divisible_sum_pairs.DivisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2})
	fmt.Println(validPair)
	fmt.Println("Ejercicio Matching Strings")
	valuesStringMarch := matching_strings.MatchingStrings([]string{"ab", "ab", "abc"}, []string{"ab", "abc", "bc"})
	fmt.Println(valuesStringMarch)
}
