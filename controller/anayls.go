package controller

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

func Analys(c *gin.Context) {
	examid := c.Query("examid")
	if examid == "" {
		c.Error(errors.New("examid 应该被传过来"))
	}

	resp := model.Analyse(examid)
	c.JSON(200, &resp)
}
