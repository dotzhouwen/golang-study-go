package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "https://api.github.com/events", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {
		r.Body.Close()
	}()

	body, err := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}

func get() {
	resp, err := http.Get("https://api.github.com/events")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}