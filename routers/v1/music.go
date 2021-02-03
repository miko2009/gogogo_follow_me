package v1

import (
	"github.com/gin-gonic/gin"
	material "github.com/miko2009/gogogo_follow_me/controller/material"
)

func RegisterMusicRouter(r *gin.Engine) {
	musicRouter := r.Group("/v1")
	{
		musicRouter.GET("/musics", material.List)
		musicRouter.GET("/musics/first", material.First)
	}
}
