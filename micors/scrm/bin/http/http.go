package main

import (
	"dev.taoism.gz.cn/go-taoism/micors/scrm/bin/http/app"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config/connection"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config/logger"
	"flag"
)

func init() {
	tomlFile := flag.String("c", "micors/scrm/config.toml", "配置文件路径")
	config.SetToml(*tomlFile)
	connection.InitConnection()
	logger.InitLogger("http")
}

func main() {
	router := app.GetRouter()
	_ = router.Run(":8080") //8032
}
