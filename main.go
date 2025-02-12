package main

import (
	"fmt"
	"hackerrank-go-practices/breaking_the_records"
	"hackerrank-go-practices/camel_case_four"
)

func main() {

	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	result := breaking_the_records.BreakingRecords(scores)
	fmt.Println(result)
	testCaseValue := []string{"S;M;plasticCup()", "C;V;mobile phone", "C;C;coffee machine", "S;C;LargeSoftwareBook", "C;M;white sheet of paper", "S;V;pictureFrame"}
	for _, testWord := range testCaseValue {
		camel_case_four.CamelCaseFour(testWord)
	}

}
