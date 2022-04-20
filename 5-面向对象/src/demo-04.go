/**
 * @Author: chentong
 * @Date: 2022/04/20 11:18
 */

package main

import "fmt"

//人
type Person05 struct {
	name string
	sex  string
	age  int
}

// 学生
type Student05 struct {
	*Person05
	id   int
	addr string
}

func main() {
	s1 := Student05{&Person05{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person05.name)
}
