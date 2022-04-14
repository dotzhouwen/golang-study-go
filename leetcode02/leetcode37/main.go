package main

import "fmt"

/**
249. 移位字符串分组
*/
func groupStrings(strings []string) [][]string {
	hash := make(map[string][]string, 0)
	for _, str := range strings {
		chars := []rune(str)
		first := chars[0]
		keyChs := make([]rune, 0)
		for _, ch := range chars {
			keyChs = append(keyChs, (ch-first+26)%26)
		}
		key := string(keyChs)
		hash[key] = append(hash[key], str)
	}
	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}

func main() {
	strings := []string{"zab", "abc", "bcd", "acef", "xyz", "az", "ba", "a", "z", "efg", "opq"}
	i := groupStrings(strings)
	fmt.Println(i)
}
