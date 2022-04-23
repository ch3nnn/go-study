/**
 * @Author: chentong
 * @Date: 2022/04/23 23:40
 */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello01(i int) {
	fmt.Println("hello goroutine", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go hello01(i)
	}
	wg.Wait()

}
