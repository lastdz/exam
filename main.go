package main

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.Init()
	r := gin.Default()

	InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
