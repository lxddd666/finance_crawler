// Package financealltickresponse
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financealltickresponse

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询alltick返回值列表
type ListReq struct {
	g.Meta `path:"/financeAlltickResponse/list" method:"get" tags:"alltick返回值" summary:"获取alltick返回值列表"`
	sysin.FinanceAlltickResponseListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceAlltickResponseListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出alltick返回值列表
type ExportReq struct {
	g.Meta `path:"/financeAlltickResponse/export" method:"get" tags:"alltick返回值" summary:"导出alltick返回值列表"`
	sysin.FinanceAlltickResponseListInp
}

type ExportRes struct{}

// ViewReq 获取alltick返回值指定信息
type ViewReq struct {
	g.Meta `path:"/financeAlltickResponse/view" method:"get" tags:"alltick返回值" summary:"获取alltick返回值指定信息"`
	sysin.FinanceAlltickResponseViewInp
}

type ViewRes struct {
	*sysin.FinanceAlltickResponseViewModel
}

// EditReq 修改/新增alltick返回值
type EditReq struct {
	g.Meta `path:"/financeAlltickResponse/edit" method:"post" tags:"alltick返回值" summary:"修改/新增alltick返回值"`
	sysin.FinanceAlltickResponseEditInp
}

type EditRes struct{}

// DeleteReq 删除alltick返回值
type DeleteReq struct {
	g.Meta `path:"/financeAlltickResponse/delete" method:"post" tags:"alltick返回值" summary:"删除alltick返回值"`
	sysin.FinanceAlltickResponseDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取alltick返回值最大排序
type MaxSortReq struct {
	g.Meta `path:"/financeAlltickResponse/maxSort" method:"get" tags:"alltick返回值" summary:"获取alltick返回值最大排序"`
	sysin.FinanceAlltickResponseMaxSortInp
}

type MaxSortRes struct {
	*sysin.FinanceAlltickResponseMaxSortModel
}
