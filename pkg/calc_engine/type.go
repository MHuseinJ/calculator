package calc_engine

import (
	"calculator/pkg/math_engine"
	"calculator/pkg/util"
)

type Command struct {
	Cmd string
	Arg float64
}
type CalculatorInstance struct {
	argCommands   []string
	noArgCommands []string
	mathEngine    math_engine.MathEngine
	Commands      []Command
	lastValue     float64
	util          util.Util
}

func NewCal(engine math_engine.MathEngine, util util.Util) CalculatorEngine {
	return &CalculatorInstance{
		[]string{"add", "subtract", "multiply", "divide", "repeat"},
		[]string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
		engine,
		[]Command{},
		0.0,
		util,
	}
}
