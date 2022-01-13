package main

import (
	"fmt"
	"time"
)

var arrRequestFloor []int

type Elevator struct {
	intMaxFloor     int
	intMinFloor     int
	intCurrentFloor int
	intSpeed        int
	intOpenDoorTime int
	strDirection    string
}

func (lift *Elevator) getDirection(arrRequestFloor []int) string {
	if len(arrRequestFloor) == 0 || arrRequestFloor[0] == 0 || arrRequestFloor[0] == lift.intCurrentFloor {
		lift.strDirection = "Don't move"
		return lift.strDirection
	}
	if arrRequestFloor[0] > lift.intCurrentFloor {
		lift.strDirection = "up"
		return lift.strDirection
	}
	if arrRequestFloor[0] < lift.intCurrentFloor {
		lift.strDirection = "down"
		return lift.strDirection
	}
	return ""
}

func (lift *Elevator) move(strDirection string, intCurrentFloor int) {
	if strDirection == "up" {
		fmt.Println("上升")
		lift.up(intCurrentFloor)
		return
	}
	if strDirection == "down" {
		fmt.Println("下降")
		lift.down(intCurrentFloor)
		return
	}
	fmt.Println("没动")
}

func (lift *Elevator) up(intCurrentFloor int) {
	time.Sleep(time.Duration(lift.intSpeed))
	fmt.Println("电梯开始上升")
	for _, requestFloor := range arrRequestFloor {
		for i := lift.intCurrentFloor; i < requestFloor; i++ {
			intCurrentFloor++
			fmt.Printf("电梯上升到了%d层\n", intCurrentFloor)
			if intCurrentFloor == requestFloor {
				fmt.Printf("电梯开门...\n")
				time.Sleep(time.Duration(lift.intOpenDoorTime))
				fmt.Printf("电梯关门\n")
				lift.intCurrentFloor = intCurrentFloor
				return
			}
			if intCurrentFloor == lift.intMaxFloor {
				return
			}
		}
	}
}

func (lift *Elevator) down(intCurrentFloor int) {
	time.Sleep(time.Duration(lift.intSpeed))
	fmt.Println("电梯开始下降")
	for _, requestFloor := range arrRequestFloor {
		for i := lift.intCurrentFloor; i > requestFloor; i-- {
			intCurrentFloor--
			fmt.Printf("电梯下降到了%d层\n", intCurrentFloor)
			if intCurrentFloor == requestFloor {
				fmt.Printf("电梯开门...\n")
				time.Sleep(time.Duration(lift.intOpenDoorTime))
				fmt.Printf("电梯关门")
				lift.intCurrentFloor = intCurrentFloor
				return
			}
			if intCurrentFloor == lift.intMaxFloor {
				return
			}
		}

	}
}

func main() {
	lift := &Elevator{
		intMaxFloor:     5,
		intMinFloor:     1,
		intCurrentFloor: 3,
		intSpeed:        1,
		intOpenDoorTime: 1,
	}
	arrRequestFloor = []int{4, 2}
	strDirection := lift.getDirection(arrRequestFloor)
	fmt.Println(strDirection)
	lift.move(lift.strDirection, lift.intCurrentFloor)
	arrRequestFloor = arrRequestFloor[:1]
	fmt.Println(arrRequestFloor)
	strDirection = lift.getDirection(arrRequestFloor)
	lift.move(lift.strDirection, lift.intCurrentFloor)
}