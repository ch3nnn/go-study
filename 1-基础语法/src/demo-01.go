/**
 * @Author: chenTong
 * @Date: 2022/03/13 23:09
 */

package main

import "fmt"

func main() {

	var s = "string"
	const X int8 = 100
	a1, a2 := 1, "abc"

	fmt.Println(s, X, a1, a2)

	// 格式化字符串 %d 表示整型数字，%s 表示字符串
	var url = "Code=%d&endDate=%s"
	sprintf := fmt.Sprintf(url, 123, "2020-12-31")
	fmt.Println(sprintf)

	// 以下几种类型为nil
	// var a *int
	// var a []int
	// var a map[string] int
	// var a chan int
	// var func(string) int
	// var a error

}

// 打印输出
// string 100 1 abc
