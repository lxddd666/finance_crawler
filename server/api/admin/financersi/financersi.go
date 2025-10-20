// Package financersi
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financersi

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询rsi线列表
type ListReq struct {
	g.Meta `path:"/financeRsi/list" method:"get" tags:"rsi线" summary:"获取rsi线列表"`
	sysin.FinanceRsiListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceRsiListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出rsi线列表
type ExportReq struct {
	g.Meta `path:"/financeRsi/export" method:"get" tags:"rsi线" summary:"导出rsi线列表"`
	sysin.FinanceRsiListInp
}

type ExportRes struct{}

// ViewReq 获取rsi线指定信息
type ViewReq struct {
	g.Meta `path:"/financeRsi/view" method:"get" tags:"rsi线" summary:"获取rsi线指定信息"`
	sysin.FinanceRsiViewInp
}

type ViewRes struct {
	*sysin.FinanceRsiViewModel
}

// EditReq 修改/新增rsi线
type EditReq struct {
	g.Meta `path:"/financeRsi/edit" method:"post" tags:"rsi线" summary:"修改/新增rsi线"`
	sysin.FinanceRsiEditInp
}

type EditRes struct{}

// DeleteReq 删除rsi线
type DeleteReq struct {
	g.Meta `path:"/financeRsi/delete" method:"post" tags:"rsi线" summary:"删除rsi线"`
	sysin.FinanceRsiDeleteInp
}

type DeleteRes struct{}
