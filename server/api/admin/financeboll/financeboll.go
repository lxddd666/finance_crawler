// Package financeboll
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financeboll

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询boll带列表
type ListReq struct {
	g.Meta `path:"/financeBoll/list" method:"get" tags:"boll带" summary:"获取boll带列表"`
	sysin.FinanceBollListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceBollListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出boll带列表
type ExportReq struct {
	g.Meta `path:"/financeBoll/export" method:"get" tags:"boll带" summary:"导出boll带列表"`
	sysin.FinanceBollListInp
}

type ExportRes struct{}

// ViewReq 获取boll带指定信息
type ViewReq struct {
	g.Meta `path:"/financeBoll/view" method:"get" tags:"boll带" summary:"获取boll带指定信息"`
	sysin.FinanceBollViewInp
}

type ViewRes struct {
	*sysin.FinanceBollViewModel
}

// EditReq 修改/新增boll带
type EditReq struct {
	g.Meta `path:"/financeBoll/edit" method:"post" tags:"boll带" summary:"修改/新增boll带"`
	sysin.FinanceBollEditInp
}

type EditRes struct{}

// DeleteReq 删除boll带
type DeleteReq struct {
	g.Meta `path:"/financeBoll/delete" method:"post" tags:"boll带" summary:"删除boll带"`
	sysin.FinanceBollDeleteInp
}

type DeleteRes struct{}
