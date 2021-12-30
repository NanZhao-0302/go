package calcuator

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBmi(height float64, weight float64) (bmi float64, err error) {
	bmi, err = gobmi.BMI(weight, height)
	return
}
