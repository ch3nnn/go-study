/**
 * @Author: chentong
 * @Date: 2022/04/19 11:56
 */

package main

import "fmt"

type User04 struct {
	id   int
	name string
}

func (self User04) Test() {
	fmt.Println(self)
}

func main() {
	u := User04{1, "Tom"}

	// 声明
	//var mValue = u.Test
	mValue := u.Test // 立即值复制 receiver，因为不是指针类型，不受后续修改影响。

	fmt.Println(u.Test)
	fmt.Println(&mValue)

	u.id, u.name = 2, "Jack"
	u.Test()

	mValue()
}
