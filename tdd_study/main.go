package main

import (
	"math"
	"sort"
)

var personFatrate = map[string]float64{}

func inputInfo(name string, fatRate ...float64) {
	//不定长参数传入的是一个数组或者是一个切片，就选择最佳体脂传入
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item { //只要比当前小就选一个更小的值
			minFatRate = item
		}
	}
	//先用键值对进行实例化，把姓名和值录入

	personFatrate[name] = minFatRate
}

func getRank(name string) (rank int, fatRate float64) {
	//我们需要根据体脂进行排名，需要先把里面的键值对反转一下
	fatRatePerson := map[float64][]string{}
	rankArr := make([]float64, 0, len(personFatrate))
	for nameItem, fatRateItem := range personFatrate {
		//原来的键变成值
		fatRatePerson[fatRateItem] = append(fatRatePerson[fatRateItem], nameItem)
		rankArr = append(rankArr, fatRateItem) //只对体脂率进行排名，所以数组中只引入fatRateItem
	}
	//需要一个排序函数，然后把每次排序的结果扔到一个数组里
	//sort函数直接把数组顺序排好
	sort.Float64s(rankArr)
	//数组的下标是0开头，所以rank的值需要下标+1
	for i, fatRateItem := range rankArr {
		//得到了体脂率对应的name数组
		_names := fatRatePerson[fatRateItem]
		for _, _name := range _names {
			if _name == name {
				rank = i + 1
				fatRate = fatRateItem
				return
			}
		}
	}
	return 0, 0
}
