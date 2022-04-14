/**
 * @Author: chentong
 * @Date: 2022/04/12 09:56
 */

package main

import "fmt"

//
// defer01
// @Description: defer 先进后出
//
func defer01() {
	var whatever [5]struct{}
	for i := range whatever {
		// 先进后出
		defer fmt.Println(i)
	}
}

//
// defer02
// @Description: defer 碰上闭包
//
func defer02() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

func main() {

	defer01()
	fmt.Println("####################")
	defer02()

}
