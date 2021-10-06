package cal

import "fmt"

type OperationDiv struct {
	Operation
}

func (o *OperationDiv) GetNumberA() float64 {
	return o.NumberA
}

func (o *OperationDiv) SetNumberA(a float64) {
	o.NumberA = a
}

func (o *OperationDiv) GetNumberB() float64 {
	return o.NumberB
}

func (o *OperationDiv) SetNumberB(b float64) {
	o.NumberB = b
}

func (o *OperationDiv) GetResult() (ret string) {
	return fmt.Sprint(o.NumberA / o.NumberB)
}
