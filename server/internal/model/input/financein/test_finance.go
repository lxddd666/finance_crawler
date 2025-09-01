// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sysin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestFinanceUpdateFields 修改测试分类字段过滤
type TestFinanceUpdateFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Remark      string `json:"remark"      dc:"备注"`
	Status      int    `json:"status"      dc:"状态"`
}

// TestFinanceInsertFields 新增测试分类字段过滤
type TestFinanceInsertFields struct {
	Name        string `json:"name"        dc:"分类名称"`
	ShortName   string `json:"shortName"   dc:"简称"`
	Description string `json:"description" dc:"描述"`
	Sort        int    `json:"sort"        dc:"排序"`
	Remark      string `json:"remark"      dc:"备注"`
	Status      int    `json:"status"      dc:"状态"`
}

// TestFinanceEditInp 修改/新增测试分类
type TestFinanceEditInp struct {
	entity.TestFinance
}

func (in *TestFinanceEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("分类名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}

	return
}

type TestFinanceEditModel struct{}

// TestFinanceDeleteInp 删除测试分类
type TestFinanceDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestFinanceDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TestFinanceDeleteModel struct{}

// TestFinanceViewInp 获取指定测试分类信息
type TestFinanceViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *TestFinanceViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TestFinanceViewModel struct {
	entity.TestFinance
}

// TestFinanceListInp 获取测试分类列表
type TestFinanceListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TestFinanceListInp) Filter(ctx context.Context) (err error) {
	return
}

type TestFinanceListModel struct {
	Id          int64       `json:"id"          dc:"分类ID"`
	Name        string      `json:"name"        dc:"分类名称"`
	ShortName   string      `json:"shortName"   dc:"简称"`
	Description string      `json:"description" dc:"描述"`
	Sort        int         `json:"sort"        dc:"排序"`
	Remark      string      `json:"remark"      dc:"备注"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// TestFinanceExportModel 导出测试分类
type TestFinanceExportModel struct {
	Id          int64       `json:"id"          dc:"分类ID"`
	Name        string      `json:"name"        dc:"分类名称"`
	ShortName   string      `json:"shortName"   dc:"简称"`
	Description string      `json:"description" dc:"描述"`
	Sort        int         `json:"sort"        dc:"排序"`
	Remark      string      `json:"remark"      dc:"备注"`
	Status      int         `json:"status"      dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// TestFinanceMaxSortInp 获取测试分类最大排序
type TestFinanceMaxSortInp struct{}

func (in *TestFinanceMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type TestFinanceMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// TestFinanceStatusInp 更新测试分类状态
type TestFinanceStatusInp struct {
	Id     int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *TestFinanceStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("分类ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type TestFinanceStatusModel struct{}
