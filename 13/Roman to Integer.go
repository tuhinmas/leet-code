package main

import (
	"fmt"
	"strings"
)

func main() {
	// i := "LVIII"
	i := "XLIV"
	fmt.Println(romanToInt(i))
}

func romanToInt(rome string) int {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := 0
	for i, v := range symbols {
		for strings.HasPrefix(rome, v) {
			res += values[i]
			rome = rome[len(v):]

		}
	}

	return res
}
