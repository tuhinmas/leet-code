package main

import "fmt"

func main() {
	i := 49
	fmt.Println(intToRoman(i))
}

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for i, v := range values {
		for num >= v {
			res += symbols[i]
			num -= v
		}
	}

	return res
}
