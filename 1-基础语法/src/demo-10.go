package main

import "fmt"

// 数组指针

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println("遍历K, V ", i, v)
	}
}

func main() {

	var arr1 [5]int
	printArr(&arr1)
	fmt.Println(arr1)

	arr2 := [...]int{2, 4, 6, 8, 10}
	printArr(&arr2)
	fmt.Println(arr2)

}
