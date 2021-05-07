package main

import (
	"fmt"
	"os"
)

func main(){

	path :="writeuser.txt"
	file,err :=os.OpenFile(path,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)


	if err!=nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "hello 沙河"
	file.Write([]byte(str)) //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据

}
func srcFile()string{


}



