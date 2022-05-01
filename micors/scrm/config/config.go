package config

import (
	"dev.taoism.gz.cn/go-taoism/util/uruntime"
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

type LogConfig struct {
	Level      string `json:"level"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

type Config struct {
	Log LogConfig

	WeWork struct {
		CropId      string
		AgentId     string
		CallBackURL string
	}

	Mongodb struct {
		Uri      string
		PoolSize uint64
	}

	Dir struct {
		VarDir    string
		SecretDir string
		RootDir   string
		AssetDir  string
	}
	Host struct {
		AppHost string
	}

	App struct {
		Env string
	}
}

var (
	tomlFile string
	config   *Config
	once     sync.Once
)

func SetToml(filePath string) {
	tomlFile = filePath
}

func GetConfig() *Config {
	once.Do(ReloadConfig)
	return config
}

func ReloadConfig() {
	filePath, err := filepath.Abs(tomlFile)
	if err != nil {
		panic(err)
	}
	newConfig := new(Config)

	if _, err := toml.DecodeFile(filePath, newConfig); err != nil {
		panic(err)
	}

	config = newConfig

	currentDir := uruntime.GetAbsPath()
	config.Dir.VarDir = fmt.Sprintf("%s/%s", currentDir, "../var")
	config.Dir.SecretDir = fmt.Sprintf("%s/%s/%s", currentDir, "../secret", config.App.Env)
	config.Dir.RootDir = fmt.Sprintf("%s/%s", currentDir, "../")
	config.Dir.AssetDir = fmt.Sprintf("%s/%s", currentDir, "../assets")
}
