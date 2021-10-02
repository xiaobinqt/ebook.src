package main

import "net/http"

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8091",
		"xsw.crt", "xsw.key", h)
}
