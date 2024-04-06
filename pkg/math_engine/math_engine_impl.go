package math_engine

import "math"

func (m *MathEngineForCalculator) Add(x float64, y float64) float64 {
	return x + y
}
func (m *MathEngineForCalculator) Multiply(x float64, y float64) float64 {
	return x * y
}
func (m *MathEngineForCalculator) Subtract(x float64, y float64) float64 {
	return x - y
}
func (m *MathEngineForCalculator) Divide(x float64, y float64) float64 {
	return x / y
}
func (m *MathEngineForCalculator) Abs(x float64) float64 {
	return math.Abs(x)
}
func (m *MathEngineForCalculator) Neg(x float64) float64 {
	return -x
}
func (m *MathEngineForCalculator) Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func (m *MathEngineForCalculator) Sqr(x float64) float64 {
	return x * x
}
func (m *MathEngineForCalculator) Cubert(x float64) float64 {
	return math.Cbrt(x)
}
func (m *MathEngineForCalculator) Cube(x float64) float64 {
	return x * x * x
}
