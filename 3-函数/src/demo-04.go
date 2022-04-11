/**
 * @Author: chentong
 * @Date: 2022/04/11 22:48
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getSqrt(4))

}
