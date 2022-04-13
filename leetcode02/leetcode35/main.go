package main

import "fmt"

/**
两个列表的最小索引总和
*/
func findRestaurant(list1 []string, list2 []string) []string {
	hash1 := make(map[string]int, 0)
	for k, v := range list1 {
		hash1[v] = k
	}

	minIndex := len(list1) + len(list2)
	res := make([]string, 0)

	for i, v := range list2 {
		if ix, exists := hash1[v]; exists {
			if i+ix < minIndex {
				res = []string{list2[i]}
				minIndex = i + ix
			} else if i+ix == minIndex {
				res = append(res, list2[i])
			}
		}
	}
	return res
}

func main() {

	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"KFC", "Shogun", "Burger King"}
	restaurant := findRestaurant(list1, list2)

	fmt.Println(restaurant)
}
