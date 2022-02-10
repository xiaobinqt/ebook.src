package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	filepath := fmt.Sprintf("%s/dev/eval/read_big_file/read.txt", wd)
	ReadBigFileByByte(filepath)
}

func ReadBigFileByByte(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("can't opened this file")
		return
	}
	defer f.Close()

	s := make([]byte, 0, 2) // 每次只读2个字节
	for {
		n, err := f.Read(s[len(s):cap(s)])
		fmt.Println("err", err, "n", n)

		// 把读出来的字符写到切片里去
		s = s[:len(s)+n]
		fmt.Println(string(s))

		if err != nil {
			if err == io.EOF { // 读完了
				return
			}
			fmt.Println("read err", err.Error())
			return
		}
	}

}
