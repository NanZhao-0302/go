package main
package main

import (
"time"
)

func main() {
	db := ConnectDb()
	g := &circle{
		conn: db,
	}
	personInfo := &PersonalInformation{
		Id:     5,
		Name:   "都怪这月色",
		Sex:    "男",
		Tall:   1.80,
		Weight: 60,
		Age:    22,
		Time:   time.Now(),
		IsShow: 1,
	}
	//先发布圈子
	err := g.Add(personInfo)
	if err != nil {
		return
	}
	//获取圈子列表
	g.Get()
	dpi := &PersonalInformation{
		Id:     4,
		IsShow: 0,
	}
	//删除圈子
	derr := g.Delete(dpi)
	if derr != nil {
		return
	}
	//再次获取圈子列表 看不到已经删除的状态数据
	g.Get()

}
