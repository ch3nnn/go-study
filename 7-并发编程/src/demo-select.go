/**
 * @Author: chentong
 * @Date: 2022/04/24 22:39
 */

package main

import (
	"fmt"
	"time"
)

/*
select 多路复用

语法格式:
select {
	case <-chan1:
   // 如果chan1成功读到数据，则进行该case处理语句
	case chan2 <- 1:
   // 如果成功向chan2写入数据，则进行该case处理语句
	default:
   // 如果上面都没有成功，则进入default处理流程
}


*/

func testSelect01(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test"
}

func testSelect02(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"

}

func main() {
	// 创建两个字符串通道
	strings01 := make(chan string)
	strings02 := make(chan string)

	go testSelect01(strings01)
	go testSelect02(strings02)

	// select可以同时监听一个或多个channel
	// 直到其中一个channel ready
	select {
	case s1 := <-strings01:
		fmt.Println("s1:", s1)
	case s2 := <-strings02:
		fmt.Println("s2:", s2)
	}

}
