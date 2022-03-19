package main

import (
	"fmt"
	"net/url"
)

func main() {
	//join := filepath.Join("https://b1.szjal.cn", "/dabd", "a.txt")
	parse, _ := url.Parse("https://b1.szjal.cn")
	fmt.Println(parse)
}


