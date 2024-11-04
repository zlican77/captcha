package main

import (
	"11_4_micro_captcha/web/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("view/*")
	r.GET("/home", controller.GetHomePage)
	r.GET("/login", controller.GetLoginPage)
	r.GET("/captcha/:uuid", controller.GetCaptcha)

	r.Run(":8080")
}
