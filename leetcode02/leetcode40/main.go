package main

import "fmt"

/**
宝石与石头
*/
func numJewelsInStones(jewels string, stones string) int {
	hash := make(map[rune]struct{})
	for _, ch := range []rune(jewels) {
		hash[ch] = struct{}{}
	}
	res := 0
	for _, ch := range []rune(stones) {
		if _, exists := hash[ch]; exists {
			res += 1
		}
	}
	return res
}

func main() {
	jewels := "z"
	stones := "ZZ"
	inStones := numJewelsInStones(jewels, stones)
	fmt.Println(inStones)
}
