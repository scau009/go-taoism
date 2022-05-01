package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller 控制器结构
type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// JsonResponse 统一接口返回结构
type JsonResponse struct {
	// Success 本次请求是否成功
	Success bool `json:"success"`
	// Data 数据
	Data interface{} `json:"data,omitempty"`
	// ErrorCode 错误码
	ErrorCode *string `json:"errorCode,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// ShowType 错误展示类型： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
	ShowType *int `json:"showType,omitempty"`
	// TraceId 错误追踪 Id，一般用于排查日志
	TraceId *string `json:"traceId,omitempty"`
	// Host 响应服务器的地址，一般用于排查日志
	Host *string `json:"host,omitempty"`
}

type OptFunc func(opt *JsonResponse)

func WithData(data interface{}) OptFunc {
	return func(opt *JsonResponse) {
		opt.Data = data
	}
}

func WithErrorCode(errorCode string) OptFunc {
	return func(opt *JsonResponse) {
		opt.ErrorCode = &errorCode
	}
}

func WithErrorMessage(errorMessage string) OptFunc {
	return func(opt *JsonResponse) {
		opt.ErrorMessage = &errorMessage
	}
}

func WithShowType(showType int) OptFunc {
	return func(opt *JsonResponse) {
		opt.ShowType = &showType
	}
}

func (r *Controller) SendJsonResponse(ctx *gin.Context, optionFunctions ...OptFunc) {
	traceId := ctx.GetString("traceId")
	host := ctx.GetString("host")
	resp := JsonResponse{
		Success: true,
		TraceId: &traceId,
		Host:    &host,
	}
	for _, f := range optionFunctions {
		f(&resp)
	}

	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,HEAD,PATCH")
	ctx.Header("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Headers", ctx.Request.Header.Get("Access-Control-Request-Headers"))

	// 所有的接口都响应 200，但是实际业务逻辑按照结构内的定义
	ctx.JSON(http.StatusOK, resp)
}
