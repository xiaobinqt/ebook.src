package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("err == ", err.Error())
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
