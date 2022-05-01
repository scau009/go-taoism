package externalcontact

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/context"
)

// Client 外部联系接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
