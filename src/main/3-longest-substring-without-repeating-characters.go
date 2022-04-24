package main

import "fmt"

/**
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}

func lengthOfLongestSubstring(s string) int {

	set := make(map[uint8]int)

	max := 0
	length := 0
	for i := 0; i < len(s); i++ {
		value, ok := set[s[i]]
		if !ok {
			length++
			if max < length {
				max = length
			}
		} else {
			length = 1
			set = make(map[uint8]int)
			i = value + 1
		}
		set[s[i]] = i
	}

	return max

}
