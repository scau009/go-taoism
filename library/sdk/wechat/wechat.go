package wechat

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/cache"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/miniprogram"
	miniConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/miniprogram/config"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/officialaccount"
	offConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/officialaccount/config"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/openplatform"
	openConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/openplatform/config"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/pay"
	payConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/pay/config"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work"
	workConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/config"
)

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

// SetCache 设置cache
func (wc *Wechat) SetCache(cahce cache.Cache) {
	wc.cache = cahce
}

// GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *openConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
}

// GetWork 获取企业微信的实例
func (wc *Wechat) GetWork(cfg *workConfig.Config) *work.Work {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return work.NewWork(cfg)
}
