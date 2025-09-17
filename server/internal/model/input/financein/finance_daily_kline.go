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

// FinanceDailyKlineUpdateFields 修改股票日K线数据表字段过滤
type FinanceDailyKlineUpdateFields struct {
	Code       string  `json:"code"       dc:"股票代码"`
	KlineType  int     `json:"klineType"  dc:"K线类型: 0-普通 1-复权"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64 `json:"openPrice"  dc:"开盘价"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"最低价"`
	Volume     int64   `json:"volume"     dc:"成交量"`
	Turnover   float64 `json:"turnover"   dc:"成交额"`
	Md5        float64 `json:"md5"        dc:"5日均线"`
	Md10       float64 `json:"md10"       dc:"10日均线"`
	Md20       float64 `json:"md20"       dc:"20日均线"`
	Md30       float64 `json:"md30"       dc:"30日均线"`
	Md60       float64 `json:"md60"       dc:"60日均线"`
	UniqueKey  string  `json:"uniqueKey"  dc:"唯一键: timestamp_code_scale"`
	Scale      int     `json:"scale"      dc:"K线周期(分钟)"`
	Day        string  `json:"day"        dc:"日期(yyyy-MM-dd)"`
}

// FinanceDailyKlineInsertFields 新增股票日K线数据表字段过滤
type FinanceDailyKlineInsertFields struct {
	Code       string  `json:"code"       dc:"股票代码"`
	KlineType  int     `json:"klineType"  dc:"K线类型: 0-普通 1-复权"`
	Timestamp  int64   `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64 `json:"openPrice"  dc:"开盘价"`
	ClosePrice float64 `json:"closePrice" dc:"收盘价"`
	HighPrice  float64 `json:"highPrice"  dc:"最高价"`
	LowPrice   float64 `json:"lowPrice"   dc:"最低价"`
	Volume     int64   `json:"volume"     dc:"成交量"`
	Turnover   float64 `json:"turnover"   dc:"成交额"`
	Md5        float64 `json:"md5"        dc:"5日均线"`
	Md10       float64 `json:"md10"       dc:"10日均线"`
	Md20       float64 `json:"md20"       dc:"20日均线"`
	Md30       float64 `json:"md30"       dc:"30日均线"`
	Md60       float64 `json:"md60"       dc:"60日均线"`
	UniqueKey  string  `json:"uniqueKey"  dc:"唯一键: timestamp_code_scale"`
	Scale      int     `json:"scale"      dc:"K线周期(分钟)"`
	Day        string  `json:"day"        dc:"日期(yyyy-MM-dd)"`
}

// FinanceDailyKlineEditInp 修改/新增股票日K线数据表
type FinanceDailyKlineEditInp struct {
	entity.FinanceDailyKline
}

func (in *FinanceDailyKlineEditInp) Filter(ctx context.Context) (err error) {
	// 验证唯一键: timestamp_code_scale
	if err := g.Validator().Rules("required").Data(in.Key).Messages("唯一键: timestamp_code_scale不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证K线周期(分钟)
	if err := g.Validator().Rules("required").Data(in.Scale).Messages("K线周期(分钟)不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceDailyKlineEditModel struct{}

// FinanceDailyKlineDeleteInp 删除股票日K线数据表
type FinanceDailyKlineDeleteInp struct {
	Id interface{} `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *FinanceDailyKlineDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceDailyKlineDeleteModel struct{}

// FinanceDailyKlineViewInp 获取指定股票日K线数据表信息
type FinanceDailyKlineViewInp struct {
	Id int64 `json:"id" v:"required#主键ID不能为空" dc:"主键ID"`
}

func (in *FinanceDailyKlineViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceDailyKlineViewModel struct {
	entity.FinanceDailyKline
}

// FinanceDailyKlineListInp 获取股票日K线数据表列表
type FinanceDailyKlineListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"主键ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceDailyKlineListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceDailyKlineListModel struct {
	Id         int64       `json:"id"         dc:"主键ID"`
	Code       string      `json:"code"       dc:"股票代码"`
	KlineType  int         `json:"klineType"  dc:"K线类型: 0-普通 1-复权"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"最低价"`
	Volume     int64       `json:"volume"     dc:"成交量"`
	Turnover   float64     `json:"turnover"   dc:"成交额"`
	Md5        float64     `json:"md5"        dc:"5日均线"`
	Md10       float64     `json:"md10"       dc:"10日均线"`
	Md20       float64     `json:"md20"       dc:"20日均线"`
	Md30       float64     `json:"md30"       dc:"30日均线"`
	Md60       float64     `json:"md60"       dc:"60日均线"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UniqueKey  string      `json:"uniqueKey"  dc:"唯一键: timestamp_code_scale"`
	Scale      int         `json:"scale"      dc:"K线周期(分钟)"`
	Day        string      `json:"day"        dc:"日期(yyyy-MM-dd)"`
}

// FinanceDailyKlineExportModel 导出股票日K线数据表
type FinanceDailyKlineExportModel struct {
	Id         int64       `json:"id"         dc:"主键ID"`
	Code       string      `json:"code"       dc:"股票代码"`
	KlineType  int         `json:"klineType"  dc:"K线类型: 0-普通 1-复权"`
	Timestamp  int64       `json:"timestamp"  dc:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  dc:"开盘价"`
	ClosePrice float64     `json:"closePrice" dc:"收盘价"`
	HighPrice  float64     `json:"highPrice"  dc:"最高价"`
	LowPrice   float64     `json:"lowPrice"   dc:"最低价"`
	Volume     int64       `json:"volume"     dc:"成交量"`
	Turnover   float64     `json:"turnover"   dc:"成交额"`
	Md5        float64     `json:"md5"        dc:"5日均线"`
	Md10       float64     `json:"md10"       dc:"10日均线"`
	Md20       float64     `json:"md20"       dc:"20日均线"`
	Md30       float64     `json:"md30"       dc:"30日均线"`
	Md60       float64     `json:"md60"       dc:"60日均线"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UniqueKey  string      `json:"uniqueKey"  dc:"唯一键: timestamp_code_scale"`
	Scale      int         `json:"scale"      dc:"K线周期(分钟)"`
	Day        string      `json:"day"        dc:"日期(yyyy-MM-dd)"`
}
