// Package financemacd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financemacd

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询macd线列表
type ListReq struct {
	g.Meta `path:"/financeMacd/list" method:"get" tags:"macd线" summary:"获取macd线列表"`
	sysin.FinanceMacdListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceMacdListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出macd线列表
type ExportReq struct {
	g.Meta `path:"/financeMacd/export" method:"get" tags:"macd线" summary:"导出macd线列表"`
	sysin.FinanceMacdListInp
}

type ExportRes struct{}

// ViewReq 获取macd线指定信息
type ViewReq struct {
	g.Meta `path:"/financeMacd/view" method:"get" tags:"macd线" summary:"获取macd线指定信息"`
	sysin.FinanceMacdViewInp
}

type ViewRes struct {
	*sysin.FinanceMacdViewModel
}

// EditReq 修改/新增macd线
type EditReq struct {
	g.Meta `path:"/financeMacd/edit" method:"post" tags:"macd线" summary:"修改/新增macd线"`
	sysin.FinanceMacdEditInp
}

type EditRes struct{}

// DeleteReq 删除macd线
type DeleteReq struct {
	g.Meta `path:"/financeMacd/delete" method:"post" tags:"macd线" summary:"删除macd线"`
	sysin.FinanceMacdDeleteInp
}

type DeleteRes struct{}
