package main

import (
	"dev.taoism.gz.cn/go-taoism/micors/scrm/bin/http/app"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config/connection"
	"dev.taoism.gz.cn/go-taoism/micors/scrm/config/logger"
	"flag"
	"fmt"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"log"
	"net/http"
	"os"
)

var BuildID = "0"

func main() {
	tomlFile := flag.String("c", "../../../config.toml", "配置文件路径")
	config.SetToml(*tomlFile)
	connection.InitConnection()
	logger.InitLogger("http")
	overseer.Run(overseer.Config{
		Program: overseerProgram,
		Address: ":5001",
		Debug:   false,
		Fetcher: &fetcher.File{Path: "main_latest"},
	})
}

func overseerProgram(state overseer.State) {

	fmt.Printf("构建 ID %s(%s), 进程(%d), 正在监听 %s \n", BuildID, state.ID, os.Getpid(), "http://0.0.0.0:5001")

	router := app.GetRouter()
	err := http.Serve(state.Listener, router)
	if err != nil {
		log.Printf("服务器启动错误: %s\n", err.Error())
	}
}
