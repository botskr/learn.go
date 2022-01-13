package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	lift := &Elevator{
		intMaxFloor:     5,
		intMinFloor:     1,
		intCurrentFloor: 2,
	}
	arrRequestFloor = []int{}
	strDirection := lift.getDirection(arrRequestFloor)
	if strDirection != "Don't move" {
		t.Fatalf("预期电梯不动，但是现在:%s", strDirection)
	}
}

func TestCase2(t *testing.T) {
	lift := &Elevator{
		intMaxFloor:     5,
		intMinFloor:     1,
		intCurrentFloor: 1,
		intSpeed:        1,
		intOpenDoorTime: 1,
	}
	arrRequestFloor = []int{3}
	strDirection := lift.getDirection(arrRequestFloor)
	if strDirection != "up" {
		t.Fatalf("预期电梯上升，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 3 {
		t.Fatalf("预期电梯停在三楼，但是现在：%d", lift.intCurrentFloor)
	}
}

func TestCase3(t *testing.T) {
	lift := &Elevator{
		intMaxFloor:     5,
		intMinFloor:     1,
		intCurrentFloor: 3,
		intSpeed:        1,
		intOpenDoorTime: 1,
	}
	arrRequestFloor = []int{4, 2}
	strDirection := lift.getDirection(arrRequestFloor)
	if strDirection != "up" {
		t.Fatalf("预期电梯上升，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 4 {
		t.Fatalf("预期电梯停在四楼，但是现在：%d", lift.intCurrentFloor)
	}
	arrRequestFloor = arrRequestFloor[1:]
	strDirection = lift.getDirection(arrRequestFloor)
	if strDirection != "down" {
		t.Fatalf("预期电梯下降，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 2 {
		t.Fatalf("预期电梯停在二楼，但是现在：%d", lift.intCurrentFloor)
	}
}

func TestCase4(t *testing.T) {
	lift := &Elevator{
		intMaxFloor:     5,
		intMinFloor:     1,
		intCurrentFloor: 3,
		intSpeed:        1,
		intOpenDoorTime: 1,
	}
	arrRequestFloor = []int{4}
	strDirection := lift.getDirection(arrRequestFloor)
	if strDirection != "up" {
		t.Fatalf("预期电梯上升，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 4 {
		t.Fatalf("预期电梯停在四楼，但是现在：%d", lift.intCurrentFloor)
	}
	arrRequestFloor = []int{5}
	strDirection = lift.getDirection(arrRequestFloor)
	if strDirection != "up" {
		t.Fatalf("预期电梯上升，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 5 {
		t.Fatalf("预期电梯停在五楼，但是现在：%d", lift.intCurrentFloor)
	}
	arrRequestFloor = []int{2}
	strDirection = lift.getDirection(arrRequestFloor)
	if strDirection != "down" {
		t.Fatalf("预期电梯下降，但是现在:%s", strDirection)
	}
	lift.move(lift.strDirection, lift.intCurrentFloor)
	if lift.intCurrentFloor != 2 {
		t.Fatalf("预期电梯停在二楼，但是现在：%d", lift.intCurrentFloor)
	}
}
