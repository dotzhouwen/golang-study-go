package main

import (
	"fmt"
	"strconv"
)

/**
单词的唯一缩写
*/
type ValidWordAbbr struct {
	hash map[string][]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	dicHash := make(map[string][]string, 0)
	for _, str := range dictionary {
		chars := []rune(str)
		if len(chars) > 2 {
			num := len(chars) - 2
			first := chars[0]
			last := chars[len(chars)-1]
			s := string(first) + strconv.Itoa(num) + string(last)
			dicHash[s] = append(dicHash[s], str)
		} else {
			dicHash[str] = append(dicHash[str], str)
		}
	}
	return ValidWordAbbr{
		hash: dicHash,
	}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	chars := []rune(word)
	var key string
	if len(chars) > 2 {
		num := len(chars) - 2
		first := chars[0]
		last := chars[len(chars)-1]
		key = string(first) + strconv.Itoa(num) + string(last)
	} else {
		key = word
	}

	if _, exists := this.hash[key]; exists {
		for _, w := range this.hash[key] {
			if w != word {
				return false
			}
		}
		return true
	} else {
		return true
	}
}

func main() {
	dictionary := []string{"deer", "door", "cake", "card"}
	validWordAbbr := Constructor(dictionary)
	unique := validWordAbbr.IsUnique("at")
	fmt.Println(unique)
}
