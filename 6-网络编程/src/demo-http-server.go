package main

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr, "连接成功")
	// 请求方式：GET POST DELETE PUT UPDATE
	fmt.Println("method:", request.Method)
	// 访问/go
	fmt.Println("url:", request.URL.Path)
	fmt.Println("header:", request.Header)
	fmt.Println("body:", request.Body)
	// 回复
	writer.Write([]byte("hello http"))

}

func main() {
	http.HandleFunc("/go", Handler)
	// addr：监听的地址
	// handler：回调函数
	fmt.Println("http server running: 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
