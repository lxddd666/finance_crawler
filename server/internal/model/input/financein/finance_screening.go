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

// FinanceScreeningUpdateFields 修改筛股字段过滤
type FinanceScreeningUpdateFields struct {
	Code       string  `json:"code"       dc:"code"`
	Boll       int     `json:"boll"       dc:"满足boll"`
	Macd       int     `json:"macd"       dc:"满足macd"`
	Kdj        int     `json:"kdj"        dc:"满足kdj"`
	Rsi        int     `json:"rsi"        dc:"满足rsi"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	Key        string  `json:"key"        dc:"timestamp和code组合"`
	Day        string  `json:"day"        dc:"日期"`
	Scale      int     `json:"scale"      dc:"分钟一条k线"`
	MatchCount int     `json:"matchCount" dc:"符合条件数量"`
}

// FinanceScreeningInsertFields 新增筛股字段过滤
type FinanceScreeningInsertFields struct {
	Code       string  `json:"code"       dc:"code"`
	Boll       int     `json:"boll"       dc:"满足boll"`
	Macd       int     `json:"macd"       dc:"满足macd"`
	Kdj        int     `json:"kdj"        dc:"满足kdj"`
	Rsi        int     `json:"rsi"        dc:"满足rsi"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	Key        string  `json:"key"        dc:"timestamp和code组合"`
	Day        string  `json:"day"        dc:"日期"`
	Scale      int     `json:"scale"      dc:"分钟一条k线"`
	MatchCount int     `json:"matchCount" dc:"符合条件数量"`
}

// FinanceScreeningEditInp 修改/新增筛股
type FinanceScreeningEditInp struct {
	entity.FinanceScreening
}

func (in *FinanceScreeningEditInp) Filter(ctx context.Context) (err error) {
	// 验证timestamp和code组合
	if err := g.Validator().Rules("required").Data(in.Key).Messages("timestamp和code组合不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceScreeningEditModel struct{}

// FinanceScreeningDeleteInp 删除筛股
type FinanceScreeningDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceScreeningDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceScreeningDeleteModel struct{}

// FinanceScreeningViewInp 获取指定筛股信息
type FinanceScreeningViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceScreeningViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceScreeningViewModel struct {
	entity.FinanceScreening
}

// FinanceScreeningListInp 获取筛股列表
type FinanceScreeningListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceScreeningListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceScreeningListModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	Boll       int         `json:"boll"       dc:"满足boll"`
	Macd       int         `json:"macd"       dc:"满足macd"`
	Kdj        int         `json:"kdj"        dc:"满足kdj"`
	Rsi        int         `json:"rsi"        dc:"满足rsi"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"timestamp和code组合"`
	Day        string      `json:"day"        dc:"日期"`
	Scale      int         `json:"scale"      dc:"分钟一条k线"`
	MatchCount int         `json:"matchCount" dc:"符合条件数量"`
}

// FinanceScreeningExportModel 导出筛股
type FinanceScreeningExportModel struct {
	Id         int64       `json:"id"         dc:"分类ID"`
	Code       string      `json:"code"       dc:"code"`
	Boll       int         `json:"boll"       dc:"满足boll"`
	Macd       int         `json:"macd"       dc:"满足macd"`
	Kdj        int         `json:"kdj"        dc:"满足kdj"`
	Rsi        int         `json:"rsi"        dc:"满足rsi"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	Key        string      `json:"key"        dc:"timestamp和code组合"`
	Day        string      `json:"day"        dc:"日期"`
	Scale      int         `json:"scale"      dc:"分钟一条k线"`
	MatchCount int         `json:"matchCount" dc:"符合条件数量"`
}
