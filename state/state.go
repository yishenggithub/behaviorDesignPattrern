package main

// 面向接口编程
// 先有一个表示状态的接口，其他具体的状态实现接口

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}
