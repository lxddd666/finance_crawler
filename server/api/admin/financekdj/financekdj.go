// Package financekdj
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financekdj

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询kdj列表
type ListReq struct {
	g.Meta `path:"/financeKdj/list" method:"get" tags:"kdj" summary:"获取kdj列表"`
	sysin.FinanceKdjListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceKdjListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出kdj列表
type ExportReq struct {
	g.Meta `path:"/financeKdj/export" method:"get" tags:"kdj" summary:"导出kdj列表"`
	sysin.FinanceKdjListInp
}

type ExportRes struct{}

// ViewReq 获取kdj指定信息
type ViewReq struct {
	g.Meta `path:"/financeKdj/view" method:"get" tags:"kdj" summary:"获取kdj指定信息"`
	sysin.FinanceKdjViewInp
}

type ViewRes struct {
	*sysin.FinanceKdjViewModel
}

// EditReq 修改/新增kdj
type EditReq struct {
	g.Meta `path:"/financeKdj/edit" method:"post" tags:"kdj" summary:"修改/新增kdj"`
	sysin.FinanceKdjEditInp
}

type EditRes struct{}

// DeleteReq 删除kdj
type DeleteReq struct {
	g.Meta `path:"/financeKdj/delete" method:"post" tags:"kdj" summary:"删除kdj"`
	sysin.FinanceKdjDeleteInp
}

type DeleteRes struct{}
