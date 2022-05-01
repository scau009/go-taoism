package datacube

import (
	"dev.taoism.gz.cn/go-taoism/library/sdk/wechat/officialaccount/context"
)

type reqDate struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// DataCube 数据统计
type DataCube struct {
	*context.Context
}

// NewCube 数据统计
func NewCube(context *context.Context) *DataCube {
	dataCube := new(DataCube)
	dataCube.Context = context
	return dataCube
}
