package main

import "net/http"

// curl -k https://127.0.0.1:8080
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("tls success"))
	})

	http.ListenAndServeTLS(":8080", "xsw.crt", "xsw.key", nil)
}
