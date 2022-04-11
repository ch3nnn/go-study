/**
 * @Author: chentong
 * @Date: 2022/04/11 14:37
 */

package main

import (
	"fmt"
)

func add(a, b int) (c int) {
	c = a + b
	return c
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	// “裸”返回 sum, avg
	return
}

func main() {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
	// 输出结果：
	// 1 2 3 3 1
}
