package router

import (
	"github.com/gin-gonic/gin"
	"ranzhouol/go_study/go_web/gin/web"
)

func InitCourse(r *gin.Engine) {
	// 1. 创建v1路由组
	v1 := r.Group("/v1")
	// 再次分组
	course := v1.Group("/course")

	// 2. 构建路由
	course.GET("", web.GetCourse)      //获取query参数
	course.GET("/:id", web.GetCourse2) //获取路径参数
	course.POST("", web.AddCourse)     //参数绑定，用于接收POST传过来的数据

	course.GET("/math", web.GetCourseMath)
	course.GET("/chinese", web.GetCourseChinese)
}
