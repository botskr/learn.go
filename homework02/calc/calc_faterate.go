package calculator

import (
	gobmi "go-bmi"
)

func CalcFatRate(bmi float64, age int, sex string) (fatRate float64) {
	return gobmi.CalcFatRate(bmi, age, sex)
}
