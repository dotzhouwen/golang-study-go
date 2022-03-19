package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	doGetWithParam()
}

func doJsonPost() {
	bodyJson, _ := json.Marshal(map[string]interface{}{
		"key": "value",
	})
	r, _ := http.DefaultClient.Post("http://httpbin.org/post", "application/json", strings.NewReader(string(bodyJson)))
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}

func doGetWithParam() {
	r, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	params := make(url.Values)
	params.Add("key1", "value1")
	params.Add("key2", "value2")

	r.URL.RawQuery = params.Encode()
	res, _ := http.DefaultClient.Do(r)

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s", body)
}

func doGet() {
	r, _ := http.Get("http://httpbin.org/get")
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}

func doPut() {
	request, _ := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	r, _ := http.DefaultClient.Do(request)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}

func doPost() {
	request, _ := http.NewRequest(http.MethodPost, "http://httpbin.org/post", nil)
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}
