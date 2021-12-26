package main

import (
	"fmt"
)

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	var x4 float64
	var y4 float64
	var x3 float64
	var y3 float64
	fmt.Println("请输入直线一的横坐标x1：")
	fmt.Scanln(&x1)
	fmt.Println("请输入直线一的纵坐标y1：")
	fmt.Scanln(&y1)
	fmt.Println("请输入直线一的横坐标x2：")
	fmt.Scanln(&x2)
	fmt.Println("请输入直线一的纵坐标y2：")
	fmt.Scanln(&y2)
	fmt.Println("请输入直线二的横坐标x3：")
	fmt.Scanln(&x3)
	fmt.Println("请输入直线二的纵坐标y3：")
	fmt.Scanln(&y3)
	fmt.Println("请输入直线二的横坐标x4：")
	fmt.Scanln(&x4)
	fmt.Println("请输入直线二的纵坐标y4：")
	fmt.Scanln(&y4)
	if (x2-x1)*(y4-y3) == (y2-y1)*(x4-x3) {
		fmt.Println("两直线平行")
	} else {
		fmt.Println("两直线不平行")
	}
}
