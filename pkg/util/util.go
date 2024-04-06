package util

type Util interface {
	Contains(s []string, str string) bool
	ReadInputFloat() (float64, error)
	ReadInputString() (string, error)
}
