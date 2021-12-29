package main

import "testing"

func TestCase1(t *testing.T) {
	//测试案例就按需求一步一步往下写

	//• 王强第一次录入的时候，他的体脂是 38
	inputInfo("王强", 0.38)
	//• 王强第二次录入的时候，他的体脂是 32
	inputInfo("王强", 0.32)
	//• 这时，王强的最佳体脂是 32
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		//为了避免后面命名冲突，可以在花括号下给一个作用域
		if rankOfWQ != 1 {
			t.Fatalf("预期王强第一，但是得到的是%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但是得到的是%d", rankOfWQ)
		}
	}
	//• 李静录入他的体脂 28
	inputInfo("李静", 0.28)
	//• 李静的最佳体脂是 28
	//• 李静排名第一，体脂 28；王强排名第二，体脂 32。
	{
		rankOfWQ, fatRateOfWQ := getRank("王强")
		if rankOfWQ != 2 {
			t.Fatalf("预期王强第二，但是得到的是%d", rankOfWQ)
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂是0.32，但是得到的是%d", rankOfWQ)
		}
	}
	{
		rankOfLJ, fatRateOfLJ := getRank("李静")
		if rankOfLJ != 1 {
			t.Fatalf("预期李静第一，但是得到的是%d", rankOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂是0.28，但是得到的是%d", rankOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	//• 王强录入体脂是 38
	//• 张伟录入体脂是 38
	//• 李静录入体脂是 28
	//• 李静排名第一，体脂 28；王强、张伟排名第二，体脂 38。
	inputInfo("王强", 0.38)
	inputInfo("张伟", 0.38)
	inputInfo("李静", 0.28)

	{
		rankOFWQ, fatRateOFWQ := getRank("王强")
		if rankOFWQ != 2 {
			t.Fatalf("预期王强是第二，但得到的是%d", rankOFWQ)
		}
		if fatRateOFWQ != 0.38 {
			t.Fatalf("预期王强的体脂率是0.38，但得到的是%d", fatRateOFWQ)
		}
	}
	{
		rankOFZW, fatRateOFZW := getRank("张伟")
		if rankOFZW != 2 {
			t.Fatalf("预期张伟是第二，但得到的是%d", rankOFZW)
		}
		if fatRateOFZW != 0.38 {
			t.Fatalf("预期张伟的体脂率是0.38，但得到的是%d", fatRateOFZW)
		}
	}
	{
		rankOFLJ, fatRateOFLJ := getRank("李静")
		if rankOFLJ != 1 {
			t.Fatalf("预期李静是第一，但得到的是%d", rankOFLJ)
		}
		if fatRateOFLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是0.28，但得到的是%d", fatRateOFLJ)
		}
	}

}

func TestCase3(t *testing.T) {

}
