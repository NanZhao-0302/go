package main

import (
	"fmt"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	//BMI 计算：
	//录入正常身高、体重，确保计算结果符合预期；
	//录入 0 或负数身高，返回错误；
	//录入 0 或负数体重，返回错误。
	bmiTest1, errCase1 := CalculateBMI(1.58, 51)
	fmt.Println("bmiTest1:", bmiTest1, " errCase1:", errCase1)
	if bmiTest1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	bmiTest2, errCase2 := CalculateBMI(0, 60)
	fmt.Println("bmiTest2:", bmiTest2, " errCase2: ", errCase2)
	if bmiTest2 != 0 || errCase2 == nil {
		t.Fatal()
	}

	bmiTest3, errCase3 := CalculateBMI(1.71, 0)
	fmt.Println("bmiTest3:", bmiTest3, " errCase3:", errCase3)
	if bmiTest3 != 0 || errCase3 == nil {
		t.Fatal()
	}
}

func TestCalculateFatRate(t *testing.T) {
	//体脂率计算：
	//录入正常 BMI、年龄、性别，确保计算结果符合预期；
	//录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
	//录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。
	fatRateCase1, suggest1, errCase1 := CalculateFatRate(21.3, 29, "女")
	fmt.Println("fatRateCase1:", fatRateCase1, " suggest1:", suggest1, " errCase1:", errCase1)
	if fatRateCase1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	_, _, errCase2 := CalculateFatRate(0, 29, "g")
	if errCase2 == nil {
		t.Fatal()
	}

	// 年龄测试
	_, _, errCase3 := CalculateFatRate(23.5, 0, "boy")
	if errCase3 == nil {
		t.Fatal()
	}

	_, _, errCase4 := CalculateFatRate(29, 151, "b")
	if errCase4 == nil {
		t.Fatal()
	}

	// 性别测试
	_, _, errCase5 := CalculateFatRate(29, 33, "demo")
	if errCase5 == nil {
		t.Fatal()
	}

}
