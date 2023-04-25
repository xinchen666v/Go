package main

import "fmt" //导入IO库

// 程序的入口点
func main() {
	s1 := "Hello" + "World" //定义字符串
	fmt.Println("Hello, world!！")
	say("Hello Go!")             //调用函数
	fmt.Println(len(s1))         //打印字符串的长度
	fmt.Println(string(s1[0:5])) //打印字符串的前五个字符
}
func say(message string) {
	fmt.Println("You said:", message)
}
