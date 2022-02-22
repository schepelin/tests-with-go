package calc

type Calc struct{}

func NewCalc() *Calc {
	return &Calc{}
}
func (*Calc) Multiply(a, b int) int {
	return a * b
}

func Sum(a, b int) int {
	return a + b
}
