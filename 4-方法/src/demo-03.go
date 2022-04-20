/**
 * @Author: chentong
 * @Date: 2022/04/19 11:03
 */

package main

import "fmt"

type User03 struct {
	id       int
	username string
}

type Manager03 struct {
	User03
	title string
}

func (self *User03) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager03) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	manager03 := Manager03{User03{1, "小毛驴"}, "Administrator"}
	fmt.Println(manager03.ToString())
	fmt.Println(manager03.User03.ToString())

}
