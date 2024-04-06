package main

import (
	"calculator/pkg/calc_engine"
	"calculator/pkg/math_engine"
	"calculator/pkg/util"
	"fmt"
)

func main() {
	fmt.Println("Welocome to simple calculator")
	utilCalc := util.NewUtil()
	engine := math_engine.NewMathEngineForCalculator()
	calc := calc_engine.NewCal(engine, utilCalc)
	input1 := ""
	var err error

	for input1 != "exit" {
		fmt.Print(">")
		input1, err = utilCalc.ReadInputString()
		if err != nil {
			fmt.Println("sorry, invalid command")
		}
		err = calc.ReadInput(input1)
		if err != nil {
			fmt.Println("sorry, invalid argument")
		}
	}
}
