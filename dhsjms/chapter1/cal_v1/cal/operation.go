package cal

import "fmt"

type Operation struct {
}

func (o *Operation) GetResult(a, b float64, operate string) (ret string) {
	switch operate {
	case "+":
		ret = fmt.Sprint(a + b)
	case "-":
		ret = fmt.Sprint(a - b)
	case "*":
		ret = fmt.Sprint(a * b)
	case "/":
		ret = fmt.Sprint(a / b)
	}

	return ret
}
