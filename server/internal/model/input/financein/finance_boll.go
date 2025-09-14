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

// FinanceBollUpdateFields 修改boll带字段过滤
type FinanceBollUpdateFields struct {
	Code              string  `json:"code"              dc:"code"`
	KlineId           int64   `json:"klineId"           dc:"k线id"`
	KlineType         int     `json:"klineType"         dc:"k线类型"`
	Timestamp         int64   `json:"timestamp"         dc:"时间戳"`
	MiddleBand        float64 `json:"middleBand"        dc:"中轨（移动平均线）"`
	UpperBand         float64 `json:"upperBand"         dc:"上轨"`
	LowerBand         float64 `json:"lowerBand"         dc:"下轨"`
	StandardDeviation float64 `json:"standardDeviation" dc:"标准差"`
	ClosePrice        float64 `json:"closePrice"        dc:"收盘价"`
	KlineNum          int     `json:"klineNum"          dc:"k线根数"`
	Multiple          int     `json:"multiple"          dc:"标准差倍数"`
	Key               string  `json:"key"               dc:"timestamp和code组合"`
	Day               string  `json:"day"               dc:"日期"`
	Scale             int     `json:"scale"             dc:"分钟一条k线"`
	Degree            float64 `json:"degree"            dc:"层度"`
}

// FinanceBollInsertFields 新增boll带字段过滤
type FinanceBollInsertFields struct {
	Code              string  `json:"code"              dc:"code"`
	KlineId           int64   `json:"klineId"           dc:"k线id"`
	KlineType         int     `json:"klineType"         dc:"k线类型"`
	Timestamp         int64   `json:"timestamp"         dc:"时间戳"`
	MiddleBand        float64 `json:"middleBand"        dc:"中轨（移动平均线）"`
	UpperBand         float64 `json:"upperBand"         dc:"上轨"`
	LowerBand         float64 `json:"lowerBand"         dc:"下轨"`
	StandardDeviation float64 `json:"standardDeviation" dc:"标准差"`
	ClosePrice        float64 `json:"closePrice"        dc:"收盘价"`
	KlineNum          int     `json:"klineNum"          dc:"k线根数"`
	Multiple          int     `json:"multiple"          dc:"标准差倍数"`
	Key               string  `json:"key"               dc:"timestamp和code组合"`
	Day               string  `json:"day"               dc:"日期"`
	Scale             int     `json:"scale"             dc:"分钟一条k线"`
	Degree            float64 `json:"degree"            dc:"层度"`
}

// FinanceBollEditInp 修改/新增boll带
type FinanceBollEditInp struct {
	entity.FinanceBoll
}

func (in *FinanceBollEditInp) Filter(ctx context.Context) (err error) {
	// 验证timestamp和code组合
	if err := g.Validator().Rules("required").Data(in.Key).Messages("timestamp和code组合不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceBollEditModel struct{}

// FinanceBollDeleteInp 删除boll带
type FinanceBollDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceBollDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceBollDeleteModel struct{}

// FinanceBollViewInp 获取指定boll带信息
type FinanceBollViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceBollViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceBollViewModel struct {
	entity.FinanceBoll
}

// FinanceBollListInp 获取boll带列表
type FinanceBollListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceBollListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceBollListModel struct {
	Id                int64       `json:"id"                dc:"分类ID"`
	Code              string      `json:"code"              dc:"code"`
	KlineId           int64       `json:"klineId"           dc:"k线id"`
	KlineType         int         `json:"klineType"         dc:"k线类型"`
	Timestamp         int64       `json:"timestamp"         dc:"时间戳"`
	MiddleBand        float64     `json:"middleBand"        dc:"中轨（移动平均线）"`
	UpperBand         float64     `json:"upperBand"         dc:"上轨"`
	LowerBand         float64     `json:"lowerBand"         dc:"下轨"`
	StandardDeviation float64     `json:"standardDeviation" dc:"标准差"`
	ClosePrice        float64     `json:"closePrice"        dc:"收盘价"`
	KlineNum          int         `json:"klineNum"          dc:"k线根数"`
	Multiple          int         `json:"multiple"          dc:"标准差倍数"`
	CreatedAt         *gtime.Time `json:"createdAt"         dc:"创建时间"`
	Key               string      `json:"key"               dc:"timestamp和code组合"`
	Day               string      `json:"day"               dc:"日期"`
	Scale             int         `json:"scale"             dc:"分钟一条k线"`
	Degree            float64     `json:"degree"            dc:"层度"`
}

// FinanceBollExportModel 导出boll带
type FinanceBollExportModel struct {
	Id                int64       `json:"id"                dc:"分类ID"`
	Code              string      `json:"code"              dc:"code"`
	KlineId           int64       `json:"klineId"           dc:"k线id"`
	KlineType         int         `json:"klineType"         dc:"k线类型"`
	Timestamp         int64       `json:"timestamp"         dc:"时间戳"`
	MiddleBand        float64     `json:"middleBand"        dc:"中轨（移动平均线）"`
	UpperBand         float64     `json:"upperBand"         dc:"上轨"`
	LowerBand         float64     `json:"lowerBand"         dc:"下轨"`
	StandardDeviation float64     `json:"standardDeviation" dc:"标准差"`
	ClosePrice        float64     `json:"closePrice"        dc:"收盘价"`
	KlineNum          int         `json:"klineNum"          dc:"k线根数"`
	Multiple          int         `json:"multiple"          dc:"标准差倍数"`
	CreatedAt         *gtime.Time `json:"createdAt"         dc:"创建时间"`
	Key               string      `json:"key"               dc:"timestamp和code组合"`
	Day               string      `json:"day"               dc:"日期"`
	Scale             int         `json:"scale"             dc:"分钟一条k线"`
	Degree            float64     `json:"degree"            dc:"层度"`
}
