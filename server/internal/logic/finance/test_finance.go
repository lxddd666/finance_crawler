// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"math"
)

type sSysTestFinance struct{}

// FinanceAlltickRequestData 用于JSON序列化的数据结构
type FinanceAlltickRequestData struct {
	Code              string `json:"code"`
	KlineType         int    `json:"kline_type"`
	KlineTimestampEnd int    `json:"kline_timestamp_end"`
	AdjustType        int    `json:"adjust_type"`
	QueryKlineNum     int    `json:"query_kline_num"`
}

// FinanceAlltickRequestQuery 用于JSON序列化的完整查询结构
type FinanceAlltickRequestQuery struct {
	Trace string                    `json:"trace"`
	Data  FinanceAlltickRequestData `json:"data"`
}

func NewSysTestFinance() *sSysTestFinance {
	return &sSysTestFinance{}
}

func init() {
	service.RegisterSysTestFinance(NewSysTestFinance())
}

// ConvertToQueryString 将 FinanceAlltickRequest 结构体转换为 JSON 查询字符串
func (s *sSysTestFinance) ConvertToQueryString(req *entity.FinanceAlltickRequest) (string, error) {
	// 如果 trace 为空，使用默认值

	trace := "1111111111111111111111111"

	query := FinanceAlltickRequestQuery{
		Trace: trace,
		Data: FinanceAlltickRequestData{
			Code:              req.Code,
			KlineType:         req.KlineType,
			KlineTimestampEnd: req.KlineTimestampEnd,
			AdjustType:        req.AdjustType,
			QueryKlineNum:     req.QueryKlineNum,
		},
	}

	jsonBytes, err := json.Marshal(query)
	if err != nil {
		return "", fmt.Errorf("JSON序列化失败: %v", err)
	}

	return string(jsonBytes), nil
}

// Model 测试分类ORM模型
func (s *sSysTestFinance) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TestFinance.Ctx(ctx), option...)
}

// List 获取测试分类列表
func (s *sSysTestFinance) List(ctx context.Context, in *sysin.TestFinanceListInp) (list []*sysin.TestFinanceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.TestFinanceListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.TestFinance.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.TestFinance.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestFinance.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.TestFinance.Columns().Sort).OrderDesc(dao.TestFinance.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试分类列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出测试分类
func (s *sSysTestFinance) Export(ctx context.Context, in *sysin.TestFinanceListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.TestFinanceExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出测试分类-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.TestFinanceExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增测试分类
func (s *sSysTestFinance) Edit(ctx context.Context, in *sysin.TestFinanceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.TestFinanceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试分类失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.TestFinanceInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试分类失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试分类
func (s *sSysTestFinance) Delete(ctx context.Context, in *sysin.TestFinanceDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除测试分类失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取测试分类最大排序
func (s *sSysTestFinance) MaxSort(ctx context.Context, in *sysin.TestFinanceMaxSortInp) (res *sysin.TestFinanceMaxSortModel, err error) {
	if err = dao.TestFinance.Ctx(ctx).Fields(dao.TestFinance.Columns().Sort).OrderDesc(dao.TestFinance.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.TestFinanceMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取测试分类指定信息
func (s *sSysTestFinance) View(ctx context.Context, in *sysin.TestFinanceViewInp) (res *sysin.TestFinanceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类信息，请稍后重试！")
		return
	}
	return
}

// Status 更新测试分类状态
func (s *sSysTestFinance) Status(ctx context.Context, in *sysin.TestFinanceStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestFinance.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新测试分类状态失败，请稍后重试！")
		return
	}
	return
}

// Start 更新测试分类状态
func (s *sSysTestFinance) Start(ctx context.Context) (err error) {
	service.SysFinanceRsi().Rsi(ctx, "sh603259", 1200)
	service.SysFinanceMacd().MacdV2(ctx, "sh603259", 1200)
	return
	//codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	//if err != nil {
	//	return
	//}
	//wg := sync.WaitGroup{}
	//// 创建大小为5的并发限制通道
	//concurrencyLimit := 5
	//semaphore := make(chan struct{}, concurrencyLimit)
	//for _, financeCode := range codeList {
	//	wg.Add(1)
	//	semaphore <- struct{}{}
	//	simple.SafeGo(gctx.New(), func(ctx context.Context) {
	//		defer wg.Done()
	//		defer func() {
	//			// 释放信号量
	//			<-semaphore
	//		}()
	//		klineList, _ := service.SysFinanceCode().GetCodeKline(ctx, financeCode.CompleteCode, 0)
	//		// 方法 1 推荐
	//		trendPlotList := FindZigZagPointsWithKLine(klineList, 2)
	//		var trendPlotListV2 []*entity.FinanceKline
	//		for _, trendPlot := range trendPlotList {
	//			trendPlotListV2 = append(trendPlotListV2, trendPlot.KLine)
	//		}
	//		//err = service.SysFinancePlot().TrendPlot(ctx, klineList, trendPlotListV2)
	//		fmt.Println(err)
	//	})
	//}
	//wg.Wait()
	//
	//return
}

// // 如果你希望将所有波峰和波谷合并到一个列表中，可以使用这个函数
func FindAllTurningPoints(klines []*entity.FinanceKline, window int) []*entity.FinanceKline {
	allPoints := make([]*entity.FinanceKline, 0)

	n := len(klines)
	if n < window*2 {
		return allPoints
	}

	for i := window; i < n-window; i++ {
		// 检查是否为波峰
		isPeak := true
		for j := i - window; j <= i+window; j++ {
			if j == i {
				continue
			}
			if klines[j].ClosePrice >= klines[i].ClosePrice {
				isPeak = false
				break
			}
		}

		// 检查是否为波谷
		isTrough := true
		for j := i - window; j <= i+window; j++ {
			if j == i {
				continue
			}
			if klines[j].ClosePrice <= klines[i].ClosePrice {
				isTrough = false
				break
			}
		}

		if isPeak || isTrough {
			allPoints = append(allPoints, klines[i])
		}
	}

	return allPoints
}

// // 法2
//
// // 修改后的ZigZag算法，返回包含KLine的转折点
type ZigZagPointWithKLine struct {
	KLine    *entity.FinanceKline
	Index    int
	IsPeak   bool
	Strength float64 // 波动强度
}

// 返回包含KLine的转折点列表
func FindZigZagPointsWithKLine(klines []*entity.FinanceKline, minChangePercent float64) []ZigZagPointWithKLine {
	if len(klines) < 3 {
		return nil
	}

	var points []ZigZagPointWithKLine
	lastExtreme := 0 // 上一个极值点索引
	lastIsPeak := klines[1].ClosePrice > klines[0].ClosePrice

	for i := 1; i < len(klines)-1; i++ {
		current := klines[i].ClosePrice
		prev := klines[i-1].ClosePrice
		next := klines[i+1].ClosePrice

		// 检测波峰
		if current > prev && current > next {
			change := math.Abs((current - klines[lastExtreme].ClosePrice) / klines[lastExtreme].ClosePrice * 100)

			if change >= minChangePercent {
				points = append(points, ZigZagPointWithKLine{
					KLine:    klines[i],
					Index:    i,
					IsPeak:   true,
					Strength: change,
				})
				lastExtreme = i
				lastIsPeak = true
			}
		}

		// 检测波谷
		if current < prev && current < next {
			change := math.Abs((current - klines[lastExtreme].ClosePrice) / klines[lastExtreme].ClosePrice * 100)

			if change >= minChangePercent {
				points = append(points, ZigZagPointWithKLine{
					KLine:    klines[i],
					Index:    i,
					IsPeak:   false,
					Strength: change,
				})
				lastExtreme = i
				lastIsPeak = false
			}
		}
	}
	fmt.Println(lastIsPeak)
	return points
}

// 如果你只需要返回KLine列表（不包含其他信息）
func FindZigZagKLinePoints(klines []*entity.FinanceKline, minChangePercent float64) []*entity.FinanceKline {
	zigzagPoints := FindZigZagPointsWithKLine(klines, minChangePercent)
	klinePoints := make([]*entity.FinanceKline, len(zigzagPoints))

	for i, point := range zigzagPoints {
		klinePoints[i] = point.KLine
	}

	return klinePoints
}

// 如果你想要分别获取波峰和波谷的KLine列表
func FindZigZagPeaksAndTroughs(klines []*entity.FinanceKline, minChangePercent float64) ([]*entity.FinanceKline, []*entity.FinanceKline) {
	zigzagPoints := FindZigZagPointsWithKLine(klines, minChangePercent)
	peaks := make([]*entity.FinanceKline, 0)
	troughs := make([]*entity.FinanceKline, 0)

	for _, point := range zigzagPoints {
		if point.IsPeak {
			peaks = append(peaks, point.KLine)
		} else {
			troughs = append(troughs, point.KLine)
		}
	}

	return peaks, troughs
}

// 增强版本：修复原算法中的问题并添加更多功能
func FindZigZagPointsEnhanced(klines []*entity.FinanceKline, minChangePercent float64) []ZigZagPointWithKLine {
	if len(klines) < 3 {
		return nil
	}

	var points []ZigZagPointWithKLine

	// 首先找到第一个极值点
	lastExtreme := findFirstExtreme(klines)
	if lastExtreme < 0 {
		return points
	}

	// 添加第一个极值点
	firstPoint := createZigZagPoint(klines, lastExtreme)
	points = append(points, firstPoint)

	// 继续寻找后续的转折点
	for i := lastExtreme + 1; i < len(klines)-1; i++ {
		current := klines[i].ClosePrice
		prev := klines[i-1].ClosePrice
		next := klines[i+1].ClosePrice

		isPeak := current > prev && current > next
		isTrough := current < prev && current < next

		if isPeak || isTrough {
			// 计算与上一个转折点的变化幅度
			lastPoint := points[len(points)-1]
			change := math.Abs((current - lastPoint.KLine.ClosePrice) / lastPoint.KLine.ClosePrice * 100)

			if change >= minChangePercent {
				newPoint := createZigZagPoint(klines, i)
				points = append(points, newPoint)
			}
		}
	}

	return points
}

// 辅助函数：找到第一个极值点
func findFirstExtreme(klines []*entity.FinanceKline) int {
	for i := 1; i < len(klines)-1; i++ {
		current := klines[i].ClosePrice
		prev := klines[i-1].ClosePrice
		next := klines[i+1].ClosePrice

		if (current > prev && current > next) || (current < prev && current < next) {
			return i
		}
	}
	return -1
}

// 辅助函数：创建ZigZag点
func createZigZagPoint(klines []*entity.FinanceKline, index int) ZigZagPointWithKLine {
	current := klines[index].ClosePrice
	prev := klines[index-1].ClosePrice
	next := klines[index+1].ClosePrice

	isPeak := current > prev && current > next

	return ZigZagPointWithKLine{
		KLine:  klines[index],
		Index:  index,
		IsPeak: isPeak,
	}
}

// 方法 3
// 计算移动平均线
func CalculateMA(klines []*entity.FinanceKline, period int) []float64 {
	ma := make([]float64, len(klines))

	for i := period - 1; i < len(klines); i++ {
		sum := 0.0
		for j := i - period + 1; j <= i; j++ {
			sum += klines[j].ClosePrice
		}
		ma[i] = sum / float64(period)
	}

	return ma
}

// 基于MA的趋势转折点检测
type TrendPoint struct {
	Kline      *entity.FinanceKline
	Price      float64
	TrendType  string // "up", "down", "consolidation"
	Confidence float64
}

func FindTrendTurningPoints(klines []*entity.FinanceKline, shortPeriod, longPeriod int) []TrendPoint {
	shortMA := CalculateMA(klines, shortPeriod)
	longMA := CalculateMA(klines, longPeriod)

	var points []TrendPoint

	for i := longPeriod; i < len(klines); i++ {
		// 金叉：短期MA上穿长期MA
		if shortMA[i] > longMA[i] && shortMA[i-1] <= longMA[i-1] {
			confidence := math.Abs((shortMA[i]-longMA[i])/longMA[i]) * 100

			points = append(points, TrendPoint{
				Kline:      klines[i],
				Price:      klines[i].ClosePrice,
				TrendType:  "up",
				Confidence: confidence,
			})
		}

		// 死叉：短期MA下穿长期MA
		if shortMA[i] < longMA[i] && shortMA[i-1] >= longMA[i-1] {
			confidence := math.Abs((longMA[i]-shortMA[i])/longMA[i]) * 100

			points = append(points, TrendPoint{
				Kline:      klines[i],
				Price:      klines[i].ClosePrice,
				TrendType:  "down",
				Confidence: confidence,
			})
		}
	}

	return points
}

// 方法4
type WaveAnalysis struct {
	Waves         []Wave
	TurningPoints []TurningPoint
	Pattern       string // "W", "M", "Uptrend", "Downtrend"
}

type Wave struct {
	StartIndex int
	EndIndex   int
	Type       string // "up", "down"
	Height     float64
	Duration   int
}

type TurningPoint struct {
	Kline      *entity.FinanceKline
	Index      int
	Price      float64
	Type       string // "peak", "trough"
	Confidence float64
}

func AnalyzeWavePattern(klines []*entity.FinanceKline) WaveAnalysis {
	analysis := WaveAnalysis{}

	// 1. 找到所有转折点
	zigzagPoints := FindZigZagPoints(klines, 2.0) // 最小2%波动

	// 2. 识别波浪
	if len(zigzagPoints) >= 3 {
		for i := 0; i < len(zigzagPoints)-1; i++ {
			current := zigzagPoints[i]
			next := zigzagPoints[i+1]

			wave := Wave{
				StartIndex: current.Index,
				EndIndex:   next.Index,
				Height:     math.Abs(next.Price - current.Price),
				Duration:   next.Index - current.Index,
			}

			if current.IsPeak && !next.IsPeak {
				wave.Type = "down"
			} else if !current.IsPeak && next.IsPeak {
				wave.Type = "up"
			}

			analysis.Waves = append(analysis.Waves, wave)
		}

		// 3. 识别模式（如W底）
		analysis.Pattern = identifyPattern(analysis.Waves, klines)
	}

	return analysis
}

func identifyPattern(waves []Wave, klines []*entity.FinanceKline) string {
	if len(waves) < 4 {
		return "unknown"
	}

	// 简单的W底模式识别
	for i := 0; i < len(waves)-3; i++ {
		// 检查是否形成下跌-上涨-下跌-上涨的波浪序列
		if waves[i].Type == "down" && waves[i+1].Type == "up" &&
			waves[i+2].Type == "down" && waves[i+3].Type == "up" {

			// 检查第二个低点是否高于或等于第一个低点（W底特征）
			firstLow := math.Min(klines[waves[i].StartIndex].ClosePrice, klines[waves[i].EndIndex].ClosePrice)
			secondLow := math.Min(klines[waves[i+2].StartIndex].ClosePrice, klines[waves[i+2].EndIndex].ClosePrice)

			if secondLow >= firstLow {
				return "W_bottom"
			}
		}
	}

	return "complex_wave"
}

func FindZigZagPoints(klines []*entity.FinanceKline, minChangePercent float64) []ZigZagPoint {
	if len(klines) < 3 {
		return nil
	}

	var points []ZigZagPoint
	lastExtreme := 0 // 上一个极值点索引
	lastIsPeak := klines[1].ClosePrice > klines[0].ClosePrice

	for i := 1; i < len(klines)-1; i++ {
		current := klines[i].ClosePrice
		prev := klines[i-1].ClosePrice
		next := klines[i+1].ClosePrice

		// 检测波峰
		if current > prev && current > next {
			change := math.Abs((current - klines[lastExtreme].ClosePrice) / klines[lastExtreme].ClosePrice * 100)

			if change >= minChangePercent {
				points = append(points, ZigZagPoint{
					Index:    i,
					Kline:    klines[i],
					Price:    current,
					IsPeak:   true,
					Strength: change,
				})
				lastExtreme = i
				lastIsPeak = true
			}
		}

		// 检测波谷
		if current < prev && current < next {
			change := math.Abs((current - klines[lastExtreme].ClosePrice) / klines[lastExtreme].ClosePrice * 100)

			if change >= minChangePercent {
				points = append(points, ZigZagPoint{
					Kline:    klines[i],
					Index:    i,
					Price:    current,
					IsPeak:   false,
					Strength: change,
				})
				lastExtreme = i
				lastIsPeak = false
			}
		}
	}
	fmt.Println(lastIsPeak)
	return points
}

type ZigZagPoint struct {
	Index    int
	Kline    *entity.FinanceKline
	Price    float64
	IsPeak   bool
	Strength float64 // 波动强度
}
