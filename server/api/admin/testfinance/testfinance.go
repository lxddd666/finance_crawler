// Package testfinance
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package testfinance

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询测试分类列表
type ListReq struct {
	g.Meta `path:"/testFinance/list" method:"get" tags:"测试分类" summary:"获取测试分类列表"`
	sysin.TestFinanceListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.TestFinanceListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出测试分类列表
type ExportReq struct {
	g.Meta `path:"/testFinance/export" method:"get" tags:"测试分类" summary:"导出测试分类列表"`
	sysin.TestFinanceListInp
}

type ExportRes struct{}

// ViewReq 获取测试分类指定信息
type ViewReq struct {
	g.Meta `path:"/testFinance/view" method:"get" tags:"测试分类" summary:"获取测试分类指定信息"`
	sysin.TestFinanceViewInp
}

type ViewRes struct {
	*sysin.TestFinanceViewModel
}

// EditReq 修改/新增测试分类
type EditReq struct {
	g.Meta `path:"/testFinance/edit" method:"post" tags:"测试分类" summary:"修改/新增测试分类"`
	sysin.TestFinanceEditInp
}

type EditRes struct{}

// DeleteReq 删除测试分类
type DeleteReq struct {
	g.Meta `path:"/testFinance/delete" method:"post" tags:"测试分类" summary:"删除测试分类"`
	sysin.TestFinanceDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取测试分类最大排序
type MaxSortReq struct {
	g.Meta `path:"/testFinance/maxSort" method:"get" tags:"测试分类" summary:"获取测试分类最大排序"`
	sysin.TestFinanceMaxSortInp
}

type MaxSortRes struct {
	*sysin.TestFinanceMaxSortModel
}

// StatusReq 更新测试分类状态
type StatusReq struct {
	g.Meta `path:"/testFinance/status" method:"post" tags:"测试分类" summary:"更新测试分类状态"`
	sysin.TestFinanceStatusInp
}

type StatusRes struct{}

// StartReq 更新测试分类状态
type StartReq struct {
	g.Meta `path:"/testFinance/start" method:"post" tags:"测试分类" summary:"开始测试"`
}

type StartRes struct{}
