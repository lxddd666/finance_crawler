// Package financeplot
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package financeplot

import (
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询plot图列表
type ListReq struct {
	g.Meta `path:"/financePlot/list" method:"get" tags:"plot图" summary:"获取plot图列表"`
	sysin.FinancePlotListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.FinancePlotListModel `json:"list"   dc:"数据列表"`
}

// CreatePlotReq 创建plot
type CreatePlotReq struct {
	g.Meta `path:"/financePlot/createPlot" method:"get" tags:"plot图" summary:"获取plot图列表"`
	sysin.FinancePlotCreate
}

type CreatePlotRes struct {
}
