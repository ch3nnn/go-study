/**
 * @Author: chentong
 * @Date: 2022/04/11 23:03
 */

package main

import "fmt"

//
// test
// @Description: 定义返回两个函数类型的返回值
// @param base
// @return func(int) int
// @return func(int) int
//
func test(base int) (func(int) int, func(int) int) {
	/* 相加 */
	add := func(i int) int {
		base += 1
		return base
	}

	/* 相减 */
	sub := func(i int) int {
		base -= 1
		return base
	}

	return add, sub

}

func main() {
	add, sub := test(10)

	fmt.Println(add(1))

	fmt.Println(sub(2))

}
