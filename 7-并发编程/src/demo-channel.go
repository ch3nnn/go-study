/**
 * @Author: chentong
 * @Date: 2022/04/24 20:30
 */

package main

import (
	"fmt"
)

//  无缓冲通道
func makeChan() {
	intChan := make(chan int)
	// 无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。
	intChan <- 10
	fmt.Println("send success")
}

func recv(c chan int) {
	var ret int
	ret = <-c
	fmt.Println(ret)
}

func makeSendChan() {
	// 创建int类型通道
	insChan := make(chan int)
	go recv(insChan)
	// 发送 20
	insChan <- 20
	fmt.Println("send success")

}

//
// makeChanSize
// @Description: 创建有缓冲的通道
//
func makeChanSize() {
	ins := make(chan int, 10)
	ins <- 10
	fmt.Println("send success")

}

func exampleChan() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 开启goroutine将0~10的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			if i, ok := <-ch1; ok {
				ch2 <- i * i
				continue
			}
			close(ch2)
			break
		}
	}()
	// 通道关闭后会退出for range循环
	for i := range ch2 {
		fmt.Println(i)
	}
}

func main() {
	//makeChan()
	//makeSendChan()
	//makeChanSize()
	exampleChan()
}
