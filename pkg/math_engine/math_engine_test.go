package math_engine

import (
	"math"
	"reflect"
	"testing"
)

func TestMathEngineForCalculator_Abs(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing abs",
			args: args{
				float64(-2),
			},
			want: float64(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Add(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing add",
			args: args{
				float64(2),
				float64(5),
			},
			want: float64(7),
		},
		{
			name: "testing add with negative value",
			args: args{
				float64(-2),
				float64(-5),
			},
			want: float64(-7),
		},
		{
			name: "testing add with negative and positive value",
			args: args{
				float64(2),
				float64(-5),
			},
			want: float64(-3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Cube(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing cube with positive value",
			args: args{
				float64(2),
			},
			want: float64(8),
		},
		{
			name: "testing cube with negative value",
			args: args{
				float64(-2),
			},
			want: float64(-8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Cube(tt.args.x); got != tt.want {
				t.Errorf("Cube() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Cubert(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing cube with negative value",
			args: args{
				float64(-8),
			},
			want: float64(-2),
		},
		{
			name: "testing cube with positive value",
			args: args{
				float64(27),
			},
			want: float64(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Cubert(tt.args.x); got != tt.want {
				t.Errorf("Cubert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Divide(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing positif divide with 0",
			args: args{
				float64(7),
				float64(0),
			},
			want: math.Inf(1),
		},
		{
			name: "testing negative divide with 0",
			args: args{
				float64(-7),
				float64(0),
			},
			want: math.Inf(-1),
		},
		{
			name: "testing divide with positive number",
			args: args{
				float64(8),
				float64(2),
			},
			want: float64(4),
		},
		{
			name: "testing divide with negative number",
			args: args{
				float64(8),
				float64(-2),
			},
			want: float64(-4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Divide(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Multiply(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing multiply with 0",
			args: args{
				float64(7),
				float64(0),
			},
			want: float64(0),
		},
		{
			name: "testing multiply with positive number",
			args: args{
				float64(8),
				float64(2),
			},
			want: float64(16),
		},
		{
			name: "testing multiply with negative number",
			args: args{
				float64(8),
				float64(-2),
			},
			want: float64(-16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Multiply(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Neg(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing neg with negative number",
			args: args{
				float64(-8),
			},
			want: float64(8),
		},
		{
			name: "testing neg with negative number",
			args: args{
				float64(8),
			},
			want: float64(-8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Neg(tt.args.x); got != tt.want {
				t.Errorf("Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Sqr(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing sqr with positive value",
			args: args{
				float64(4),
			},
			want: float64(16),
		},
		{
			name: "testing sqr with negative value",
			args: args{
				float64(-4),
			},
			want: float64(16),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Sqr(tt.args.x); got != tt.want {
				t.Errorf("Sqr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMathEngineForCalculator_Sqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing sqrt with positive value",
			args: args{
				float64(16),
			},
			want: float64(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Sqrt(tt.args.x); got != tt.want {
				if !math.IsNaN(got) {
					t.Errorf("Sqrt() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestMathEngineForCalculator_Sqrt_NaN(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing sqrt with negative value",
			args: args{
				float64(-16),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Sqrt(tt.args.x); !math.IsNaN(got) {
				t.Errorf("Sqrt() = %v, it suppostube NaN", got)
			}
		})
	}
}

func TestMathEngineForCalculator_Subtract(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "testing subtract with negative value",
			args: args{
				float64(-16),
				float64(-16),
			},
			want: float64(0),
		},
		{
			name: "testing subtract with positive value",
			args: args{
				float64(16),
				float64(15),
			},
			want: float64(1),
		},
		{
			name: "testing subtract with positive negative value",
			args: args{
				float64(16),
				float64(-15),
			},
			want: float64(31),
		},
		{
			name: "testing subtract with negative positive value",
			args: args{
				float64(-16),
				float64(15),
			},
			want: float64(-31),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MathEngineForCalculator{}
			if got := m.Subtract(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMathEngine(t *testing.T) {
	tests := []struct {
		name string
		want MathEngine
	}{
		{
			"Testing constructor Math Engine for Calculator",
			&MathEngineForCalculator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMathEngineForCalculator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMathEngineForCalculator() = %v, want %v", got, tt.want)
			}
		})
	}
}
