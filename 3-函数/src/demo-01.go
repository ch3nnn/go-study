/**
 * @Author: chentong
 * @Date: 2022/04/11 14:13
 */

package main

import (
	"fmt"
)

/* 定义相互交换值的函数 */
func swap(x *int, y *int) {
	var temp int

	temp = *x /* 保存 x 的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y*/

}

func main() {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Println(a, b)

}
