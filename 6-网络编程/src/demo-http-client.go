package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	get, err := http.Get("http://localhost:8080/go")
	if err != nil {
		fmt.Println(err)
	}
	defer get.Body.Close()
	// 状态码
	fmt.Println(get.Status)
	fmt.Println(get.Header)

	var buf [1024]byte
	for {
		n, err := get.Body.Read(buf[:])
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			res := string(buf[:n])
			fmt.Println(res)
			break
		}

	}
}
