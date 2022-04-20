/**
 * @Author: chentong
 * @Date: 2022/04/20 10:16
 */

package main

import "fmt"

func test01() {
	intArray := [5]int{1, 2, 3, 4, 5}
	intArray[1] = 123
	fmt.Println(intArray)
	index := 10
	intArray[index] = 10
	fmt.Println(intArray)

}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		// 自己抛
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func test02() {
	getCircleArea(-5)
}
func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func test04() {
	test03()
	fmt.Println("test04")
}

func main() {
	//test01()
	//test02()
	//test03()
	test04()
}
