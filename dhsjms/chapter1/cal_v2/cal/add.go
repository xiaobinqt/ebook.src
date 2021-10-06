package cal

import "fmt"

type OperationAdd struct {
	Operation
}

func (o *OperationAdd) GetNumberA() float64 {
	return o.NumberA
}

func (o *OperationAdd) SetNumberA(a float64) {
	o.NumberA = a
}

func (o *OperationAdd) GetNumberB() float64 {
	return o.NumberB
}

func (o *OperationAdd) SetNumberB(b float64) {
	o.NumberB = b
}

func (o *OperationAdd) GetResult() (ret string) {
	return fmt.Sprint(o.NumberA + o.NumberB)
}
