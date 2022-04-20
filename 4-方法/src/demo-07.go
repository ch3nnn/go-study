/**
 * @Author: chentong
 * @Date: 2022/04/20 10:29
 */

package main

import (
	"fmt"
	"os"
	"time"
)

//
// PathError
// @Description: 路径错误
//
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

//
// Error
// @Description: 错误方法
// @receiver p PathError
// @return string
//
func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s",
		p.path, p.op, p.createTime, p.message)
}

//
// Open
// @Description: 打开文件
// @param filename 文件路径
// @return error
//
func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}

}
