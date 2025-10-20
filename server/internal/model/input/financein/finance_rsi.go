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

// FinanceRsiUpdateFields 修改rsi线字段过滤
type FinanceRsiUpdateFields struct {
	Code           string  `json:"code"           dc:"code"`
	Rsi            float64 `json:"rsi"            dc:"相对强弱指数的数值"`
	SmoothingMa    float64 `json:"smoothingMa"    dc:"平滑移动平均线"`
	BollUpper      int     `json:"bollUpper"      dc:"布林带上轨"`
	BollLower      int     `json:"bollLower"      dc:"布林带下轨"`
	IsBoll         int     `json:"isBoll"         dc:"是否触及布林带边界"`
	BullDivergence int     `json:"bullDivergence" dc:"看涨背离信号"`
	BearDivergence int     `json:"bearDivergence" dc:"看跌背离信号"`
	MaLength       int     `json:"maLength"       dc:"参数"`
	Key            string  `json:"key"            dc:"timestamp和code组合:scalse"`
	Day            string  `json:"day"            dc:"日期"`
	Scale          int     `json:"scale"          dc:"分钟一条k线"`
}

// FinanceRsiInsertFields 新增rsi线字段过滤
type FinanceRsiInsertFields struct {
	Code           string  `json:"code"           dc:"code"`
	Rsi            float64 `json:"rsi"            dc:"相对强弱指数的数值"`
	SmoothingMa    float64 `json:"smoothingMa"    dc:"平滑移动平均线"`
	BollUpper      int     `json:"bollUpper"      dc:"布林带上轨"`
	BollLower      int     `json:"bollLower"      dc:"布林带下轨"`
	IsBoll         int     `json:"isBoll"         dc:"是否触及布林带边界"`
	BullDivergence int     `json:"bullDivergence" dc:"看涨背离信号"`
	BearDivergence int     `json:"bearDivergence" dc:"看跌背离信号"`
	MaLength       int     `json:"maLength"       dc:"参数"`
	Key            string  `json:"key"            dc:"timestamp和code组合:scalse"`
	Day            string  `json:"day"            dc:"日期"`
	Scale          int     `json:"scale"          dc:"分钟一条k线"`
}

// FinanceRsiEditInp 修改/新增rsi线
type FinanceRsiEditInp struct {
	entity.FinanceRsi
}

func (in *FinanceRsiEditInp) Filter(ctx context.Context) (err error) {
	// 验证timestamp和code组合:scalse
	if err := g.Validator().Rules("required").Data(in.Key).Messages("timestamp和code组合:scalse不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceRsiEditModel struct{}

// FinanceRsiDeleteInp 删除rsi线
type FinanceRsiDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceRsiDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceRsiDeleteModel struct{}

// FinanceRsiViewInp 获取指定rsi线信息
type FinanceRsiViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceRsiViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceRsiViewModel struct {
	entity.FinanceRsi
}

// FinanceRsiListInp 获取rsi线列表
type FinanceRsiListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *FinanceRsiListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceRsiListModel struct {
	Id             int64       `json:"id"             dc:"分类ID"`
	Code           string      `json:"code"           dc:"code"`
	Rsi            float64     `json:"rsi"            dc:"相对强弱指数的数值"`
	SmoothingMa    float64     `json:"smoothingMa"    dc:"平滑移动平均线"`
	BollUpper      int         `json:"bollUpper"      dc:"布林带上轨"`
	BollLower      int         `json:"bollLower"      dc:"布林带下轨"`
	IsBoll         int         `json:"isBoll"         dc:"是否触及布林带边界"`
	BullDivergence int         `json:"bullDivergence" dc:"看涨背离信号"`
	BearDivergence int         `json:"bearDivergence" dc:"看跌背离信号"`
	MaLength       int         `json:"maLength"       dc:"参数"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	Key            string      `json:"key"            dc:"timestamp和code组合:scalse"`
	Day            string      `json:"day"            dc:"日期"`
	Scale          int         `json:"scale"          dc:"分钟一条k线"`
}

// FinanceRsiExportModel 导出rsi线
type FinanceRsiExportModel struct {
	Id             int64       `json:"id"             dc:"分类ID"`
	Code           string      `json:"code"           dc:"code"`
	Rsi            float64     `json:"rsi"            dc:"相对强弱指数的数值"`
	SmoothingMa    float64     `json:"smoothingMa"    dc:"平滑移动平均线"`
	BollUpper      int         `json:"bollUpper"      dc:"布林带上轨"`
	BollLower      int         `json:"bollLower"      dc:"布林带下轨"`
	IsBoll         int         `json:"isBoll"         dc:"是否触及布林带边界"`
	BullDivergence int         `json:"bullDivergence" dc:"看涨背离信号"`
	BearDivergence int         `json:"bearDivergence" dc:"看跌背离信号"`
	MaLength       int         `json:"maLength"       dc:"参数"`
	CreatedAt      *gtime.Time `json:"createdAt"      dc:"创建时间"`
	Key            string      `json:"key"            dc:"timestamp和code组合:scalse"`
	Day            string      `json:"day"            dc:"日期"`
	Scale          int         `json:"scale"          dc:"分钟一条k线"`
}
