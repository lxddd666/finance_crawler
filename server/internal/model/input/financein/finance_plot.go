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
)

// FinancePlotUpdateFields 修改plot图字段过滤
type FinancePlotUpdateFields struct {
	Code      string `json:"code"      dc:"code"`
	Indicator string `json:"indicator" dc:"指标"`
	Day       string `json:"day"       dc:"日期"`
	Path      string `json:"path"      dc:"path"`
}

// FinancePlotInsertFields 新增plot图字段过滤
type FinancePlotInsertFields struct {
	Code      string `json:"code"      dc:"code"`
	Indicator string `json:"indicator" dc:"指标"`
	Day       string `json:"day"       dc:"日期"`
	Path      string `json:"path"      dc:"path"`
}

// FinancePlotEditInp 修改/新增plot图
type FinancePlotEditInp struct {
	entity.FinancePlot
}

func (in *FinancePlotEditInp) Filter(ctx context.Context) (err error) {

	return
}

type FinancePlotEditModel struct{}

// FinancePlotDeleteInp 删除plot图
type FinancePlotDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinancePlotDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancePlotDeleteModel struct{}

// FinancePlotViewInp 获取指定plot图信息
type FinancePlotViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinancePlotViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancePlotViewModel struct {
	entity.FinancePlot
}

// FinancePlotListInp 获取plot图列表
type FinancePlotListInp struct {
	form.PageReq
	Id int64 `json:"id" dc:"分类ID"`
}

func (in *FinancePlotListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinancePlotListModel struct {
	Id        int64  `json:"id"        dc:"分类ID"`
	Code      string `json:"code"      dc:"code"`
	Indicator string `json:"indicator" dc:"指标"`
	Day       string `json:"day"       dc:"日期"`
	Path      string `json:"path"      dc:"path"`
}

// FinancePlotExportModel 导出plot图
type FinancePlotExportModel struct {
	Id        int64  `json:"id"        dc:"分类ID"`
	Code      string `json:"code"      dc:"code"`
	Indicator string `json:"indicator" dc:"指标"`
	Day       string `json:"day"       dc:"日期"`
	Path      string `json:"path"      dc:"path"`
}

type FinancePlotCreate struct {
	Code      string `json:"code"      dc:"code"`
	Indicator string `json:"indicator" dc:"指标"`
	Day       string `json:"day"       dc:"日期"`
}
