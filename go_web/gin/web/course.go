package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCourse(ctx *gin.Context) {
	// 获取请求URL中的query参数（ 在?之后的键值对 ）
	id := ctx.DefaultQuery("id", "nil") //获取到query字段，不存在就用默认字段

	// ctx.JSON返回JSON数据，包括状态码和gin.H的JSON数据
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get Course",
		"id":      id,
	})
}

func GetCourse2(ctx *gin.Context) {
	// 获取请求URL中的路径参数
	id := ctx.Param("id")

	// ctx.JSON返回JSON数据，包括状态码和gin.H的JSON数据
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get Course",
		"id":      id,
	})
}

func GetCourseMath(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get Course math",
	})
}
func GetCourseChinese(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get Course Chinese",
	})
}

// 传输的数据可能是form类型，也可能是JSON类型，都需要tag
// binding:"required" 必传字段，binding:"number" 字段类型必须为数字
type Course struct {
	Name    string  `json:"name" form:"name" binding:"required"`
	Teacher string  `json:"teacher" form:"teacher" binding:"required"`
	Price   float64 `json:"price" form:"price" binding:"number"`
}

// 接受POST方法传过来的请求体数据
func AddCourse(ctx *gin.Context) {
	req := &Course{}
	// 参数绑定 .ShouldBind()
	err := ctx.ShouldBind(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 直接返回原数据
	ctx.JSON(http.StatusOK, req)
}
