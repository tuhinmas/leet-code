package main

import "fmt"

func main() {
	s := "abcabcbb"
	lengthOfLongestSubstringV2(s)
	// lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstringV2(s string) int {
	var curChar *string
	length := 0

	for i := 0; i < len(s); i++ {
		x := string(s[i])
		curChar = &x

	outerLoop:
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				break
			}

			for x := 0; x < len(*curChar); x++ {
				if (*curChar)[x] == s[j] {
					break outerLoop
				}
			}

			*curChar += string(s[j])
		}

		if len(*curChar) > length {
			length = len(*curChar)
		}
	}
	fmt.Println(length)

	return length
}

func lengthOfLongestSubstring(s string) int {
	var curChar string
	var subStr []string
	for i := 0; i < len(s); i++ {
		curChar = string(s[i])

	outerLoop:
		for j := i + 1; j < len(s); j++ {
			for x := 0; x < len(curChar); x++ {
				if string(curChar[x]) == string(s[j]) {
					break outerLoop
				}
			}
			if string(s[i]) == string(s[j]) {
				break
			}

			curChar += string(s[j])
		}
		subStr = append(subStr, curChar)
	}

	length := 0
	for _, char := range subStr {
		if len(char) > length {
			length = len(char)
		}
	}

	fmt.Print(length)
	return length
}
