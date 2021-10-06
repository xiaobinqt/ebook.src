package cal

type Operation struct {
	NumberA float64
	NumberB float64
}

type OperationInterface interface {
	GetResult() (ret string)
	SetNumberA(float64)
	SetNumberB(float64)
	GetNumberB() float64
	GetNumberA() float64
}
