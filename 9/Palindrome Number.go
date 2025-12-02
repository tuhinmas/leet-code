package main

import (
	"fmt"
	"math"
)

func main() {
	i := 121
	// i = 192
	// i = 1901
	// i = 2332
	// i = 2331
	isPalindrome(i)
	isPalindromeGPT(i)
}

func isPalindrome(x int) bool {

	/* negative always false palindrome */
	if x*-1 > 0 {
		return false
	}
	length := 0
	temp := x
	for temp > 0 {
		temp /= int(math.Pow(10, 1))
		length++
	}

	/* odd */
	if length%2 == 1 {
		start := length/2 + 1

		left, right := start, start
		for left <= length && right >= 0 {
			temp = x
			first := 0
			for i := 0; i < left; i++ {
				first = temp % int(math.Pow(10, 1))
				temp = temp / int(math.Pow(10, 1))
			}

			temp = x
			second := 0
			for i := 0; i < right; i++ {
				second = temp % int(math.Pow(10, 1))
				temp = temp / int(math.Pow(10, 1))
			}

			if first != second {
				return false
			}
			left++
			right--
		}
	} else {
		start := length / 2

		left, right := start+1, start
		for left <= length && right >= 0 {
			temp = x
			first := 0
			for i := 0; i < left; i++ {
				first = temp % int(math.Pow(10, 1))
				temp = temp / int(math.Pow(10, 1))
			}

			temp = x
			second := 0
			for i := 0; i < right; i++ {
				second = temp % int(math.Pow(10, 1))
				temp = temp / int(math.Pow(10, 1))
			}

			if first != second {
				return false
			}

			left++
			right--
		}
	}

	return true
}

func isPalindromeGPT(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10 // Add the last digit to reversedHalf
		x /= 10                               // Remove the last digit from x
		fmt.Println("revers:", reversedHalf, "x:", x)
	}

	// Check if the original half matches the reversed half
	return x == reversedHalf || x == reversedHalf/10
}
