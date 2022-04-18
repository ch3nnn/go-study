/**
 * @Author: chentong
 * @Date: 2022/04/18 23:43
 */

package main

import "fmt"

// User User结构体
type User struct {
	Name  string
	Email string
}

//
// Notify
// @Description:
// @receiver user
//
func (user User) Notify() {
	fmt.Printf("%v : %v \n", user.Name, user.Email)
}

func main() {
	// 值类型调用方法
	user := User{"golang", "golang@test.com"}
	user.Notify()

	// 指针类型调用方法
	user2 := User{"python", "python@test.com"}
	user3 := &user2
	user3.Notify()
}
