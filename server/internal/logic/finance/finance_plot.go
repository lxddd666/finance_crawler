// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
)

type sSysFinancePlot struct{}

func NewSysFinancePlot() *sSysFinancePlot {
	return &sSysFinancePlot{}
}

func init() {
	service.RegisterSysFinancePlot(NewSysFinancePlot())
}

// Model plot图ORM模型
func (s *sSysFinancePlot) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinancePlot.Ctx(ctx), option...)
}

// List 获取plot图列表
func (s *sSysFinancePlot) List(ctx context.Context, in *sysin.FinancePlotListInp) (list []*sysin.FinancePlotListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.FinancePlotListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.FinancePlot.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.FinancePlot.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取plot图列表失败，请稍后重试！")
		return
	}
	return
}

// CreatePlot 指标图创建
func (s *sSysFinancePlot) CreatePlot(ctx context.Context, in sysin.FinancePlotCreate) (err error) {
	switch in.Indicator {
	case consts.PlotIndicatorKline:
		s.KlinePlot(ctx, in.Code)
	case consts.PlotIndicatorMacd:
		s.MacdPlot(ctx, in.Code)
	case consts.PlotIndicatorBoll:
		s.BollPlot(ctx, in.Code)
	case consts.PlotIndicatorKdj:
		s.KdjPlot(ctx, in.Code)
	case consts.PlotIndicatorRsi:
		s.RsiPlot(ctx, in.Code)
	}
	return
}

func (s *sSysFinancePlot) KlinePlot(ctx context.Context, code string) {
	return
}

func (s *sSysFinancePlot) MacdPlot(ctx context.Context, code string) {
	return
}

func (s *sSysFinancePlot) KdjPlot(ctx context.Context, code string) {
	return
}

func (s *sSysFinancePlot) BollPlot(ctx context.Context, code string) {
	return
}

func (s *sSysFinancePlot) RsiPlot(ctx context.Context, code string) {
	return
}
