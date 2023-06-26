package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	ff := cors.Default()
	r.Use(ff)
	r.GET("/analyse", controller.Analys)
}
