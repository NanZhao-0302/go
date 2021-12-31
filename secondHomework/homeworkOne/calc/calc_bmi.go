package calcuator

import (
	gobmi "github.com/armstrongli/go-bmi"
)

func CalcBmi(height float64, weight float64) (bmi float64, err error) {
	//直接调用github上的代码
	bmi, err = gobmi.BMI(weight, height)
	return
}
