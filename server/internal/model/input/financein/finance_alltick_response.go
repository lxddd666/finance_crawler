// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceAlltickResponseUpdateFields 修改alltick返回值字段过滤
type FinanceAlltickResponseUpdateFields struct {
	Msg        string      `json:"msg"        dc:"msg"`
	Trace      string      `json:"trace"      dc:"trace"`
	Data       *gjson.Json `json:"data"       dc:"data"`
	Code       string      `json:"code"       dc:"code"`
	Sort       int         `json:"sort"       dc:"排序"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     float64     `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
}

// FinanceAlltickResponseInsertFields 新增alltick返回值字段过滤
type FinanceAlltickResponseInsertFields struct {
	Msg        string      `json:"msg"        dc:"msg"`
	Trace      string      `json:"trace"      dc:"trace"`
	Data       *gjson.Json `json:"data"       dc:"data"`
	Code       string      `json:"code"       dc:"code"`
	Sort       int         `json:"sort"       dc:"排序"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     float64     `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
}

// FinanceAlltickResponseEditInp 修改/新增alltick返回值
type FinanceAlltickResponseEditInp struct {
	entity.FinanceAlltickResponse
}

func (in *FinanceAlltickResponseEditInp) Filter(ctx context.Context) (err error) {
	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceAlltickResponseEditModel struct{}

// FinanceAlltickResponseDeleteInp 删除alltick返回值
type FinanceAlltickResponseDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceAlltickResponseDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceAlltickResponseDeleteModel struct{}

// FinanceAlltickResponseViewInp 获取指定alltick返回值信息
type FinanceAlltickResponseViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceAlltickResponseViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceAlltickResponseViewModel struct {
	entity.FinanceAlltickResponse
}

// FinanceAlltickResponseListInp 获取alltick返回值列表
type FinanceAlltickResponseListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceAlltickResponseListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceAlltickResponseListModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Msg        string      `json:"msg"        dc:"msg"`
	Trace      string      `json:"trace"      dc:"trace"`
	Code       string      `json:"code"       dc:"code"`
	Sort       int         `json:"sort"       dc:"排序"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     float64     `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
}

// FinanceAlltickResponseExportModel 导出alltick返回值
type FinanceAlltickResponseExportModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Msg        string      `json:"msg"        dc:"msg"`
	Trace      string      `json:"trace"      dc:"trace"`
	Code       string      `json:"code"       dc:"code"`
	Sort       int         `json:"sort"       dc:"排序"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     float64     `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
}

// FinanceAlltickResponseMaxSortInp 获取alltick返回值最大排序
type FinanceAlltickResponseMaxSortInp struct{}

func (in *FinanceAlltickResponseMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceAlltickResponseMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}
