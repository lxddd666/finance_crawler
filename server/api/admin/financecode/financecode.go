// Package financecode
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financecode

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票代码列表
type ListReq struct {
	g.Meta `path:"/financeCode/list" method:"get" tags:"股票代码" summary:"获取股票代码列表"`
	sysin.FinanceCodeListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceCodeListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票代码列表
type ExportReq struct {
	g.Meta `path:"/financeCode/export" method:"get" tags:"股票代码" summary:"导出股票代码列表"`
	sysin.FinanceCodeListInp
}

type ExportRes struct{}

// ViewReq 获取股票代码指定信息
type ViewReq struct {
	g.Meta `path:"/financeCode/view" method:"get" tags:"股票代码" summary:"获取股票代码指定信息"`
	sysin.FinanceCodeViewInp
}

type ViewRes struct {
	*sysin.FinanceCodeViewModel
}

// EditReq 修改/新增股票代码
type EditReq struct {
	g.Meta `path:"/financeCode/edit" method:"post" tags:"股票代码" summary:"修改/新增股票代码"`
	sysin.FinanceCodeEditInp
}

type EditRes struct{}

// DeleteReq 删除股票代码
type DeleteReq struct {
	g.Meta `path:"/financeCode/delete" method:"post" tags:"股票代码" summary:"删除股票代码"`
	sysin.FinanceCodeDeleteInp
}

type DeleteRes struct{}

// ImportCodeReq 导入股票code
type ImportCodeReq struct {
	g.Meta `path:"/financeCode/import" method:"post" tags:"股票代码" summary:"导入股票代码"`
	sysin.FinanceImportCodeInp
}

type ImportCodeRes struct{}
