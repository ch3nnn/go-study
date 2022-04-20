/**
 * @Author: chentong
 * @Date: 2022/04/20 11:39
 */

package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student06 struct{}

func (stu *Student06) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	// 无法将 'Student06{}' (类型 Student06) 用作类型 People 类型未实现 'People'，因为 'Speak' 方法有指针接收器
	//var peo People = Student06{}
	var peo People = &Student06{}
	think := "sb"
	fmt.Println(peo.Speak(think))
}
