/**
 * @Author: chentong
 * @Date: 2022/04/20 22:54
 */

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	// 关闭连接
	defer conn.Close()

	// 标准输出和标准错误文件描述符打开文件
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		// 如果输出Q退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		read, err := conn.Read(buf[:])
		if err != nil {
			return
		}
		fmt.Println("接收数据: ", string(buf[:read]))

	}

}
