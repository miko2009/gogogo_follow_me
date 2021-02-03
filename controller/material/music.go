package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/miko2009/gogogo_follow_me/controller"
	MaterialService "github.com/miko2009/gogogo_follow_me/service/material"

)

func List(ctx *gin.Context) {

	materialId := 250200078
	msg :=  "success"
	data := MaterialService.Show(materialId)
	controller.Render(ctx, msg, data)
	return
}

func First(ctx *gin.Context) {
	msg :=  "success"
	data := MaterialService.First()
	controller.Render(ctx, msg, data)
	return
}