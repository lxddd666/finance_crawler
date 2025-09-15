// Package financedailykline
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financedailykline

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询股票日K线数据表列表
type ListReq struct {
	g.Meta `path:"/financeDailyKline/list" method:"get" tags:"股票日K线数据表" summary:"获取股票日K线数据表列表"`
	sysin.FinanceDailyKlineListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinanceDailyKlineListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出股票日K线数据表列表
type ExportReq struct {
	g.Meta `path:"/financeDailyKline/export" method:"get" tags:"股票日K线数据表" summary:"导出股票日K线数据表列表"`
	sysin.FinanceDailyKlineListInp
}

type ExportRes struct{}

// ViewReq 获取股票日K线数据表指定信息
type ViewReq struct {
	g.Meta `path:"/financeDailyKline/view" method:"get" tags:"股票日K线数据表" summary:"获取股票日K线数据表指定信息"`
	sysin.FinanceDailyKlineViewInp
}

type ViewRes struct {
	*sysin.FinanceDailyKlineViewModel
}

// EditReq 修改/新增股票日K线数据表
type EditReq struct {
	g.Meta `path:"/financeDailyKline/edit" method:"post" tags:"股票日K线数据表" summary:"修改/新增股票日K线数据表"`
	sysin.FinanceDailyKlineEditInp
}

type EditRes struct{}

// DeleteReq 删除股票日K线数据表
type DeleteReq struct {
	g.Meta `path:"/financeDailyKline/delete" method:"post" tags:"股票日K线数据表" summary:"删除股票日K线数据表"`
	sysin.FinanceDailyKlineDeleteInp
}

type DeleteRes struct{}
