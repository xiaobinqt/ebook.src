package main

import (
	"encoding/hex"
	"fmt"
)

func f1() []byte {
	v := []byte{0x33, 0x32, 0x34, 0x35, 0xc8, 0x56, 0xc9, 0x56, 0xc9, 0x56}
	for i := 0; i < len(v); i++ {
		v[i] = v[i] - 0x33
	}
	fmt.Println(v)
	return v
}

func main() {

	fmt.Println(string(f1()))

	//return
	// 注意"Hello"与"encodedStr"不相等，encodedStr是用字符串来表示16进制
	//src := []byte("Hello")
	//encodedStr := hex.EncodeToString(src)
	//// [72 101 108 108 111]
	//fmt.Println(src)
	//// 48656c6c6f -> 48(4*16+8=72) 65(6*16+5=101) 6c 6c 6f
	//fmt.Println("33323435c856c956c956", encodedStr)

	test, _ := hex.DecodeString("025512149351503515035")
	fmt.Println(string(test))

}
