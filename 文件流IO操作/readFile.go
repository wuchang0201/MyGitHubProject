package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	file, err := os.Open("user.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 10)
	//循环读取文件，文件读完输出
	for{
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}

		fmt.Println(string(tmp[:n]))
	}

}
