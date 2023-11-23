package router

import (
	"github.com/gin-gonic/gin"
)

// 用于初始化路由组
func InitRouter(r *gin.Engine) {
	// 初始化user路由组
	InitUser(r)
	// 初始化course路由组
	InitCourse(r)
}
