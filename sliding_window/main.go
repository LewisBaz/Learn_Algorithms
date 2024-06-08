package slidingwindow

import "fmt"

func Run() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo","bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar","foo","the"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo","barr","wing","ding","wing"}))
}