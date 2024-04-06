package calc_engine

type CalculatorEngine interface {
	ReadInput(input string) error
	ProceedInput(value float64, command string, arg float64) float64
	repeat(result float64, input float64, commands []Command) (float64, error)
}
