package main

import "fmt"

// 程序的入口点
func main() {
	fmt.Println("Hello, world!！")
	say("Hello Go!") //调用函数
}
func say(message string) {
	fmt.Println("You said:", message)
}
