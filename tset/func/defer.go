package main

import (
	"fmt"
	"os"
	"time"
)

func openFile() {
	fileName := "/"
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
	//文件存在的话
	//panic之后运行
	defer fileObj.Close()
}
func deferGuess() {
	startTime := time.Now()
	defer fmt.Println("duration:", time.Now().Sub(startTime))
	time.Sleep(5 * time.Second)
	fmt.Println("start time: ", startTime)
}