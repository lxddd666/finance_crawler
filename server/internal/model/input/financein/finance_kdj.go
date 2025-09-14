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

// FinanceKdjUpdateFields 修改kdj字段过滤
type FinanceKdjUpdateFields struct {
	Code       string  `json:"code"       dc:"code"`
	K          float64 `json:"k"          dc:"k值"`
	D          float64 `json:"d"          dc:"D值"`
	J          float64 `json:"j"          dc:"J值"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"该K线最低价"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	Key        string  `json:"key"        dc:"timestamp和code组合"`
	Day        string  `json:"day"        dc:"日期"`
	Scale      int     `json:"scale"      dc:"分钟一条k线"`
}

// FinanceKdjInsertFields 新增kdj字段过滤
type FinanceKdjInsertFields struct {
	Code       string  `json:"code"       dc:"code"`
	K          float64 `json:"k"          dc:"k值"`
	D          float64 `json:"d"          dc:"D值"`
	J          float64 `json:"j"          dc:"J值"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"该K线最低价"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	Key        string  `json:"key"        dc:"timestamp和code组合"`
	Day        string  `json:"day"        dc:"日期"`
	Scale      int     `json:"scale"      dc:"分钟一条k线"`
}

// FinanceKdjEditInp 修改/新增kdj
type FinanceKdjEditInp struct {
	entity.FinanceKdj
}

func (in *FinanceKdjEditInp) Filter(ctx context.Context) (err error) {
	// 验证timestamp和code组合
	if err := g.Validator().Rules("required").Data(in.Key).Messages("timestamp和code组合不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceKdjEditModel struct{}

// FinanceKdjDeleteInp 删除kdj
type FinanceKdjDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceKdjDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKdjDeleteModel struct{}

// FinanceKdjViewInp 获取指定kdj信息
type FinanceKdjViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceKdjViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKdjViewModel struct {
	entity.FinanceKdj
}

// FinanceKdjListInp 获取kdj列表
type FinanceKdjListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceKdjListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceKdjListModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	K          float64     `json:"k"          dc:"k值"`
	D          float64     `json:"d"          dc:"D值"`
	J          float64     `json:"j"          dc:"J值"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"timestamp和code组合"`
	Day        string      `json:"day"        dc:"日期"`
	Scale      int         `json:"scale"      dc:"分钟一条k线"`
}

// FinanceKdjExportModel 导出kdj
type FinanceKdjExportModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	K          float64     `json:"k"          dc:"k值"`
	D          float64     `json:"d"          dc:"D值"`
	J          float64     `json:"j"          dc:"J值"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"该K线最低价"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"timestamp和code组合"`
	Day        string      `json:"day"        dc:"日期"`
	Scale      int         `json:"scale"      dc:"分钟一条k线"`
}

// 买入信号类型
type BuySignal struct {
	IsValid      bool
	SignalType   string
	Strength     string // 信号强度: 弱, 中, 强
	Description  string
	CurrentValue *entity.FinanceKdj
}
