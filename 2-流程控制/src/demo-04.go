/**
 * @Author: chentong
 * @Date: 2022/04/10 04:05
 */

package main

func main() {
	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。

		println(&s)

		if i == 0 {
			s = s[:]   // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。

		}

		println(i, v)
	}
}
