package gee

import (
	"fmt"
	"testing"
)

func Test_parsePattern(t *testing.T) {
	x := parsePattern("/hello/:xxx")
	fmt.Println(x)
}
