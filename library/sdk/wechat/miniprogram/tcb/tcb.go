package tcb

import "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/miniprogram/context"

// Tcb Tencent Cloud Base
type Tcb struct {
	*context.Context
}

// NewTcb new Tencent Cloud Base
func NewTcb(context *context.Context) *Tcb {
	return &Tcb{
		context,
	}
}
