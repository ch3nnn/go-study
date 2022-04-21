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
// @Description: 网络编程 upd 客户端
//
func main() {

	udpAddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	udpConn, err := net.DialUDP("udp", nil, &udpAddr)
	if err != nil {
		fmt.Printf("连接服务端失败: %s \n", err)
		return
	}
	defer udpConn.Close()
	sendData := []byte("hello udp server")
	_, err = udpConn.Write(sendData)
	if err != nil {
		fmt.Printf("发送数据失败: %s \n", err)
		return
	}
	readData := [1024]byte{}
	udp, addr, err := udpConn.ReadFromUDP(readData[:])
	if err != nil {
		fmt.Printf("接收数据失败: %s \n", err)
		return
	}
	fmt.Printf("recv: %v addr: %v count: %v \n", string(readData[:udp]), addr, udp)
}
