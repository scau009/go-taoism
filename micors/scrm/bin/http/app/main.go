package app

import (
	"dev.taoism.gz.cn/go-taoism/library/middleware/access_logger"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetRouter() *gin.Engine {
	r := gin.New()
	//配置上传文件大小
	r.MaxMultipartMemory = 50
	//配置路由信息
	r.Use(access_logger.GinLogger(), access_logger.GinRecovery(true))
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"app":       "scrm api",
			"version":   "1.0.0",
			"timestamp": time.Now(),
		})
	})

	ctl := controller.NewController()
	r.GET("/api/login", ctl.GetLoginQrCode)
	return r
}
