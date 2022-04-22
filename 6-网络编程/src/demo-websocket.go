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

//websocket实现
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(":12345")
}
