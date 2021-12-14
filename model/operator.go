package model

type Operator int

const (
	None Operator = iota
	Plus
	Minus
	Multiply
	Divide
	Equals
)

func (o Operator) String() string {
	switch o {
	case None:
		return "None"
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Multiply:
		return "Multiply"
	case Divide:
		return "Divide"
	case Equals:
		return "Equals"
	}
	panic("unknown operator")
}
