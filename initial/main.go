package main

import (
	"fmt"
	// "main/arrays"
	stringsf "main/strings"

	// binarysearch "main/binary_search"
	// digitsalgs "main/digits_algs"
	// oneway "main/linked_lists/one_way"
	hascycle "main/linked_lists/one_way/has_cycle"
	// "main/simple"
)

var (
	hundredSlice  = createSlice(100)
	thousandSlice = createSlice(1000)
)

func main() {
	// fmt.Println(simple.GCD(4851, 3003))
	// fmt.Println(simple.Randomize_array(hundredSlice))

	// fmt.Println(digitsalgs.DiceThatTossACoin())
	// fmt.Println(digitsalgs.EqualValuesDice())
	// fmt.Println(digitsalgs.ChooseRandomFrom(thousandSlice, 5))
	// fmt.Println(digitsalgs.Poker_card_giver(5))

	// oneway.Show()

	// fmt.Println(binarysearch.Search([]int{-1,0,3,5,9,12}, 3))
	// nodeCycle()

	// arrays.MergeArraysInAsc([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	// arrays.MergeArrays([]int{1,2,4,6}, 4, []int{1,2,5,6,3}, 5)

	// fmt.Println(stringsf.LengthOfLastWord("a "))

	// digitsalgs.IntToRoman(2745)
	// digitsalgs.RomanToInt("LVIII")
	// res := digitsalgs.IsPalindrome(0)
	// fmt.Println(digitsalgs.MyPow(-1, 2147483647))
	// fmt.Println(digitsalgs.MyPow(1.0000000000001, -2147483648))
	// fmt.Println(digitsalgs.MyPow(2, 12))

	fmt.Println(stringsf.IsIsomorphic("badc", "baba"))
}

func createSlice(count int) []int {
	mySlice := make([]int, count)
	for i := 0; i < count; i++ {
		mySlice[i] = i
	}
	return mySlice
}

func nodeCycle() {
	var node = &hascycle.ListNode{Val: 1}
	var nextNode = &hascycle.ListNode{Val: 1}
	node.Next = nextNode
	nextNode.Next = node
	var hasC = hascycle.HasCycle(node)
	var head = hascycle.DetectCycle(node)
	fmt.Println(hasC)
	fmt.Println(head)
}
