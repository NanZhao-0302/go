package main

import "sort"

var personFatrate = map[string]float64{}

func inputInfo(name string, fatRate float64) {
	//先用键值对进行实例化，把姓名和值录入
	personFatrate[name] = fatRate
}

func getRank(name string) (rank int, fatRate float64) {
	//我们需要根据体脂进行排名，需要先把里面的键值对反转一下
	fatRatePerson := map[float64]string{}
	rankArr := make([]float64, 0, len(personFatrate))
	for nameItem, fatRateItem := range personFatrate {
		//原来的键变成值
		fatRatePerson[fatRateItem] = nameItem
		rankArr = append(rankArr, fatRateItem) //只对体脂率进行排名，所以数组中只引入fatRateItem
	}
	//需要一个排序函数，然后把每次排序的结果扔到一个数组里
	//sort函数直接把数组顺序排好
	sort.Float64s(rankArr)
	//数组的下标是0开头，所以rank的值需要下标+1
	for i, fatRateItem := range rankArr {
		_name := fatRatePerson[fatRateItem]
		if _name == name {
			rank = i + 1
			fatRate = fatRateItem
			return
		}
	}
	return 0, 0
}
