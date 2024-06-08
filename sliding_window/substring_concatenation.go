package slidingwindow

import (
	timetrack "main/initial/time_track"
	"time"
)

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	defer timetrack.TimeTrack(time.Now(), "findSubstring")
    indexes := make([]int,0)
	wordLen := len(words[0])
	startIndex := 0
	endIndex := wordLen * len(words)

	wordsMap := make(map[string]int)

	for _, v := range words {
		wordsMap[v] += 1
	}

	for endIndex <= len(s) {
		wordsMapCopy := make(map[string]int)
		for k, v := range wordsMap {
			wordsMapCopy[k] = v
		}
		window := s[startIndex:endIndex]
		for i := 0; i < len(window); i += wordLen {
			wordEnd := i + wordLen
			wordsMapCopy[window[i:wordEnd]] += 1
		}
		goodMap := true
		for k, v := range wordsMap {
			if wordsMapCopy[k] != v+v {
				goodMap = false
				break
			}
		}
		if goodMap {
			indexes = append(indexes, startIndex)
		}
		wordsMapCopy = wordsMap
		startIndex += 1
		endIndex += 1
	}

	return indexes
}