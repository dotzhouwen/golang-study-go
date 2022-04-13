package main

import "fmt"

/*
424. 替换后的最长重复字符
*/
func characterReplacement(s string, k int) int {
	chars := []rune(s)
	cnt := make([]int, 26)

	ans := 0
	left := 0

	for right := 0; right < len(chars); right++ {
		cnt[chars[right]-'A']++
		for !check(cnt, k) {
			cnt[chars[left]-'A']--
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func check(cnt []int, k int) bool {
	max := 0
	sum := 0
	for i := 0; i < len(cnt); i++ {
		if cnt[i] > max {
			max = cnt[i]
		}
		sum += cnt[i]
	}
	return sum-max <= k
}

func main() {
	s := "AAABAAABBB"
	k := 1
	replacement := characterReplacement(s, k)
	fmt.Println(replacement)
}
