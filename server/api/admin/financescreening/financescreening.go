// Package financescreening
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financescreening

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询筛股列表
type ListReq struct {
	g.Meta `path:"/financeScreening/list" method:"get" tags:"筛股" summary:"获取筛股列表"`
	sysin.FinanceScreeningListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceScreeningListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出筛股列表
type ExportReq struct {
	g.Meta `path:"/financeScreening/export" method:"get" tags:"筛股" summary:"导出筛股列表"`
	sysin.FinanceScreeningListInp
}

type ExportRes struct{}

// ViewReq 获取筛股指定信息
type ViewReq struct {
	g.Meta `path:"/financeScreening/view" method:"get" tags:"筛股" summary:"获取筛股指定信息"`
	sysin.FinanceScreeningViewInp
}

type ViewRes struct {
	*sysin.FinanceScreeningViewModel
}

// EditReq 修改/新增筛股
type EditReq struct {
	g.Meta `path:"/financeScreening/edit" method:"post" tags:"筛股" summary:"修改/新增筛股"`
	sysin.FinanceScreeningEditInp
}

type EditRes struct{}

// DeleteReq 删除筛股
type DeleteReq struct {
	g.Meta `path:"/financeScreening/delete" method:"post" tags:"筛股" summary:"删除筛股"`
	sysin.FinanceScreeningDeleteInp
}

type DeleteRes struct{}

// ScreeningDailyReq 删除筛股
type ScreeningDailyReq struct {
	g.Meta `path:"/financeScreening/screeningDaily" method:"post" tags:"筛股" summary:"日常筛股"`
}

type ScreeningDailyRes struct{}
