package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) {
	// 获取请求URL中的query参数（ 在?之后的键值对 ）
	id := ctx.DefaultQuery("id", "nil") //获取到query字段，不存在就用默认字段

	// ctx.JSON返回JSON数据，包括状态码和gin.H的JSON数据
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get user",
		"id":      id,
	})
}

// 传输的数据可能是form类型，也可能是JSON类型，都需要tag
// binding:"required" 必传字段
// binding:"e164"     字段类型必须为手机号码格式
// binding:"number" 字段类型必须为数字
type User struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Phone string `json:"phone" form:"phone" binding:"required,e164"`
	Age   int32  `json:"age" form:"age" binding:"number"`
}

// 接受POST方法传过来的请求体数据
func AddUser(ctx *gin.Context) {
	req := &User{}
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
