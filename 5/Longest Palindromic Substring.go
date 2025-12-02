package main

import "fmt"

func main() {
	s := "babad"
	// s = "a"
	// s = "ac"
	// s = "aacabdkacaa"
	// s = "cstgvkbrxacmpxdxxktktvpdzcuxmnhvuxdgsmskgeeawzeikhtmhdvnccbrgifpzmiuytfmeyfoxsntrdtxeuxcqsndexbgfxnthqwveujqzemloooyddparbjcuiwpopjwvvmwblsamkhjhlnoxinkpsempexmipifsfwzxbetgvfnkngzxcpizwctpdlpngjpyovmjllxfiwktghkxvyelwjwdztujmunijfsfdvmhgqhlpouewgyznphlmccjmqaqncwbeqheohibafqfunfywmrvqvjygjwqoclijwkcfiuaiymeagdbwyejnvtosxylptbtyoahfzfmwzrkhzdamknleroffmsqcaryibamgdpcumlhrblypddzhaxfeztokgogzgvpfvlmetiwsamhdidmvxavleryfyakendwrbslcavlqkerrnvpuzhdgwzuyorxzbkzhxhpbvkflgxouvaavxrdzsjlgrmogzvlhhdidldsxqhrqlryaanffhxnutcycnczuedtrwcnfiqrtoafvdfnfhxhyjivzalozrbrajboecfyalisxxanduzraqdrbzsbkobaieqpzcawrunxucypqyjnmrlrlivrrwwhdpekeelsphhunzbhkkejvqfopjsuholutgmtnleqdrntbqgrobnuhqpdkbjtikijkdiwqvnxgajaaqgswrdamzv"
	longestPalindrome(s)
	longestPalindromeV2(s)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {

		// Expand for odd-length palindromes
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1

			}
			left--
			right++
		}

		// Expand for even-length palindromes
		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	return s[start : start+maxLen]
}

func longestPalindromeV2(s string) string {
	var curr string

	for i := 0; i < len(s); i++ {

		palindrom := string(s[i])
		check := string(s[i])

		for j := i + 1; j < len(s); j++ {

			check += string(s[j])
			is_plindrom := true
			for x := 0; x < len(check); x++ {
				if check[x] != check[len(check)-x-1] {
					// if len(check) > len(curr) {
					// 	curr = check
					// }
					is_plindrom = false
					break
				}
				// else {
				// 	is_plindrom = false
				// 	break
				// }
			}

			if len(check) > len(curr) && is_plindrom {
				curr = check
			}
		}

		if len(curr) == 0 {
			curr = palindrom
		}
	}

	fmt.Println(curr)
	return curr
}
