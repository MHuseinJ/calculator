package calc_engine

import (
	"calculator/pkg/math_engine"
	"calculator/pkg/util"
	"fmt"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestCalculator_ProceedInput(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockMathEngine := math_engine.NewMockMathEngine(ctrlMock)
	mockUtil := util.NewMockUtil(ctrlMock)
	type fields struct {
		argCommands   []string
		noArgCommands []string
		mathEngine    math_engine.MathEngine
		Commands      []Command
		lastValue     float64
	}
	type args struct {
		value   float64
		command string
		arg     float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		mock   func()
	}{
		{
			name: "testing with args input like add",
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				mathEngine:    mockMathEngine,
				Commands:      make([]Command, 0),
				lastValue:     float64(7),
			},
			args: args{
				value:   float64(7),
				command: "add",
				arg:     float64(8),
			},
			want: float64(15),
			mock: func() {
				mockMathEngine.EXPECT().Add(float64(7), float64(8)).Return(float64(15))
			},
		},
		{
			name: "testing with no args input like cancel",
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				mathEngine:    math_engine.NewMathEngineForCalculator(),
				Commands:      make([]Command, 0),
				lastValue:     float64(7),
			},
			args: args{
				value:   float64(7),
				command: "cancel",
			},
			want: float64(0),
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			cal := &CalculatorInstance{
				argCommands:   tt.fields.argCommands,
				noArgCommands: tt.fields.noArgCommands,
				mathEngine:    mockMathEngine,
				Commands:      tt.fields.Commands,
				lastValue:     tt.fields.lastValue,
				util:          mockUtil,
			}
			if got := cal.ProceedInput(tt.args.value, tt.args.command, tt.args.arg); got != tt.want {
				t.Errorf("ProceedInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_ReadInput(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockMathEngine := math_engine.NewMockMathEngine(ctrlMock)
	mockUtil := util.NewMockUtil(ctrlMock)

	type fields struct {
		argCommands   []string
		noArgCommands []string
		Commands      []Command
		lastValue     float64
	}
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		mock    func()
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "testing input with args",
			mock: func() {
				mockUtil.EXPECT().Contains([]string{"add", "subtract", "multiply", "divide", "repeat"}, "add").Return(true)
				mockUtil.EXPECT().ReadInputFloat().Return(float64(5), nil)
				mockMathEngine.EXPECT().Add(float64(8), float64(5)).Return(float64(5))
			},
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				Commands:      make([]Command, 0),
				lastValue:     float64(8),
			},
			args: args{
				input: "add",
			},
			wantErr: false,
		},
		{
			name: "testing input with false command",
			mock: func() {
				mockUtil.EXPECT().Contains([]string{"add", "subtract", "multiply", "divide", "repeat"}, "adds").Return(false)
				mockUtil.EXPECT().Contains([]string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"}, "adds").Return(false)
			},
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				Commands:      make([]Command, 0),
				lastValue:     float64(8),
			},
			args: args{
				input: "adds",
			},
			wantErr: true,
		},
		{
			name: "testing input with false arg",
			mock: func() {
				mockUtil.EXPECT().Contains([]string{"add", "subtract", "multiply", "divide", "repeat"}, "divide").Return(true)
				mockUtil.EXPECT().ReadInputFloat().Return(float64(0), fmt.Errorf("wrong arg"))
			},
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				Commands:      make([]Command, 0),
				lastValue:     float64(8),
			},
			args: args{
				input: "divide",
			},
			wantErr: true,
		},
		{
			name: "testing input with no arg",
			mock: func() {
				mockUtil.EXPECT().Contains([]string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"}, "sqr").Return(true)
				mockUtil.EXPECT().Contains([]string{"add", "subtract", "multiply", "divide", "repeat"}, "sqr").Return(false)
				mockMathEngine.EXPECT().Sqr(float64(8)).Return(float64(64))
			},
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				Commands:      make([]Command, 0),
				lastValue:     float64(8),
			},
			args: args{
				input: "sqr",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			cal := &CalculatorInstance{
				argCommands:   tt.fields.argCommands,
				noArgCommands: tt.fields.noArgCommands,
				mathEngine:    mockMathEngine,
				Commands:      tt.fields.Commands,
				lastValue:     tt.fields.lastValue,
				util:          mockUtil,
			}
			if err := cal.ReadInput(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("ReadInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculator_repeat(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockMathEngine := math_engine.NewMockMathEngine(ctrlMock)
	mockUtil := util.NewMockUtil(ctrlMock)
	type fields struct {
		argCommands   []string
		noArgCommands []string
		mathEngine    math_engine.MathEngine
		Commands      []Command
		lastValue     float64
		util          util.Util
	}
	type args struct {
		result   float64
		input    float64
		commands []Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
		mock    func()
	}{
		{
			name: "testing repeat with more than one commands",
			fields: fields{
				argCommands:   []string{"add", "subtract", "multiply", "divide", "repeat"},
				noArgCommands: []string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				mathEngine:    mockMathEngine,
				Commands:      []Command{Command{"add", float64(2)}, Command{"multiply", float64(2)}},
				util:          mockUtil,
				lastValue:     0.0,
			},
			args: args{
				result:   float64(0),
				input:    float64(2),
				commands: []Command{{"add", float64(2)}, {"multiply", float64(2)}},
			},
			want:    float64(4),
			wantErr: false,
			mock: func() {
				mockMathEngine.EXPECT().Add(float64(0), float64(2)).Return(float64(2))
				mockMathEngine.EXPECT().Multiply(float64(2), float64(2)).Return(float64(4))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			cal := &CalculatorInstance{
				argCommands:   tt.fields.argCommands,
				noArgCommands: tt.fields.noArgCommands,
				mathEngine:    tt.fields.mathEngine,
				Commands:      tt.fields.Commands,
				lastValue:     tt.fields.lastValue,
			}
			got, err := cal.repeat(tt.args.result, tt.args.input, tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("repeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repeat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCal(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockMathEngine := math_engine.NewMockMathEngine(ctrlMock)
	mockUtil := util.NewMockUtil(ctrlMock)

	type args struct {
		engine math_engine.MathEngine
		util   util.Util
	}
	tests := []struct {
		name string
		args args
		mock func()
		want CalculatorEngine
	}{
		{
			name: "testing constructor Math Engine for Calculator",
			args: args{
				engine: mockMathEngine,
				util:   mockUtil,
			},
			mock: func() {},
			want: &CalculatorInstance{
				[]string{"add", "subtract", "multiply", "divide", "repeat"},
				[]string{"cancel", "exit", "abs", "neg", "sqrt", "sqr", "cubert", "cube"},
				mockMathEngine,
				[]Command{},
				0.0,
				mockUtil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			if got := NewCal(tt.args.engine, tt.args.util); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCal() = %v, want %v", got, tt.want)
			}
		})
	}
}
