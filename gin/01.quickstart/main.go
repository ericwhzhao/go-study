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
