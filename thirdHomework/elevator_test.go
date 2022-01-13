package main

import "testing"

// 测试案例 1
// 楼层有 5 层, 所有电梯楼层没有人请求电梯, 电梯不动
func TestCase1InitFloor(t *testing.T) {
	elevator := NewElevatorDefault("Elevator1", 5)
	t.Log("TestCase1 elevator:", elevator.floorNumber, "  currentFloorNumber:", elevator.currentFloorNumber)
	if elevator.floorNumber != 5 {
		t.Fail()
	}

	if elevator.currentFloorNumber != 1 {
		t.Fail()
	}
}

// 测试案例 2
// 楼层有 5 层, 电梯在 1 层. 三楼按电梯, 电梯向三楼行进, 并停在三楼
func TestCase2RunElevator(t *testing.T) {
	elevator := NewElevatorDefault("Elevator2", 5)
	elevator.RunTarget(3)
}

// 测试案例 3
// 楼层有 5 层, 电梯在 3 层, 上来一些人后, 目标楼层: 4 楼, 2 楼. 电梯向上到 4 楼, 然后转头到 2 楼, 最后停在 2 楼
func TestCase3RunElevator(t *testing.T) {
	elevator := NewElevatorCurrent("Elevator3", 5, 3)
	elevator.RunTarget(4)
	elevator.RunTarget(2)
}

// 测试案例 4
// 楼层有 5 层, 电梯在 3 层. 上来一些人后, 目标楼层: 4 楼, 5 楼, 2 楼. 电梯先向上到 4 楼, 然后到 5 楼, 之后转头到 2楼, 最后停在 2楼
func TestCase4RunElevator(t *testing.T) {
	elevator := NewElevatorCurrent("Elevator4", 5, 3)
	elevator.RunTarget(4)
	elevator.RunTarget(5)
	elevator.RunTarget(2)
}
