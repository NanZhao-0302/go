package main

import (
	"fmt"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	//BMI案例
	//录入正常身高、体重，确保计算结果符合预期；
	bmiTest1, errCase1 := CalculateBMI(1.58, 51)
	fmt.Println("bmiTest1:", bmiTest1, " errCase1:", errCase1)
	if bmiTest1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	//录入 0 或负数身高，返回错误；
	bmiTest2, errCase2 := CalculateBMI(0, 51)
	fmt.Println("bmiTest2:", bmiTest2, " errCase2: ", errCase2)
	if bmiTest2 != 0 || errCase2 == nil {
		t.Fatal()
	}

	bmiTest3, errCase3 := CalculateBMI(-5, 51)
	fmt.Println("bmiTest3:", bmiTest3, " errCase3: ", errCase3)
	if bmiTest3 < 0 || errCase3 == nil {
		t.Fatal()
	}

	//录入 0 或负数体重，返回错误。
	bmiTest4, errCase4 := CalculateBMI(1.58, 0)
	fmt.Println("bmiTest4:", bmiTest4, " errCase4:", errCase4)
	if bmiTest4 != 0 || errCase4 == nil {
		t.Fatal()
	}

	bmiTest5, errCase5 := CalculateBMI(1.58, -80)
	fmt.Println("bmiTest5:", bmiTest5, " errCase5: ", errCase5)
	if bmiTest5 < 0 || errCase5 == nil {
		t.Fatal()
	}
}

func TestCalculateFatRate(t *testing.T) {
	//体脂率案例
	//录入正常 BMI、年龄、性别，确保计算结果符合预期；

	//录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。
	fatRateCase1, suggest1, errCase1 := CalculateFatRate(21.3, 26, "女")
	fmt.Println("fatRateCase1:", fatRateCase1, " suggest1:", suggest1, " errCase1:", errCase1)
	if fatRateCase1 < 0 || errCase1 != nil {
		t.Fatal()
	}

	//录入非法 BMI，返回错误；
	_, _, errCase2 := CalculateFatRate(0, 26, "女")
	if errCase2 == nil {
		t.Fatal()
	}

	_, _, errCase3 := CalculateFatRate(-5, 26, "女")
	if errCase3 == nil {
		t.Fatal()
	}

	//录入非法年龄，返回错误；
	_, _, errCase4 := CalculateFatRate(23.5, 0, "女")
	if errCase4 == nil {
		t.Fatal()
	}

	_, _, errCase5 := CalculateFatRate(29, 151, "男")
	if errCase5 == nil {
		t.Fatal()
	}

	//录入非法性别，返回错误；
	_, _, errCase6 := CalculateFatRate(26, 26, "你猜")
	if errCase6 == nil {
		t.Fatal()
	}

}
