package lonely_integer

func Lonelyinteger(a []int32) int32 {
	var result int32
	for _, value := range a {
		result ^= value
	}
	return result
}
