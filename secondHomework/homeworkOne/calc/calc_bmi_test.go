package calcuator

import (
	"testing"
)

func TestCalaBmi(t *testing.T) {
	inputHeigut, inputWeight := 1.0, 1.0
	expectedOutput := 1.0
	//在关键部位打了一个log
	//t.Logf("开始计算，输入：height:%f,weight:%f,期望结果：%f", inputHeigut, inputWeight, expectedOutput)
	actualOutput, err := CalcBmi(inputHeigut, inputWeight)
	//t.Logf("实际得到：%f,error:%v", actualOutput, err)
	if err != nil { //出错了，意味着测试失败了，error通常用%v
		t.Fatalf("expecting no err,but got :%v", err) //t.Fatalf严重错误，后面的测试没有意义了，直接结束
	}
	if expectedOutput != actualOutput {
		t.Errorf("expecting%f,but got %f", expectedOutput, actualOutput) //失败了，但是返回错误
	}
}
