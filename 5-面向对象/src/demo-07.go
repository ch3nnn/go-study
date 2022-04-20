/**
 * @Author: chentong
 * @Date: 2022/04/20 11:59
 */

package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println(c.name, "喵喵喵")
}

func (c cat) move() {
	fmt.Println(c.name, "猫会动")
}

func main() {
	var x animal
	x = cat{"小明"}
	x.move()
	x.say()

	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

}
