package context

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/credential"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
