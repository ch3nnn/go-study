/**
 * @Author: chentong
 * @Date: 2022/04/19 10:56
 */

package main

import "fmt"

type User01 struct {
	id       int
	username string
}

type Manager struct {
	// 匿名字段
	User01
}

func (self *User01) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}

func main() {
	m := Manager{User01{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}
