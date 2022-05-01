package config

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/cache"
)

// Config .config for 微信开放平台
type Config struct {
	AppID          string `json:"app_id"`           // appid
	AppSecret      string `json:"app_secret"`       // appsecret
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // EncodingAESKey
	Cache          cache.Cache
}
