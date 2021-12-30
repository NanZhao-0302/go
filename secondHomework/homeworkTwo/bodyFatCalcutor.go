package main

import "fmt"

func CalculateBMI(height float64, weight float64) (float64, error) {
	if height <= 0 {
		err := fmt.Errorf("身高不能为0或者负数")
		return 0, err
	}

	if weight <= 0 {
		err := fmt.Errorf("体重不能为0或者负数")
		return 0, err

	}

	bmi := weight / (height * height)
	return bmi, nil
}

func CalculateFatRate(bmi float64, age int, gender string) (float64, string, error) {

	var suggest string = ""
	var genderType = 0
	// 校验 bmi
	if bmi <= 0 {
		error := fmt.Errorf("bmi录入不能为负数")
		return 0, suggest, error
	}

	// 校验年龄
	if age <= 0 || age > 150 {
		error := fmt.Errorf("年龄不能为负数或者大于150")
		return 0, suggest, error
	}

	// 校验性别
	if "b" == gender || "boy" == gender {
		genderType = 1
	} else if "g" == gender || "girl" == gender {
		genderType = 0
	} else {
		error := fmt.Errorf("性别输入无效 男: b boy  女: g girl")
		return 0, suggest, error
	}

	fatRate := 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(genderType)
	return fatRate, suggest, nil
}
