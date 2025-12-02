package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	l1 := createL1()
	l2 := createL2()

	addTwoNumbers(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createL1() *ListNode {
	var l1 = []int{
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	}
	var node ListNode
	var nodes []ListNode
	for i := len(l1) - 1; i >= 0; i-- {
		node = ListNode{
			Val:  l1[i],
			Next: nil,
		}

		nodes = append([]ListNode{node}, nodes...)
	}

	if len(nodes) > 1 {
		for i := 0; i <= len(nodes)-2; i++ {
			nodes[i].Next = &nodes[i+1]
		}
	}

	return &nodes[0]
}

func createL2() *ListNode {
	var l1 = []int{
		5, 6, 4,
	}
	var node ListNode
	var nodes []ListNode
	for i := len(l1) - 1; i >= 0; i-- {
		node = ListNode{
			Val:  l1[i],
			Next: nil,
		}

		nodes = append([]ListNode{node}, nodes...)
	}

	if len(nodes) > 1 {
		for i := 0; i <= len(nodes)-2; i++ {
			nodes[i].Next = &nodes[i+1]
		}
	}

	return &nodes[0]
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node1 []int
	var node2 []int
	current := l1
	for current != nil {
		node1 = append(node1, current.Val)
		current = current.Next
	}

	current = l2
	for current != nil {
		node2 = append(node2, current.Val)
		current = current.Next
	}

	/* create sum integer */
	var sum1 string
	for i := len(node1) - 1; i >= 0; i-- {
		sum1 += strconv.Itoa(node1[i])
	}

	var sum2 string
	for i := len(node2) - 1; i >= 0; i-- {
		sum2 += strconv.Itoa(node2[i])
	}

	sum1 = strings.TrimSpace(sum1)
	val := new(big.Int)
	val2 := new(big.Int)

	// Attempt to parse the string
	val, ok := val.SetString(sum1, 10)
	if !ok {
		fmt.Println("Error: Invalid big integer format.")
	}

	// result += val
	val2, ok = val2.SetString(sum2, 10)
	if !ok {
		fmt.Println("Error: Invalid big integer format.")
	}

	result := new(big.Int).Set(val).Add(val, val2) // sum = a + b

	/* convert result to string and slice */
	strResult := result.String()
	// strResult := "111"

	var sliceOfStr []int

	for i := 0; i < len(strResult); i++ {
		num, _ := strconv.Atoi(string(strResult[i]))
		sliceOfStr = append(sliceOfStr, num)
	}

	/* create reverse of slice */
	sliceOfStrReverse := make([]int, len(sliceOfStr))
	for i := 0; i < len(sliceOfStr); i++ {
		sliceOfStrReverse[i] = sliceOfStr[len(sliceOfStr)-1-i]
	}

	/* create link node*/
	var node ListNode
	var nodes []ListNode
	for i := len(sliceOfStrReverse) - 1; i >= 0; i-- {
		node = ListNode{
			Val:  sliceOfStrReverse[i],
			Next: nil,
		}

		nodes = append([]ListNode{node}, nodes...)
	}

	if len(nodes) > 1 {
		for i := 0; i <= len(nodes)-2; i++ {
			nodes[i].Next = &nodes[i+1]

			// fmt.Print(nodes[i].Val, ",")
		}
	}

	return &nodes[0]
}
