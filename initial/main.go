package main

import (
	"fmt"
	"main/arrays"
	// "main/sorting"
	// stringsf "main/strings"

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
	// arrays.TwoSum([]int{0, 0, 3, 4}, 0)

	// fmt.Println(stringsf.LengthOfLastWord("a "))

	// digitsalgs.IntToRoman(2745)
	// digitsalgs.RomanToInt("LVIII")
	// res := digitsalgs.IsPalindrome(0)
	// fmt.Println(digitsalgs.MyPow(-1, 2147483647))
	// fmt.Println(digitsalgs.MyPow(1.0000000000001, -2147483648))
	// fmt.Println(digitsalgs.MyPow(2, 12))

	// fmt.Println(stringsf.IsIsomorphic("badc", "baba"))

	// sorting.SortColors([]int{2,0,1,4,6,3,2,1,1,0,5,6,0,3})
	// sorting.FindTheDifference("abcd", "abcde")

	// arrays.MajorityElement([]int{3,2,3})
	// fmt.Println(arrays.MaximumPopulation([][]int{{1950, 1961}, {1960,1971}, {1970,1981}}))
	// fmt.Println(arrays.MaxDistance([]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}))
	// fmt.Println(arrays.MaxDistance([]int{2, 2, 2}, []int{10, 10, 1}))
	// fmt.Println(arrays.MaxDistance([]int{30, 29, 19, 5}, []int{25,25,25,25,25}))
	fmt.Println(arrays.MaxDistance([]int{9820,8937,7936,4855,4830,4122,2327,1342,1167,815,414}, []int{9889,9817,9800,9777,9670,9646,9304,8977,8974,8802,8626,8622,8456}))
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
