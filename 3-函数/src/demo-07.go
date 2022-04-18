/**
 * @Author: chentong
 * @Date: 2022/04/15 10:19
 */

package main

func main() {
	test07()
}

func test07() {

	// 匿名函数
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
}
