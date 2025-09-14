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

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceKlineUpdateFields 修改k线字段过滤
type FinanceKlineUpdateFields struct {
	Code       string  `json:"code"       dc:"code"`
	KlineType  int     `json:"klineType"  dc:"k线类型"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64 `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64 `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"该K线最低价"`
	Volume     int64   `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64 `json:"turnover"   dc:"该K线成交金额"`
	Key        string  `json:"key"        dc:"唯一key，timestamp和code组合"`
	Scale      int     `json:"scale"      dc:"多少分钟一根K线"`
	Day        string  `json:"day"        dc:"日期"`
}

// FinanceKlineInsertFields 新增k线字段过滤
type FinanceKlineInsertFields struct {
	Code       string  `json:"code"       dc:"code"`
	KlineType  int     `json:"klineType"  dc:"k线类型"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64 `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64 `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"该K线最低价"`
	Volume     int64   `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64 `json:"turnover"   dc:"该K线成交金额"`
	Key        string  `json:"key"        dc:"唯一key，timestamp和code组合"`
	Scale      int     `json:"scale"      dc:"多少分钟一根K线"`
	Day        string  `json:"day"        dc:"日期"`
}

// FinanceKlineEditInp 修改/新增k线
type FinanceKlineEditInp struct {
	entity.FinanceKline
}

func (in *FinanceKlineEditInp) Filter(ctx context.Context) (err error) {
	// 验证唯一key，timestamp和code组合
	if err := g.Validator().Rules("required").Data(in.Key).Messages("唯一key，timestamp和code组合不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证多少分钟一根K线
	if err := g.Validator().Rules("required").Data(in.Scale).Messages("多少分钟一根K线不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceKlineEditModel struct{}

// FinanceKlineDeleteInp 删除k线
type FinanceKlineDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceKlineDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKlineDeleteModel struct{}

// FinanceKlineViewInp 获取指定k线信息
type FinanceKlineViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceKlineViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKlineViewModel struct {
	entity.FinanceKline
}

// FinanceKlineListInp 获取k线列表
type FinanceKlineListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceKlineListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKlineListModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     int64       `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"唯一key，timestamp和code组合"`
	Scale      int         `json:"scale"      dc:"多少分钟一根K线"`
	Day        string      `json:"day"        dc:"日期"`
}

// FinanceKlineExportModel 导出k线
type FinanceKlineExportModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	KlineType  int         `json:"klineType"  dc:"k线类型"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Volume     int64       `json:"volume"     dc:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   dc:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"唯一key，timestamp和code组合"`
	Scale      int         `json:"scale"      dc:"多少分钟一根K线"`
	Day        string      `json:"day"        dc:"日期"`
}
