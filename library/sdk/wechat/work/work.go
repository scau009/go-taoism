package work

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/credential"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/config"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/context"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/externalcontact"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/kf"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/msgaudit"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/oauth"
)

// Work 企业微信
type Work struct {
	ctx *context.Context
}

// NewWork init work
func NewWork(cfg *config.Config) *Work {
	defaultAkHandle := credential.NewWorkAccessToken(cfg.CorpID, cfg.CorpSecret, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Work{ctx: ctx}
}

// GetContext get Context
func (wk *Work) GetContext() *context.Context {
	return wk.ctx
}

// GetOauth get oauth
func (wk *Work) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wk.ctx)
}

// GetMsgAudit get msgAudit
func (wk *Work) GetMsgAudit() (*msgaudit.Client, error) {
	return msgaudit.NewClient(wk.ctx.Config)
}

// GetKF get kf
func (wk *Work) GetKF() (*kf.Client, error) {
	return kf.NewClient(wk.ctx.Config)
}

// GetExternalContact get external_contact
func (wk *Work) GetExternalContact() *externalcontact.Client {
	return externalcontact.NewClient(wk.ctx)
}
