package arrays

import (
	"sort"
)

func MaximumPopulation(logs [][]int) int {
	alivedMap := make(map[int]int)
	for _, log := range logs {
		for birth_i := log[0]; birth_i < log[1]; birth_i++ {
			alivedMap[birth_i] += 1
		}
	}
    return maxKey(alivedMap)
}

func maxKey(m map[int]int) int {
	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var maxCount int
	var maxNumber int 
	for _, k := range keys {
		if m[k] > maxCount {
			maxCount = m[k]
			maxNumber = k
		}
	}
	return maxNumber
}