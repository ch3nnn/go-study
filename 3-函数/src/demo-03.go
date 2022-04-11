/**
 * @Author: chentong
 * @Date: 2022/04/11 15:03
 */

package main

func add03(x, y int) (z int) {
	defer func() {
		z += 100
	}()

	z = x + y
	return
}

func add04(x, y int) (z int) {
	defer func() {
		println(z + 1) // 输出: 204
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
}

func main() {
	println(add03(1, 2))
	println(add04(1, 2))
}
