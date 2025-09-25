// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
	"image/color"
	"log"
	"os"
	"time"
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

func (s *sSysFinancePlot) TrendPlot(ctx context.Context, klineList []*entity.FinanceKline, trendList []*entity.FinanceKline) (err error) {
	codeInfo := klineList[0]

	p := plot.New()
	p.Title.Text = "Finance Trend Chart"
	p.Title.TextStyle.Font.Size = vg.Points(16)
	p.X.Label.Text = "Date"
	p.Y.Label.Text = "Close Price"

	// 设置X轴为时间格式
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	// 自定义颜色
	color1 := color.Black                            // klineList - 黑色
	color2 := color.RGBA{R: 255, G: 0, B: 0, A: 255} // trendList - 红色

	// 转换数据
	points1, err := createXYPoints(klineList)
	if err != nil {
		return err
	}

	points2, err := createXYPoints(trendList)
	if err != nil {
		return err
	}

	// 创建第一条线（klineList - 黑色实线）
	line1, err := plotter.NewLine(points1)
	if err != nil {
		return err
	}
	line1.LineStyle.Width = vg.Points(1.5)
	line1.LineStyle.Color = color1

	// 第一条线的散点（较小的黑点）
	scatter1, err := plotter.NewScatter(points1)
	if err != nil {
		return err
	}
	scatter1.GlyphStyle.Color = color1
	scatter1.GlyphStyle.Radius = vg.Points(2) // 较小的点
	scatter1.GlyphStyle.Shape = draw.CircleGlyph{}

	// 创建第二条线（trendList - 红色虚线）
	line2, err := plotter.NewLine(points2)
	if err != nil {
		return err
	}
	line2.LineStyle.Width = vg.Points(2)
	line2.LineStyle.Color = color2
	line2.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// 第二条线的散点（较大的红色实心圆）
	scatter2, err := plotter.NewScatter(points2)
	if err != nil {
		return err
	}
	scatter2.GlyphStyle.Color = color2
	scatter2.GlyphStyle.Radius = vg.Points(5) // 更大的点，更明显
	scatter2.GlyphStyle.Shape = draw.CircleGlyph{}
	// 可选：添加白色边框使红点更突出
	scatter2.GlyphStyle.Shape = draw.RingGlyph{}

	// 添加到图表
	p.Add(line1, scatter1, line2, scatter2)
	p.Legend.Add("Kline Data", line1)
	p.Legend.Add("Trend Line", line2)
	p.Legend.TextStyle.Font.Size = vg.Points(12)

	// 保存为PNG文件
	saveDir := "D:\\go\\src\\finance_crawlerV2\\hotgo\\server\\internal\\logic\\finance\\image\\"
	savePath := saveDir + fmt.Sprintf("%s_finance_trend.png", codeInfo.Code)
	if _, err = os.Stat(saveDir); os.IsNotExist(err) {
		if err = os.MkdirAll(saveDir, 0755); err != nil {
			return fmt.Errorf("创建文件夹失败: %v", err)
		}
		log.Printf("创建文件夹: %s", saveDir)
	}
	if err = p.Save(10*vg.Inch, 6*vg.Inch, savePath); err != nil {
		return err
	}

	log.Println("趋势图表已保存为: " + savePath)
	return nil
}

// createXYPoints 将FinanceKline列表转换为plotter.XYs格式
func createXYPoints(klineList []*entity.FinanceKline) (plotter.XYs, error) {
	points := make(plotter.XYs, len(klineList))

	for i, kline := range klineList {
		// 解析日期字符串为时间类型
		t, err := time.Parse("2006-01-02", kline.Day)
		if err != nil {
			return nil, err
		}

		// 使用Unix时间戳作为X轴值
		points[i].X = float64(t.Unix())
		points[i].Y = kline.ClosePrice
	}

	return points, nil
}
