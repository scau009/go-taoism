package controller

import (
	"dev.taoism.gz.cn/go-taoism/micors/scrm/service/wework"
	"github.com/gin-gonic/gin"
)

func (r *Controller) GetLoginQrCode(ctx *gin.Context) {

	r.SendJsonResponse(ctx, WithData(wework.NewWeWorkService().GetLoginQrCode()))

}
