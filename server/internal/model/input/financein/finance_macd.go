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

// FinanceMacdUpdateFields 修改macd线字段过滤
type FinanceMacdUpdateFields struct {
	Code         string  `json:"code"         dc:"code"`
	FastPeriod   int     `json:"fastPeriod"   dc:"快线周期"`
	SlowPeriod   int     `json:"slowPeriod"   dc:"慢线周期"`
	SignalPeriod int     `json:"signalPeriod" dc:"信号周期"`
	Dif          float64 `json:"dif"          dc:"EMA12 - EMA26 快线"`
	Dea          float64 `json:"dea"          dc:"DIF的9日EMA（26+9-1)慢线"`
	Macd         float64 `json:"macd"         dc:"macd柱子"`
	Timestamp    int64   `json:"timestamp"    dc:"时间戳"`
	ClosePrice   float64 `json:"closePrice"   dc:"收盘价"`
	KlineNum     int     `json:"klineNum"     dc:"k线根数"`
	Key          string  `json:"key"          dc:"timestamp和code组合"`
	Day          string  `json:"day"          dc:"日期"`
	Scale        int     `json:"scale"        dc:"分钟一条k线"`
}

// FinanceMacdInsertFields 新增macd线字段过滤
type FinanceMacdInsertFields struct {
	Code         string  `json:"code"         dc:"code"`
	FastPeriod   int     `json:"fastPeriod"   dc:"快线周期"`
	SlowPeriod   int     `json:"slowPeriod"   dc:"慢线周期"`
	SignalPeriod int     `json:"signalPeriod" dc:"信号周期"`
	Dif          float64 `json:"dif"          dc:"EMA12 - EMA26 快线"`
	Dea          float64 `json:"dea"          dc:"DIF的9日EMA（26+9-1)慢线"`
	Macd         float64 `json:"macd"         dc:"macd柱子"`
	Timestamp    int64   `json:"timestamp"    dc:"时间戳"`
	ClosePrice   float64 `json:"closePrice"   dc:"收盘价"`
	KlineNum     int     `json:"klineNum"     dc:"k线根数"`
	Key          string  `json:"key"          dc:"timestamp和code组合"`
	Day          string  `json:"day"          dc:"日期"`
	Scale        int     `json:"scale"        dc:"分钟一条k线"`
}

// FinanceMacdEditInp 修改/新增macd线
type FinanceMacdEditInp struct {
	entity.FinanceMacd
}

func (in *FinanceMacdEditInp) Filter(ctx context.Context) (err error) {
	// 验证timestamp和code组合
	if err := g.Validator().Rules("required").Data(in.Key).Messages("timestamp和code组合不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceMacdEditModel struct{}

// FinanceMacdDeleteInp 删除macd线
type FinanceMacdDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceMacdDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceMacdDeleteModel struct{}

// FinanceMacdViewInp 获取指定macd线信息
type FinanceMacdViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceMacdViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceMacdViewModel struct {
	entity.FinanceMacd
}

// FinanceMacdListInp 获取macd线列表
type FinanceMacdListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceMacdListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceMacdListModel struct {
	Id           int64       `json:"id"           dc:"分类ID"`
	Code         string      `json:"code"         dc:"code"`
	FastPeriod   int         `json:"fastPeriod"   dc:"快线周期"`
	SlowPeriod   int         `json:"slowPeriod"   dc:"慢线周期"`
	SignalPeriod int         `json:"signalPeriod" dc:"信号周期"`
	Dif          float64     `json:"dif"          dc:"EMA12 - EMA26 快线"`
	Dea          float64     `json:"dea"          dc:"DIF的9日EMA（26+9-1)慢线"`
	Macd         float64     `json:"macd"         dc:"macd柱子"`
	Timestamp    int64       `json:"timestamp"    dc:"时间戳"`
	ClosePrice   float64     `json:"closePrice"   dc:"收盘价"`
	KlineNum     int         `json:"klineNum"     dc:"k线根数"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	Key          string      `json:"key"          dc:"timestamp和code组合"`
	Day          string      `json:"day"          dc:"日期"`
	Scale        int         `json:"scale"        dc:"分钟一条k线"`
}

// FinanceMacdExportModel 导出macd线
type FinanceMacdExportModel struct {
	Id           int64       `json:"id"           dc:"分类ID"`
	Code         string      `json:"code"         dc:"code"`
	FastPeriod   int         `json:"fastPeriod"   dc:"快线周期"`
	SlowPeriod   int         `json:"slowPeriod"   dc:"慢线周期"`
	SignalPeriod int         `json:"signalPeriod" dc:"信号周期"`
	Dif          float64     `json:"dif"          dc:"EMA12 - EMA26 快线"`
	Dea          float64     `json:"dea"          dc:"DIF的9日EMA（26+9-1)慢线"`
	Macd         float64     `json:"macd"         dc:"macd柱子"`
	Timestamp    int64       `json:"timestamp"    dc:"时间戳"`
	ClosePrice   float64     `json:"closePrice"   dc:"收盘价"`
	KlineNum     int         `json:"klineNum"     dc:"k线根数"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	Key          string      `json:"key"          dc:"timestamp和code组合"`
	Day          string      `json:"day"          dc:"日期"`
	Scale        int         `json:"scale"        dc:"分钟一条k线"`
}
