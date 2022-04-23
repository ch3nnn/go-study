/**
 * @Author: chentong
 * @Date: 2022/04/23 23:35
 */

package main

import (
	"fmt"
	"time"
)

// 启动当个goroutine
func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello()
	fmt.Println("main func done!")
	time.Sleep(time.Second)
}
