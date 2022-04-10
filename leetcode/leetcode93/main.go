package main

import "fmt"

/**
1423. 可获得的最大点数
*/
func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		sum += cardPoints[i]
	}

	temp := sum
	j := len(cardPoints) - k

	for i := 0; i < k; i++ {
		temp -= cardPoints[j]
		temp += cardPoints[i]
		if temp > sum {
			sum = temp
		}
		j++
	}
	return sum
}

func main() {
	cardPoints := []int{1, 79, 80, 1, 1, 1, 200, 1}
	k := 3
	score := maxScore(cardPoints, k)
	fmt.Println(score)
}
