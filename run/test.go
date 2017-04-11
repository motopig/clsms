package main

import (
	"blue"
	"fmt"
)

func main() {
	blue.ACCOUNT = "xxx"
	blue.PASSWORD = "yyy"
	err, status := blue.Send("13888888888", "你好")
	if status > 0 {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success")
	}
}
