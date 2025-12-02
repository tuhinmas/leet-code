package main

func main() {
	s := "PAYPALISHIRING"
	// s = "AB"
	// s = "A"
	numRows := 4
	// convert(s, numRows)
	coverChatGPT(s, numRows)
}

func convert(s string, numRows int) string {
	var matrix = make([][]string, numRows)
	var possibleColumn = 1
	var result string
	cycle := 1

	if len(s) > numRows {
		if numRows == 1 {
			cycle = len(s)
		} else {
			cycle = len(s) / ((2 * numRows) - 2)
		}
	}

	remaind := 0
	if ((2 * numRows) - 2) > 0 {
		remaind = len(s) % ((2 * numRows) - 2)
	}

	if (numRows - 1) > 0 {
		possibleColumn = cycle * (numRows - 1)
	}

	if numRows == 1 {
		possibleColumn = len(s)
	}

	if remaind > numRows {
		possibleColumn += (remaind - numRows) + 1
		cycle += 1
	} else if remaind > 0 {
		possibleColumn += 1
		cycle += 1
	}

	// fmt.Println("cycle", cycle)
	// fmt.Println("possibleColumn", possibleColumn)
	for i := numRows - 1; i >= 0; i-- {
		matrix[i] = make([]string, possibleColumn)
	}

	col := 0
	for x := 0; x < int(cycle); x++ {

		leftUndex := x*(2*numRows-2) + numRows - 1
		if (numRows - 1) == 0 {
			leftUndex = x
		}
		rightIndex := x*(2*numRows-2) + numRows - 1
		for i := numRows - 1; i >= 0; i-- {

			/* zig */
			if leftUndex < len(s) {
				// fmt.Println(leftUndex, "-", rightIndex)
				// fmt.Println("col:", col)
				// fmt.Println("cahr:", string(s[leftUndex]))
				matrix[i][col] = string(s[leftUndex])
			}

			/* zag */
			if rightIndex < len(s) && numRows > 2 {
				matrix[i][col+(numRows-i-1)] = string(s[rightIndex])
			}

			/* increase index string */
			rightIndex++
			leftUndex--
		}

		/* first cycle increase */
		if (numRows - 1) > 0 {
			col += numRows - 1
		} else {
			col++
		}
	}

	for i := 0; i < numRows; i++ {
		// fmt.Println(matrix[i])
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			result += matrix[i][j]
		}
	}

	// fmt.Println("result", result)
	return result
}

func coverChatGPT(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	for _, char := range s {
		rows[curRow] += string(char)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
