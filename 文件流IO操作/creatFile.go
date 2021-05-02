package main

import (
	"fmt"
	"os"
)

func main(){

	path :="writeuser.txt"
	file,err :=os.Create(path)
	if err!=nil{
		fmt.Println(err)
	}else{
		str := "hello 沙河"
		file.Write([]byte(str)) //写入字节切片数据
		file.WriteString("hello 小王子") //直接写入字符串数据
		file.Close()

	}


}


