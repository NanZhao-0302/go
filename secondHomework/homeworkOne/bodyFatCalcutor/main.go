package main

import (
	"fmt"
	"learn.go/secondHomework/homeworkOne/calc"
)

//1.获取一些信息 getFromInput()
//2.计算出体脂率 fateRate = calculateFatRate()
//3.给出匹配信息和建议 suggestion =checkHealthAndGetSuggestion()

//因为有多个函数，所以先定义全局变量
var name string
var weight float64
var height float64
var age int
var sex string
var sexWeight int

var bmi float64
var fatRate float64
var totalFatRate float64
var cnn int = 0
var avgFatRate float64
var whether string

//写一个可无限工作，且有截止条件的体脂计算器
func main() {
	for {
		//录入信息
		weight, height, age, sex := getFromInput()
		//计算体脂率
		calculateFatRate(bmi, age, sexWeight, weight, height)
		//给出建议
		checkHealthAndGetSuggestion(sex, age, fatRate)
		//计算平均值
		avgFat(fatRate)
		//是否继续
		if forward := whetherContinue(); !forward {
			//if可以先跟一个块，得到一个产出，然后再对这个产出判断，如果false就break退出循环
			break
		}
	}
}

func getFromInput() (weight float64, height float64, age int, sex string) {
	//1.获取一些信息 getFromInput()
	fmt.Print("姓名:")
	fmt.Scanln(&name)

	fmt.Print("体重（Kg）:")
	fmt.Scanln(&weight)

	fmt.Print("身高（m）:")
	fmt.Scanln(&height)

	fmt.Print("年龄:")
	fmt.Scanln(&age)

	fmt.Print("性别（男/女）:")
	fmt.Scanln(&sex)
	return
}

func calculateFatRate(bmi float64, age int, sexWeight int, weight float64, height float64) (fatRate float64) {
	//2.1 计算出bmi bmi = weight / (height * height)
	bmi, _ = calcuator.CalcBmi(height, weight)
	//2.2 计算出体脂率 fatRate = (1.23*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fatRate = calcuator.CalcFateRate(bmi, age, sex)
	fmt.Println("体脂率是：", fatRate)
	return fatRate
}

func checkHealthAndGetSuggestion(sex string, age int, fatRate float64) {
	//3.给出匹配信息和建议 suggestion =checkHealthAndGetSuggestion()
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatRate >= 0 && fatRate < 0.1 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.1 && fatRate < 0.16 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.16 && fatRate < 0.21 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.21 && fatRate < 0.26 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate >= 0 && fatRate < 0.11 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.11 && fatRate < 0.17 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.17 && fatRate < 0.22 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.22 && fatRate < 0.27 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else if age >= 60 {
			if fatRate >= 0 && fatRate < 0.13 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.13 && fatRate < 0.19 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.19 && fatRate < 0.24 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.24 && fatRate < 0.29 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大了")
		}
	} else {
		if age >= 18 && age <= 39 {
			if fatRate >= 0 && fatRate < 0.2 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.2 && fatRate < 0.27 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.27 && fatRate < 0.34 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.34 && fatRate < 0.39 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else if age >= 40 && age <= 59 {
			if fatRate >= 0 && fatRate < 0.21 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.21 && fatRate < 0.28 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.28 && fatRate < 0.35 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.35 && fatRate < 0.4 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else if age >= 60 {
			if fatRate >= 0 && fatRate < 0.22 {
				fmt.Println("目前是：偏瘦，要多吃多练")
			} else if fatRate >= 0.22 && fatRate < 0.29 {
				fmt.Println("目前是：标准，好好保持哟")
			} else if fatRate >= 0.29 && fatRate < 0.36 {
				fmt.Println("目前是：偏重，该有些适当的运动啦")
			} else if fatRate >= 0.36 && fatRate < 0.41 {
				fmt.Println("目前是：肥胖，快点减肥")
			} else {
				fmt.Println("目前是：严重肥胖，再不减肥就要去医院了")
			}
		} else {
			fmt.Println("我们不看未成年人的体脂率，变化太大了")
		}
	}
}

func avgFat(fatRate float64) {
	//计算平均体脂率
	cnn++
	totalFatRate += fatRate
	avgFatRate = totalFatRate / float64(cnn)
	fmt.Println("共录入", cnn, "人，", "平均体脂率为", avgFatRate)
}

func whetherContinue() bool {
	fmt.Print("是否继续输入（y/n）:")
	fmt.Scanln(&whether)
	//写一个可以无限工作的计算器，需要有一个结束条件
	if whether != "y" {
		return false
	}
	return true
}
