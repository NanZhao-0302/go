package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	fmt.Println("第七周 第三次作业 模拟 电梯")

	elevator1 := NewElevatorDefault("demo1", 5)
	fmt.Println(currentTime(), "电梯层数:", elevator1.floorNumber)

	fmt.Println("第七周 第三次作业 模拟 电梯 完成")

}

type Elevator struct {
	id string // 电梯标识 方便调试

	floorNumber int // 电梯楼层数 一般仅初始化一次, 此案例层数 从1开始, 不支持负数 就是没有地下层数

	belowFlowNumber int // 负数楼层 如果有负数楼层, 单独设置

	currentFloorNumber int // 当前所在层数

	speed int // 运行速度 毫秒

	direction string // 运行方向 up-上 / down-下

	nextTargetFlow int // 下个目标方向

	maxFloorNumber int // 最大目标楼层

	minFloorNumber int // 最低目标楼层

	floorMap map[int]bool // 每层电梯的 是否需要开门 如果运行中有人按了楼层 在此保存

}

func currentTime() string {
	current := time.Now().Format("2006-01-02 15:04:05.999")
	return current
}

func (elevator *Elevator) currentFloorInfo() string {
	currentFloor := currentTime() + "  " + elevator.id + "  当前楼层: " + strconv.Itoa(elevator.currentFloorNumber)
	return currentFloor
}

func Run(elevator Elevator) {
	fmt.Println("电梯运行")
}

// RunTarget
// 无法解决 依次去往多个目标
//
func (elevator *Elevator) RunTarget(target int) {
	fmt.Println(elevator.currentFloorInfo(), "    电梯运行去往", target, " 楼层")
	if target >= 1 && target <= elevator.floorNumber {
		if target == elevator.currentFloorNumber {
			fmt.Println(elevator.currentFloorInfo(), " 目标楼层: ", target, " 与当前楼层相等, 不运行")
		} else if target > elevator.currentFloorNumber {
			for start := elevator.currentFloorNumber; start < target; start++ {
				time.Sleep(time.Duration(elevator.speed) * time.Millisecond)
				elevator.currentFloorNumber = start + 1
				fmt.Println(elevator.currentFloorInfo())
				// elevator.floorOpen()
			}
		} else if target < elevator.currentFloorNumber {
			for start := elevator.currentFloorNumber; start > target; start-- {
				time.Sleep(time.Duration(elevator.speed) * time.Millisecond)
				elevator.currentFloorNumber = start - 1
				fmt.Println(elevator.currentFloorInfo())
				// elevator.floorOpen()
			}

		}

		if elevator.currentFloorNumber == target {
			elevator.floorOpen()
		}

	} else {
		fmt.Println(elevator.currentFloorInfo(), " 目标楼层: ", target, " 无效")
	}
}

func NewElevatorDefault(id string, floorNumber int) *Elevator {
	//elevator := new(Elevator)
	//elevator.id = id
	//elevator.floorNumber = floorNumber
	// 设置默认值
	//elevator.currentFloorNumber = 1 // 默认楼层 1
	elevator := NewElevatorCurrent(id, floorNumber, 1)
	return elevator
}

func NewElevatorCurrent(id string, floorNumber int, currentFloorNumber int) *Elevator {
	elevator := new(Elevator)
	elevator.id = id
	elevator.floorNumber = floorNumber
	elevator.currentFloorNumber = currentFloorNumber
	elevator.speed = 1000
	fmt.Println("创建电梯成功 ", elevator.currentFloorInfo(), " 总楼层数:", elevator.floorNumber)
	return elevator
}

func up(elevator Elevator) {
	fmt.Println(elevator.id, " 上行 ")
}

func down(elevator Elevator) {
	fmt.Println(elevator.id, " 下行 ")
}

func (elevator Elevator) floorOpen() {
	elevator.pause()
	elevator.open()
	elevator.close()
}

func (elevator *Elevator) pause() {
	fmt.Println(elevator.id, " 暂停 ")
}

func (elevator Elevator) open() {
	fmt.Println(elevator.id, " 开门 ")
}

func (elevator Elevator) close() {
	fmt.Println(elevator.id, " 关门 ")
}
