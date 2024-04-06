package math_engine

type MathEngine interface {
	Add(x float64, y float64) float64
	Multiply(x float64, y float64) float64
	Subtract(x float64, y float64) float64
	Divide(x float64, y float64) float64
	Abs(x float64) float64
	Neg(x float64) float64
	Sqrt(x float64) float64
	Sqr(x float64) float64
	Cubert(x float64) float64
	Cube(x float64) float64
}
