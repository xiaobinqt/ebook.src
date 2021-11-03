package main

import (
	"fmt"
	"reflect"
)

type X struct {
}

// 固定参数
func (X) Test(x, y int) (int, error) {
	return x + y, fmt.Errorf("err:%d", x+y)
}

// 变参
func (X) Format(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}

func main() {
	var a X

	v := reflect.ValueOf(&a)
	m := v.MethodByName("Test")

	in := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	}

	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}

	// 变参
	m = v.MethodByName("Format")

	out = m.Call([]reflect.Value{ // 所有参数都需要处理
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf("x"),
		reflect.ValueOf(100),
	})
	fmt.Println(out)

	out = m.CallSlice([]reflect.Value{
		reflect.ValueOf("%s=%d"),
		reflect.ValueOf([]interface{}{ // 仅一个即可
			"x", 100,
		}),
	})

	fmt.Println(out)
}
