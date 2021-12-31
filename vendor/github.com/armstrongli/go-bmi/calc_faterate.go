package gobmi

//扩展的文件
import "fmt"

func CalcFateRate(bmi float64, age int, sex string) (fatRate float64) {
	if bmi < 0 {
		panic("bmi不能为0或者负数")
	}
	if age < 0 {
		panic("年龄不能为0或者负数")
	}
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}

	if sexWeight < 0 {
		panic("性别权重不能为0或者负数")
	}

	fatRate = (1.23*bmi + getAgeWeight(age)*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率是：", fatRate)
	return
}

func getAgeWeight(age int) (ageWeight float64) {
	ageWeight = 0.23
	if age >= 30 && age <= 40 {
		ageWeight = 0.22
	}
	return

}
