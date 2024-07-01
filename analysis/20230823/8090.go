package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("I am 8090"))
	})
	http.ListenAndServe("0.0.0.0:8090", nil)
}
