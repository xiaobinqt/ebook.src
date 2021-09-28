package main

import "net/http"

func main() {
	req, _ := http.NewRequest(http.MethodGet,
		"http://www.baidu.com", nil)
	req.Header.Add("xx", "ss")

	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	//http.Transport{}
	//http.RoundTripper()
	client.Do(req)
}
