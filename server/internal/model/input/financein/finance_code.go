// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// FinanceCodeUpdateFields 修改股票代码字段过滤
type FinanceCodeUpdateFields struct {
	Code     string `json:"code"     dc:"代码"`
	Name     string `json:"name"     dc:"名称"`
	Exchange string `json:"exchange" dc:"交易所"`
}

// FinanceCodeInsertFields 新增股票代码字段过滤
type FinanceCodeInsertFields struct {
	Code     string `json:"code"     dc:"代码"`
	Name     string `json:"name"     dc:"名称"`
	Exchange string `json:"exchange" dc:"交易所"`
}

// FinanceCodeEditInp 修改/新增股票代码
type FinanceCodeEditInp struct {
	entity.FinanceCode
}

func (in *FinanceCodeEditInp) Filter(ctx context.Context) (err error) {
	// 验证代码
	if err := g.Validator().Rules("required").Data(in.Code).Messages("代码不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type FinanceCodeEditModel struct{}

// FinanceCodeDeleteInp 删除股票代码
type FinanceCodeDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceCodeDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceCodeDeleteModel struct{}

// FinanceCodeViewInp 获取指定股票代码信息
type FinanceCodeViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *FinanceCodeViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceCodeViewModel struct {
	entity.FinanceCode
}

// FinanceCodeListInp 获取股票代码列表
type FinanceCodeListInp struct {
	form.PageReq
	Id int64 `json:"id" dc:"分类ID"`
}

func (in *FinanceCodeListInp) Filter(ctx context.Context) (err error) {
	return
}

type FinanceCodeListModel struct {
	Id       int64  `json:"id"       dc:"分类ID"`
	Code     string `json:"code"     dc:"代码"`
	Name     string `json:"name"     dc:"名称"`
	Exchange string `json:"exchange" dc:"交易所"`
}

// FinanceCodeExportModel 导出股票代码
type FinanceCodeExportModel struct {
	Id       int64  `json:"id"       dc:"分类ID"`
	Code     string `json:"code"     dc:"代码"`
	Name     string `json:"name"     dc:"名称"`
	Exchange string `json:"exchange" dc:"交易所"`
}

// FinanceImportCodeInp 获取指定股票代码信息
type FinanceImportCodeInp struct {
	File *ghttp.UploadFile `json:"file" type:"file"    dc:"zip、rar文件"`
}
