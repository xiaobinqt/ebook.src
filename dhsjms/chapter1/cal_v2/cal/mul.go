package cal

import "fmt"

type OperationMul struct {
	Operation
}

func (o *Operation) GetNumberA() float64 {
	return o.NumberA
}

func (o *Operation) SetNumberA(a float64) {
	o.NumberA = a
}

func (o *Operation) GetNumberB() float64 {
	return o.NumberB
}

func (o *Operation) SetNumberB(b float64) {
	o.NumberB = b
}

func (o *OperationMul) GetResult() (ret string) {
	return fmt.Sprint(o.NumberA * o.NumberB)
}
