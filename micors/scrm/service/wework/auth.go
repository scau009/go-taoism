package wework

import "dev.taoism.gz.cn/go-taoism/micors/scrm/config"

func (w *weworkService) GetLoginQrCode() string {
	return w.wcw.GetOauth().GetQrContentTargetURL(config.GetConfig().WeWork.CallBackURL)
}
