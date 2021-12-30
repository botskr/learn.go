package gobmi

func CalcFateRate(bmi float64, age int, sexVal int) (bfr float64) {
	bfr = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*(float64(sexVal))) / 100
	return
}
