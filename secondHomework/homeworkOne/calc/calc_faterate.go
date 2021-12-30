package calcuator

import gobmi "github.com/armstrongli/go-bmi"

func CalcFateRate(bmi float64, age int, sex string) (fatRate float64) {
	gobmi.CalcFateRate(bmi, age, sex)
	return
}
