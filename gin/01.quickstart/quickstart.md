# Gin快速入门
## 安装
1、安装Gin包

```bash
$ go get -u github.com/gin-gonic/gin
```
2、导入Gin包

在`gin/quickstart/main.go`中，导入

```bash
$ import "github.com/gin-gonic/gin"
```
## 启动

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	//1、创建路由router
	r := gin.Default()
	//2、Http的GET请求
	//gin.Context，封装了request和response
	r.GET("/gin", func(c *gin.Context) {
		c.Writer.Write([]byte("hello, gin!"))
	})
	//3、监听端口，默认是8080
	r.Run(":8080")
}
```