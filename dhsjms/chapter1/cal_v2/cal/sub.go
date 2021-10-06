package cal

import "fmt"

type OperationSub struct {
	Operation
}

func (o *OperationSub) GetNumberA() float64 {
	return o.NumberA
}

func (o *OperationSub) SetNumberA(a float64) {
	o.NumberA = a
}

func (o *OperationSub) GetNumberB() float64 {
	return o.NumberB
}

func (o *OperationSub) SetNumberB(b float64) {
	o.NumberB = b
}

func (o *OperationSub) GetResult() (ret string) {
	return fmt.Sprint(o.NumberA - o.NumberB)
}
