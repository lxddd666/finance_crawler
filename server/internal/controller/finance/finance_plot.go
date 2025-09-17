// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financeplot"
	"hotgo/internal/service"
)

var (
	FinancePlot = cFinancePlot{}
)

type cFinancePlot struct{}

// List 查看plot图列表
func (c *cFinancePlot) List(ctx context.Context, req *financeplot.ListReq) (res *financeplot.ListRes, err error) {
	//list, totalCount, err := service.SysFinanceKline().List(ctx, &req.FinanceKlineListInp)

	return
}

func (c *cFinancePlot) CreatePlot(ctx context.Context, req *financeplot.CreatePlotReq) (res *financeplot.CreatePlotRes, err error) {
	err = service.SysFinancePlot().CreatePlot(ctx, req.FinancePlotCreate)

	return
}
