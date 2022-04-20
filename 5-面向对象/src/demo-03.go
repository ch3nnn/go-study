/**
 * @Author: chentong
 * @Date: 2022/04/20 11:16
 */

package main

import "fmt"

//人
type Person03 struct {
	name string
	sex  string
	age  int
}

// 自定义类型
type mystr string

// 学生
type Student03 struct {
	Person03
	int
	mystr
}

func main() {
	s1 := Student03{Person03{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
}
