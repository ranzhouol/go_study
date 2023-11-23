package router

import (
	"github.com/gin-gonic/gin"
	"ranzhouol/go_study/go_web/gin/web"
)

func InitUser(r *gin.Engine) {
	// 1. 创建v1路由组
	v1 := r.Group("/v1")
	// 再次分组
	user := v1.Group("/user")

	// 2. 构建路由
	user.GET("", web.GetUser)
	user.POST("", web.AddUser)
}
