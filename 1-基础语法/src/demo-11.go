package main

import "fmt"

// 指针数组

func test1(a [5]*int) {
	*a[1] = 2
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i])
	}
	fmt.Println()
}

func main() {

	// 声明指针数组
	var a [5]*int
	fmt.Println(a) // [<nil> <nil> <nil> <nil> <nil>]
	// 为指针每个元素赋指针
	for i := 0; i < 5; i++ {
		temp := i
		// &temp 指针
		a[i] = &temp
	}
	// 循环遍历取出指针对应的值
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i]) // 0 1 2 3 4
	}
	fmt.Println()
	test1(a) // 0 2 2 3 4
	for i := 0; i < 5; i++ {
		fmt.Print(" ", *a[i]) //  0 2 2 3 4
	}
}
