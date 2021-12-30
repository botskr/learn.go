package main

import "fmt"

func GetUserBmi(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("录入的体重小于0kg,请核对信息后重新录入")
		return
	}
	if heightM <= 0 {
		err = fmt.Errorf("输入的身高小于等于0m,请核对信息后重新录入")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

func GetUserFateRate(bmi float64, age int, sex string) (fateRate float64, suggest string, err error) {
	fateRate = 0.0
	suggest = ""
	if bmi <= 0 {
		err = fmt.Errorf("bmi录入不能为负数")
		return
	}

	if age <= 0 || age > 150 {
		err = fmt.Errorf("年龄不能为负数或者大于150")
		return
	}
	mapSex := map[string]int{
		"男": 1,
		"女": 0,
	}
	sexVal, ok := mapSex[sex]
	if ok == false {
		err = fmt.Errorf("性别输入异常，只能男/女")
	}
	fateRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*(float64(sexVal))) / 100
	return
}
