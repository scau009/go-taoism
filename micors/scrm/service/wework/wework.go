package wework

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/cache"
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work"
	workConfig "dev.taoism.gz.cn/go-taoism/library/sdk/wechat/work/config"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config"
	"sync"
)

type weworkService struct {
	wcw *work.Work
}

var (
	once sync.Once
	wws  *weworkService
)

func NewWeWorkService() *weworkService {
	once.Do(func() {
		wc := wechat.NewWechat()
		wc.SetCache(cache.NewMemory())
		wcw := wc.GetWork(&workConfig.Config{
			CorpID:  config.GetConfig().WeWork.CropId,
			AgentID: config.GetConfig().WeWork.AgentId,
		})
		wws = &weworkService{
			wcw: wcw,
		}
	})
	return wws
}
