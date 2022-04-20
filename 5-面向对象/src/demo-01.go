/**
 * @Author: chentong
 * @Date: 2022/04/20 11:01
 */

package main

import "fmt"

type Person struct {
	name string
	sex  string
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func main() {
	var student = Student{Person{"小名", "男", 18}, 1, "杭州"}
	fmt.Println(student)
}
