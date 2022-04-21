/**
 * @Author: chentong
 * @Date: 2022/04/21 10:25
 */

package main

import (
	"fmt"
	"net"
)

//
// main
// @Description: 网络编程 upd 服务端
//
func main() {

	// upd地址
	updAddr := net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	}
	fmt.Printf("UDP server running...  IP: %v Port: %v  \n", updAddr.IP, updAddr.Port)

	listen, err := net.ListenUDP("udp", &updAddr)
	if err != nil {
		fmt.Println("listen failed err: ", err)
		return
	}
	// 关闭连接
	defer func(listen *net.UDPConn) {
		err := listen.Close()
		if err != nil {
			fmt.Println("listen failed err: ", err)
		}
	}(listen)
	for {
		// 接收数据
		var data [1024]byte
		udp, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("listen failed err: ", err)
			continue
		}
		fmt.Printf("data: %v add: %v count: %v \n\n", string(data[:udp]), addr, udp)
		// 发送数据
		_, err = listen.WriteToUDP(data[:udp], addr)
		if err != nil {
			fmt.Println("listen failed err: ", err)
			continue
		}
	}

}
