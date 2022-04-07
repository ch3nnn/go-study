/**
 * @Author: chentong
 * @Date: 2022/04/07 22:26
 */

package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 10
	/* 使用 if 语句判断布尔表达式 */
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)
}
