package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义中间件
var token = "123"

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("access_token")
		fmt.Println("access_token", accessToken)
		if accessToken != token {
			ctx.JSON(http.StatusForbidden, gin.H{
				"message": "token 校验失败",
			})
			ctx.Abort() //阻止调用后面的处理函数
		}

		ctx.Next() //调用后面的处理函数
	}
}
