package main

import (
	"log"
	"os"
)
//追加文件
func main(){
	path :="user.log"
	file,err :=os.OpenFile(path,os.O_APPEND|os.O_CREATE,os.ModePerm)

	if err != nil{
		log.Fatal(err)
	}

	//输出到相应的文件夹，而不是前台输出
	log.SetOutput(file)
	log.Print("第一条日志！！！")
}
