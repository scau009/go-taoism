package controller

import "github.com/gin-gonic/gin"

func (r *Controller) GetCurrentUser(ctx *gin.Context) {
	r.SendJsonResponse(ctx, WithErrorCode("401"), WithErrorMessage("未登录"))

	//r.SendJsonResponse(ctx,WithData(gin.H{
	//	"name":   "admin",
	//	"avatar": "https://avatars.githubusercontent.com/u/18747266?s=40&v=4",
	//	"userid": "1",
	//}))
}
