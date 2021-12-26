package main

import (
	"fmt"
)

// 课后作业 1：计算多个人的平均体脂

// 实现完整的体脂计算器
// 连续输入三人的姓名、性别、身高、体重、年龄信息
// 计算每个人的 BMI、体脂率
// 输出：
// 每个人姓名、BMI、体脂率、建议
// 总人数、平均体脂率

func main() {
	for i := 0; ; i++ {
		var name string
		fmt.Print("姓名：")
		fmt.Scanln(&name)
		var sex string
		fmt.Print("性别（男/女）：")
		fmt.Scanln(&sex)
		var tall float64
		fmt.Print("身高（米）：")
		fmt.Scanln(&tall)
		var weight float64
		fmt.Print("体重（公斤）：")
		fmt.Scanln(&weight)
		var age int
		fmt.Print("年龄：")
		fmt.Scanln(&age)

		var sexWeight int
		if sex == "男" {
			sexWeight = 1
		} else {
			sexWeight = 0
		}

		BMI := weight / (tall * tall)
		fatRate := (1.2*BMI + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100

		switch sexWeight {
		case 1:
			if age >= 18 && age <= 39 {
				if fatRate <= 0.1 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.16 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.21 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.26 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.11 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.17 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.22 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.27 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else if age >= 60 {
				if fatRate <= 0.13 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.19 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.24 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.29 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else {
				fmt.Println("你的年龄太小了，我都不知道说什么了。。")
			}
		case 0:
			if age >= 18 && age <= 39 {
				if fatRate <= 0.2 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.27 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.34 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.39 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else if age >= 40 && age <= 59 {
				if fatRate <= 0.21 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.28 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.35 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.4 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else if age >= 60 {
				if fatRate <= 0.22 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏瘦：太瘦啦！多吃点哦。")
				} else if fatRate <= 0.29 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "标准：很不错，请继续加油，维持住鸭。")
				} else if fatRate <= 0.36 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "偏重：胖了一点点哦，注意减重。")
				} else if fatRate <= 0.41 {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "肥胖：很胖哦，该运动运动了，注意饮食。")
				} else {
					fmt.Println(name, "你的BMI是", BMI, "体脂率是", fatRate, "严重肥胖：你的体脂率太高了！应予以高度重视。")
				}
			} else {
				fmt.Println("你的年龄太小了，我都不知道说什么了。。")
			}
		default:
			fmt.Println("您的输入不符合规范，请重新输入……")
		}
		var tFatRate float64
		tFatRate = tFatRate + fatRate
		aFatRate := tFatRate / float64(i+1)

		fmt.Println("总人数", i+1, "平均体脂率", aFatRate)
		var whethercontinue string
		fmt.Print("是否继续输入？（y/n）")
		fmt.Scanln(&whethercontinue)
		if whethercontinue != "y" {
			break
		}
	}

}
