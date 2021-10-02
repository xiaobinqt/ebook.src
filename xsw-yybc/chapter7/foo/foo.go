/*
11111111111

2222222222222222222

33333333333333
*/
package foo

// Foo xxxxxxxx
// 哈哈哈哈哈哈哈哈
func Foo(a, b int) (ret int, err error) {
	if a > b {
		return a, nil
	}

	return b, nil
}

// BUG(jack): #1 第一号 bug
// BUG(tom): #2 第二号 bug
