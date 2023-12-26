package main

import (
	"fmt"
	binarysearch "main/binary_search"
	digitsalgs "main/digits_algs"
	oneway "main/linked_lists/one_way"
	"main/simple"
)

var (
	hundredSlice = createSlice(100)
	thousandSlice = createSlice(1000)
)


func main() {
	fmt.Println(simple.GCD(4851, 3003))
	fmt.Println(simple.Randomize_array(hundredSlice))

	fmt.Println(digitsalgs.DiceThatTossACoin())
	fmt.Println(digitsalgs.EqualValuesDice())
	fmt.Println(digitsalgs.ChooseRandomFrom(thousandSlice, 5))
	fmt.Println(digitsalgs.Poker_card_giver(5))

	oneway.Show()

	fmt.Println(binarysearch.Search([]int{-1,0,3,5,9,12}, 3))
}

func createSlice(count int) []int {
	mySlice := make([]int, count)
	for i := 0; i < count; i++ {
		mySlice[i] = i
	}
	return mySlice
}