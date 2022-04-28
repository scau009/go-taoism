package uruntime

import (
	"path"
	"runtime"
)

func GetAbsPath() string {
	if _, file, _, ok := runtime.Caller(1); ok {
		return path.Dir(file)
	}
	panic("获取路径失败")
}
