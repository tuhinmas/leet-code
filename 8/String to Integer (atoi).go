package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	s := " -042"   // -42
	s = "042"      // 42
	s = "42"       // 42
	s = "1337c0d3" // 1337
	// s = "0-1"           // 0
	// s = "words and 987" // 0
	// s = "-91283472332"  // -2147483648
	// s = "   +0 123"     // 0
	// myAtoi(s)
	myAtoiGPT(s)
}

func myAtoi(s string) int {
	result := ""
	lead := true
	for i := 0; i < len(s); i++ {

		/* ignoring any leading white space*/
		if string(s[i]) == " " && lead {
			continue
		}

		// if !unicode.IsDigit(rune(s[i])) && (string(s[i]) != "-" || string(s[i]) != "+") {
		if !unicode.IsDigit(rune(s[i])) && lead && string(s[i]) != "-" && string(s[i]) != "+" {
			break
		}

		if !unicode.IsDigit(rune(s[i])) {
			if !lead {
				break
			} else if string(s[i]) != "-" && string(s[i]) != "+" {
				break
			}
		}

		result = result + string(s[i])
		lead = false
		fmt.Println("result:", result)
	}

	if result == "" {
		return 0
	}

	final, err := strconv.ParseInt(result, 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	if final < math.MinInt32 {
		final = math.MinInt32
	} else if final > math.MaxInt32 {
		final = math.MaxInt32
	}

	fmt.Println(int(final))

	return int(final)
}

func myAtoiGPT(s string) int {
	const (
		intMax = 1<<31 - 1 // 2147483647
		intMin = -1 << 31  // -2147483648
	)

	i, n := 0, len(s)
	// Skip leading whitespace
	for i < n && s[i] == ' ' {
		i++
	}

	fmt.Println("n:", n)
	fmt.Println("i:", i)
	// Handle optional sign
	sign := 1
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Parse digits and build the number
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Check for overflow/underflow before adding the digit
		if result > (intMax-digit)/10 {
			if sign == 1 {
				return intMax
			}
			return intMin
		}
		result = result*10 + digit
		fmt.Println(int(s[i]), ":", int(s[i]-'0'))
		i++
	}

	fmt.Println(result * sign)
	return result * sign
}
