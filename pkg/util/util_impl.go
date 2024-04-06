package util

import "fmt"

func (u *UtilInstance) Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (u *UtilInstance) ReadInputFloat() (float64, error) {
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		return float64(input), err
	} else {
		return float64(input), nil
	}
}

func (u *UtilInstance) ReadInputString() (string, error) {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		return input, err
	} else {
		return input, nil
	}
}
