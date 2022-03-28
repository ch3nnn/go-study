/**
 * @Author: chentong
 * @Date: 2022/03/28 16:36
 */

package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 c 的结果 %d \n", c)

	c = a - b
	fmt.Printf("第二行 c 的结果 %d \n", c)

	c = a * b
	fmt.Printf("第三行 c 的结果 %d \n", c)

	c = a / b
	fmt.Printf("第四行 c 的结果 %d \n", c)

	c = a % b
	fmt.Printf("第五行 c 的结果 %d \n", c)

	a++
	fmt.Printf("第六行 a 的结果 %d \n", a)

	a--
	fmt.Printf("第七行 a 的结果 %d \n", a)

}
