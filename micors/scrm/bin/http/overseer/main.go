package main

import (
	"dev.taoism.gz.cn/go-taoism/micors/scrm/bin/http/app"
	"fmt"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"log"
	"net/http"
	"os"
)

var BuildID = "0"

func main() {

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
