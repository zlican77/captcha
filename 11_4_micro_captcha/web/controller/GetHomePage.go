package controller

import "github.com/gin-gonic/gin"

func GetHomePage(ctx *gin.Context) {
	ctx.Writer.WriteString("hello home")
}
