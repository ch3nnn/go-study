/**
 * @Author: chentong
 * @Date: 2022/04/20 11:50
 */

package main

import "fmt"

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var washingMachine WashingMachine
	washingMachine = new(haier)
	washingMachine.dry()
	washingMachine.wash()
}
