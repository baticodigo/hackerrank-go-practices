package main

import (
	"fmt"
	"hackerrank-go-practices/breaking_the_records"
)

func main() {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	result := breaking_the_records.BreakingRecords(scores)
	fmt.Println(result)
}
