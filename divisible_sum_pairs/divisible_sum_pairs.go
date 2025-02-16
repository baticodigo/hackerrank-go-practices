package divisible_sum_pairs

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	countValidPairs := 0
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			sum1, sum2 := ar[i], ar[j]
			if ((sum1 + sum2) % k) == 0 {
				countValidPairs++
			}
		}
	}
	return int32(countValidPairs)
}
