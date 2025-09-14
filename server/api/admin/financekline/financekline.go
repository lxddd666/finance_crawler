// Package financekline
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financekline

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询k线列表
type ListReq struct {
	g.Meta `path:"/financeKline/list" method:"get" tags:"k线" summary:"获取k线列表"`
	sysin.FinanceKlineListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceKlineListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出k线列表
type ExportReq struct {
	g.Meta `path:"/financeKline/export" method:"get" tags:"k线" summary:"导出k线列表"`
	sysin.FinanceKlineListInp
}

type ExportRes struct{}

// ViewReq 获取k线指定信息
type ViewReq struct {
	g.Meta `path:"/financeKline/view" method:"get" tags:"k线" summary:"获取k线指定信息"`
	sysin.FinanceKlineViewInp
}

type ViewRes struct {
	*sysin.FinanceKlineViewModel
}

// EditReq 修改/新增k线
type EditReq struct {
	g.Meta `path:"/financeKline/edit" method:"post" tags:"k线" summary:"修改/新增k线"`
	sysin.FinanceKlineEditInp
}

type EditRes struct{}

// DeleteReq 删除k线
type DeleteReq struct {
	g.Meta `path:"/financeKline/delete" method:"post" tags:"k线" summary:"删除k线"`
	sysin.FinanceKlineDeleteInp
}

type DeleteRes struct{}
