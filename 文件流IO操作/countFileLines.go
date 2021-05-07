package main

import (
	"fmt"
	"io/ioutil"
)

func main(){

	countFileLine("user.txt")
}

func countFileLine(name string) {

	fmt.Println("count file:", name)
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return
	}

	count := 1
	for i := 1; i < len(data); i++ {
		if data[i] == '\n' {
			count++
		}
	}
	fmt.Println("line:", count)

}
