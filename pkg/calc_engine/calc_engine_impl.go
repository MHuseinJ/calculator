package calc_engine

import (
	"calculator/pkg/util"
	"fmt"
)

func (cal *CalculatorInstance) ReadInput(input string) error {
	arg, err := examineInput(cal.util, cal.argCommands, cal.noArgCommands, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	cal.lastValue = cal.ProceedInput(cal.lastValue, input, arg)
	if input != "exit" {
		fmt.Printf("%.1f\n", cal.lastValue)
	}
	return err
}

func (cal *CalculatorInstance) ProceedInput(value float64, command string, arg float64) float64 {
	result := value
	switch command {
	case "add":
		{
			result = cal.mathEngine.Add(value, arg)
			cmd := Command{
				Cmd: "add",
				Arg: arg,
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "subtract":
		{
			result = cal.mathEngine.Subtract(value, arg)
			cmd := Command{
				Cmd: "subtract",
				Arg: arg,
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "multiply":
		{
			result = cal.mathEngine.Multiply(value, arg)
			cmd := Command{
				Cmd: "multiply",
				Arg: arg,
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "divide":
		{
			result = cal.mathEngine.Divide(value, arg)
			cmd := Command{
				Cmd: "divide",
				Arg: arg,
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "sqr":
		{
			result = cal.mathEngine.Sqr(value)
			cmd := Command{
				Cmd: "sqr",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "sqrt":
		{
			result = cal.mathEngine.Sqrt(value)
			cmd := Command{
				Cmd: "sqrt",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "cube":
		{
			result = cal.mathEngine.Cube(value)
			cmd := Command{
				Cmd: "cube",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "cubert":
		{
			result = cal.mathEngine.Cubert(value)
			cmd := Command{
				Cmd: "cubert",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "neg":
		{
			result = cal.mathEngine.Neg(value)
			cmd := Command{
				Cmd: "neg",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "abs":
		{
			result = cal.mathEngine.Abs(value)
			cmd := Command{
				Cmd: "abs",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "cancel":
		{
			result = 0.0
			cmd := Command{
				Cmd: "cancel",
			}
			cal.Commands = append(cal.Commands, cmd)
		}
	case "repeat":
		{
			resultRepeat, err := cal.repeat(value, arg, cal.Commands)
			if err != nil {
				fmt.Println(err.Error())
				return value
			} else {
				cmd := Command{
					Cmd: "repeat",
					Arg: arg,
				}
				cal.Commands = append(cal.Commands, cmd)
			}
			result = resultRepeat
		}
	}
	return result
}

func (cal *CalculatorInstance) repeat(result float64, input float64, commands []Command) (float64, error) {
	if len(commands) < int(input) || input == 0 {
		return result, fmt.Errorf("invalid command")
	} else {
		if input <= 1 {
			popCmd := commands[len(commands)-int(input)]
			result = cal.ProceedInput(result, popCmd.Cmd, popCmd.Arg)
		} else {
			popCmd := commands[len(commands)-int(input)]
			result = cal.ProceedInput(result, popCmd.Cmd, popCmd.Arg)
			resultRepeat, err := cal.repeat(result, input-1, commands[len(commands)-int(input):])
			if err != nil {
				return result, err
			}
			result = resultRepeat
		}
	}
	return result, nil
}

func examineInput(util util.Util, argCommands []string, noArgCommands []string, input string) (float64, error) {
	var arg float64
	if util.Contains(argCommands, input) {
		return util.ReadInputFloat()
	} else if util.Contains(noArgCommands, input) {
		return arg, nil
	} else {
		return arg, fmt.Errorf("invalid command")
	}
}
