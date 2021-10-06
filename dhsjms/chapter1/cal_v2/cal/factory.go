package cal

type OperationFactory struct {
}

func (o *OperationFactory) CreateOperate(operate string) OperationInterface {
	var oper OperationInterface
	switch operate {
	case "+":
		oper = new(OperationAdd)
	case "-":
		oper = new(OperationSub)
	case "*":
		oper = new(OperationMul)
	case "/":
		oper = new(OperationDiv)
	}

	return oper
}
