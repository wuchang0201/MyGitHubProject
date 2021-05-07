package main

import "fmt"

func main() {
	// 将匿名函数保存到变量
	add := func() {
		fmt.Println("hrello")
	}
	add() // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
}
