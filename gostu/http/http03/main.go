package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0)")

	resp, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
