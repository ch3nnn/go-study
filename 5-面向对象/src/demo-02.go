/**
 * @Author: chentong
 * @Date: 2022/04/20 11:10
 */

package main

import "fmt"

//人
type Person02 struct {
	name string
	sex  string
	age  int
}

type Student02 struct {
	Person02
	id   int
	addr string
	//同名字段
	name string
}

func main() {
	var s Student02
	// 给自己字段赋值了
	s.name = "5lmh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下
	s.Person02.name = "枯藤"
	fmt.Println(s)
}
