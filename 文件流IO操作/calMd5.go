package main

import (
	"crypto/md5"
	"fmt"
)

func main(){

	fmt.Println(Md5Str("123456"))

}

func Md5Str(str string) string {

	return fmt.Sprintf("%x",md5.Sum([]byte(str)))
}
