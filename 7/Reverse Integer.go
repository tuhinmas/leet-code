package main

import (
	"fmt"
	"math"
)

func main() {
	x := 123
	x = 1534236469
	reverse(x)
}

func reverse(x int) int {
	reversed := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	fmt.Println(math.MaxInt32)
	fmt.Println(reversed * sign)

	if reversed*sign < math.MinInt32 || reversed*sign > math.MaxInt32 {
		return 0
	}

	return reversed * sign
}
