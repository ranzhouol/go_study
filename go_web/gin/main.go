package main

import (
	"github.com/gin-gonic/gin"
	"ranzhouol/go_study/go_web/gin/middleware"
	"ranzhouol/go_study/go_web/gin/router"
)

func main() {
	r := gin.Default()
	// 全局中间件
	r.Use(middleware.Auth())

	// 简单方法
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.PUT("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put pong",
		})
	})
	r.DELETE("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete pong",
		})
	})
	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post pong",
		})
	})

	// 路由组
	router.InitRouter(r)

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}
