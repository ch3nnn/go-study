/**
 * @Author: chentong
 * @Date: 2022/04/22 10:02
 */

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

/*
webSocket是什么

WebSocket是一种在单个TCP连接上进行全双工通信的协议
WebSocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
在WebSocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输
需要安装第三方包：
cmd中：go get -u -v github.com/gorilla/websocket

参考资料: https://zhuanlan.zhihu.com/p/329991604

需要使用的框架：

1、gin框架：gin-gonic/gin

2、websocket框架：gorilla/websocket

3、api文档：
	https://godoc.org/github.com/gin-gonic/gin
	https://godoc.org/github.com/go

*/

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//
// index
// @Description: 渲染 websocket 页面
// @param c
//
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "websocket.html", nil)

}

////
// ping
// @Description: websocket 实现
// @param c
//
func ping(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// 写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	r := gin.Default()
	// LoadHTMLGlob()方法可以加载模板文件
	r.LoadHTMLGlob("6-网络编程/src/*.html")
	// 路由
	r.GET("/", index)
	r.GET("/ws", ping)

	r.Run("localhost:8000")
}
