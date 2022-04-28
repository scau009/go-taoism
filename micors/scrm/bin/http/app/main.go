package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	//配置上传文件大小
	r.MaxMultipartMemory = 50
	//配置路由信息

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"app":       "scrm api",
			"version":   "1.0.0",
			"timestamp": time.Now(),
		})
	})

	return r
}
