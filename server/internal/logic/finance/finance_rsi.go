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
	"github.com/gogf/gf/v2/util/gconv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility"
	"log"
	"math"
	"os"

	"github.com/gogf/gf/v2/database/gdb"
)

type sSysFinanceRsi struct{}

func NewSysFinanceRsi() *sSysFinanceRsi {
	return &sSysFinanceRsi{}
}

func init() {
	service.RegisterSysFinanceRsi(NewSysFinanceRsi())
}

// Model rsi线ORM模型
func (s *sSysFinanceRsi) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceRsi.Ctx(ctx), option...)
}

// RSIConfig RSI配置参数
type RSIConfig struct {
	Length              int
	Source              string // "close", "open", "high", "low"
	CalculateDivergence bool
	MAType              string // "None", "SMA", "SMA + Bollinger Bands", "EMA", "SMMA (RMA)", "WMA", "VWMA"
	MALength            int
	BBMultiplier        float64
	LookbackRight       int
	LookbackLeft        int
	RangeUpper          int
	RangeLower          int
}

// RSIData RSI计算结果
type RSIData struct {
	RSI            float64
	SmoothingMA    float64
	BBUpper        float64
	BBLower        float64
	IsBB           bool
	BullDivergence bool
	BearDivergence bool
	RegularBullish []bool
	RegularBearish []bool
	Day            string
}

// PriceData 价格数据
type PriceData struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
	Day    string
}

// RSICalculator RSI计算器
type RSICalculator struct {
	config          RSIConfig
	prices          []PriceData
	rsiValues       []float64
	upChanges       []float64
	downChanges     []float64
	upRMA           *RMA
	downRMA         *RMA
	maCalculator    *MACalculator
	stdevCalculator *StdevCalculator
}

// NewRSICalculator 创建RSI计算器
func NewRSICalculator(config RSIConfig) *RSICalculator {
	return &RSICalculator{
		config:          config,
		prices:          make([]PriceData, 0),
		rsiValues:       make([]float64, 0),
		upChanges:       make([]float64, 0),
		downChanges:     make([]float64, 0),
		upRMA:           NewRMA(config.Length),
		downRMA:         NewRMA(config.Length),
		maCalculator:    NewMACalculator(config.MAType, config.MALength),
		stdevCalculator: NewStdevCalculator(config.MALength),
	}
}

// Update 更新RSI计算
func (r *RSICalculator) Update(price PriceData) RSIData {
	r.prices = append(r.prices, price)

	// 获取源数据
	var source float64
	switch r.config.Source {
	case "open":
		source = price.Open
	case "high":
		source = price.High
	case "low":
		source = price.Low
	default:
		source = price.Close
	}

	// 计算价格变化
	var change float64
	if len(r.prices) > 1 {
		prevSource := getSource(r.prices[len(r.prices)-2], r.config.Source)
		change = source - prevSource
	}

	// 计算上涨和下跌变化
	upChange := math.Max(change, 0)
	downChange := math.Max(-change, 0)

	// 计算RMA（Wilder's Smoothing Method）
	upRMA := r.upRMA.Update(upChange)
	downRMA := r.downRMA.Update(downChange)

	// 计算RSI
	var rsi float64
	if downRMA == 0 {
		rsi = 100
	} else if upRMA == 0 {
		rsi = 0
	} else {
		rs := upRMA / downRMA
		rsi = 100 - (100 / (1 + rs))
	}

	r.rsiValues = append(r.rsiValues, rsi)

	// 限制历史数据长度
	if len(r.prices) > 1000 {
		r.prices = r.prices[1:]
	}
	if len(r.rsiValues) > 1000 {
		r.rsiValues = r.rsiValues[1:]
	}

	// 计算平滑移动平均
	var smoothingMA, bbUpper, bbLower float64
	var isBB bool

	if r.config.MAType != "None" {
		smoothingMA = r.maCalculator.Update(rsi)

		if r.config.MAType == "SMA + Bollinger Bands" {
			isBB = true
			// 计算布林带
			stdev := r.stdevCalculator.Update(rsi)
			bbUpper = smoothingMA + (stdev * r.config.BBMultiplier)
			bbLower = smoothingMA - (stdev * r.config.BBMultiplier)
		}
	}

	// 计算背离（需要足够的历史数据）
	bullDivergence, bearDivergence := false, false
	if r.config.CalculateDivergence && len(r.rsiValues) >= r.config.LookbackRight+1 {
		bullDivergence, bearDivergence = r.calculateDivergence()
	}

	return RSIData{
		RSI:            rsi,
		SmoothingMA:    smoothingMA,
		BBUpper:        bbUpper,
		BBLower:        bbLower,
		IsBB:           isBB,
		BullDivergence: bullDivergence,
		BearDivergence: bearDivergence,
	}
}

// 获取源数据
func getSource(price PriceData, sourceType string) float64 {
	switch sourceType {
	case "open":
		return price.Open
	case "high":
		return price.High
	case "low":
		return price.Low
	default:
		return price.Close
	}
}

// RMA (Wilder's Smoothing) 计算器
type RMA struct {
	length  int
	alpha   float64
	current float64
	isReady bool
}

func NewRMA(length int) *RMA {
	return &RMA{
		length:  length,
		alpha:   1.0 / float64(length),
		current: 0,
		isReady: false,
	}
}

func (r *RMA) Update(value float64) float64 {
	if !r.isReady {
		r.current = value
		r.isReady = true
		return r.current
	}

	r.current = (value * r.alpha) + (r.current * (1 - r.alpha))
	return r.current
}

// MACalculator 移动平均计算器
type MACalculator struct {
	maType  string
	length  int
	sma     *SMA
	ema     *EMA
	rma     *RMA
	wma     *WMA
	vwma    *VWMA
	prices  []float64
	volumes []float64
}

func NewMACalculator(maType string, length int) *MACalculator {
	return &MACalculator{
		maType:  maType,
		length:  length,
		sma:     NewSMA(length),
		ema:     NewEMA(length),
		rma:     NewRMA(length),
		wma:     NewWMA(length),
		vwma:    NewVWMA(length),
		prices:  make([]float64, 0),
		volumes: make([]float64, 0),
	}
}

func (m *MACalculator) Update(value float64) float64 {
	m.prices = append(m.prices, value)

	switch m.maType {
	case "SMA", "SMA + Bollinger Bands":
		return m.sma.Update(value)
	case "EMA":
		return m.ema.Update(value)
	case "SMMA (RMA)":
		return m.rma.Update(value)
	case "WMA":
		return m.wma.Update(value)
	case "VWMA":
		// 简化版VWMA，使用固定成交量1
		return m.vwma.Update(value, 1.0)
	default:
		return value
	}
}

// WMA 加权移动平均
type WMA struct {
	length int
	prices []float64
}

func NewWMA(length int) *WMA {
	return &WMA{
		length: length,
		prices: make([]float64, 0),
	}
}

func (w *WMA) Update(price float64) float64 {
	w.prices = append(w.prices, price)
	if len(w.prices) > w.length {
		w.prices = w.prices[1:]
	}

	if len(w.prices) == 0 {
		return 0
	}

	var sum, weightSum float64
	for i, p := range w.prices {
		weight := float64(i + 1)
		sum += p * weight
		weightSum += weight
	}

	return sum / weightSum
}

// VWMA 成交量加权移动平均
type VWMA struct {
	length    int
	prices    []float64
	volumes   []float64
	priceSum  float64
	volumeSum float64
}

func NewVWMA(length int) *VWMA {
	return &VWMA{
		length:    length,
		prices:    make([]float64, 0),
		volumes:   make([]float64, 0),
		priceSum:  0,
		volumeSum: 0,
	}
}

func (v *VWMA) Update(price, volume float64) float64 {
	v.prices = append(v.prices, price)
	v.volumes = append(v.volumes, volume)
	v.priceSum += price * volume
	v.volumeSum += volume

	if len(v.prices) > v.length {
		oldestPrice := v.prices[0]
		oldestVolume := v.volumes[0]
		v.prices = v.prices[1:]
		v.volumes = v.volumes[1:]
		v.priceSum -= oldestPrice * oldestVolume
		v.volumeSum -= oldestVolume
	}

	if v.volumeSum == 0 {
		return 0
	}

	return v.priceSum / v.volumeSum
}

// StdevCalculator 标准差计算器
type StdevCalculator struct {
	length int
	prices []float64
	sma    *SMA
}

func NewStdevCalculator(length int) *StdevCalculator {
	return &StdevCalculator{
		length: length,
		prices: make([]float64, 0),
		sma:    NewSMA(length),
	}
}

func (s *StdevCalculator) Update(price float64) float64 {
	s.prices = append(s.prices, price)
	if len(s.prices) > s.length {
		s.prices = s.prices[1:]
	}

	if len(s.prices) < 2 {
		return 0
	}

	// 计算平均值
	mean := s.sma.Update(price)

	// 计算方差
	var variance float64
	for _, p := range s.prices {
		variance += math.Pow(p-mean, 2)
	}
	variance /= float64(len(s.prices))

	// 返回标准差
	return math.Sqrt(variance)
}

// 计算背离
func (r *RSICalculator) calculateDivergence() (bool, bool) {
	if len(r.rsiValues) < r.config.LookbackRight+1 {
		return false, false
	}

	// 寻找RSI的极值点
	plFound := r.findPivotLow()
	phFound := r.findPivotHigh()

	bullCond := false
	bearCond := false

	if plFound {
		// 正则牛市背离：RSI形成更高的低点，但价格形成更低的低点
		rsiHL := r.checkRSIHigherLow()
		priceLL := r.checkPriceLowerLow()
		bullCond = priceLL && rsiHL
	}

	if phFound {
		// 正则熊市背离：RSI形成更低的高点，但价格形成更高的高点
		rsiLH := r.checkRSILowerHigh()
		priceHH := r.checkPriceHigherHigh()
		bearCond = priceHH && rsiLH
	}

	return bullCond, bearCond
}

// 寻找RSI的低点枢轴
func (r *RSICalculator) findPivotLow() bool {
	if len(r.rsiValues) < r.config.LookbackLeft+r.config.LookbackRight+1 {
		return false
	}

	currentIndex := len(r.rsiValues) - 1 - r.config.LookbackRight
	currentValue := r.rsiValues[currentIndex]

	// 检查左侧
	for i := 1; i <= r.config.LookbackLeft; i++ {
		if currentIndex-i < 0 {
			break
		}
		if r.rsiValues[currentIndex-i] <= currentValue {
			return false
		}
	}

	// 检查右侧
	for i := 1; i <= r.config.LookbackRight; i++ {
		if currentIndex+i >= len(r.rsiValues) {
			break
		}
		if r.rsiValues[currentIndex+i] <= currentValue {
			return false
		}
	}

	return true
}

// 寻找RSI的高点枢轴
func (r *RSICalculator) findPivotHigh() bool {
	if len(r.rsiValues) < r.config.LookbackLeft+r.config.LookbackRight+1 {
		return false
	}

	currentIndex := len(r.rsiValues) - 1 - r.config.LookbackRight
	currentValue := r.rsiValues[currentIndex]

	// 检查左侧
	for i := 1; i <= r.config.LookbackLeft; i++ {
		if currentIndex-i < 0 {
			break
		}
		if r.rsiValues[currentIndex-i] >= currentValue {
			return false
		}
	}

	// 检查右侧
	for i := 1; i <= r.config.LookbackRight; i++ {
		if currentIndex+i >= len(r.rsiValues) {
			break
		}
		if r.rsiValues[currentIndex+i] >= currentValue {
			return false
		}
	}

	return true
}

// 检查RSI是否形成更高的低点
func (r *RSICalculator) checkRSIHigherLow() bool {
	// 简化实现，实际需要比较当前低点和前一个低点
	currentIndex := len(r.rsiValues) - 1 - r.config.LookbackRight
	if currentIndex < 1 {
		return false
	}

	currentLow := r.rsiValues[currentIndex]
	// 这里需要找到前一个有效的低点进行比较
	// 简化：假设当前值比前一个值高
	return currentLow > r.rsiValues[currentIndex-1]
}

// 检查价格是否形成更低的低点
func (r *RSICalculator) checkPriceLowerLow() bool {
	currentIndex := len(r.prices) - 1 - r.config.LookbackRight
	if currentIndex < 1 {
		return false
	}

	currentLow := r.prices[currentIndex].Low
	// 简化：假设当前值比前一个值低
	return currentLow < r.prices[currentIndex-1].Low
}

// 检查RSI是否形成更低的高点
func (r *RSICalculator) checkRSILowerHigh() bool {
	currentIndex := len(r.rsiValues) - 1 - r.config.LookbackRight
	if currentIndex < 1 {
		return false
	}

	currentHigh := r.rsiValues[currentIndex]
	// 简化：假设当前值比前一个值低
	return currentHigh < r.rsiValues[currentIndex-1]
}

// 检查价格是否形成更高的高点
func (r *RSICalculator) checkPriceHigherHigh() bool {
	currentIndex := len(r.prices) - 1 - r.config.LookbackRight
	if currentIndex < 1 {
		return false
	}

	currentHigh := r.prices[currentIndex].High
	// 简化：假设当前值比前一个值高
	return currentHigh > r.prices[currentIndex-1].High
}

// 检查是否在范围内
func (r *RSICalculator) inRange(condition bool) bool {
	// 简化实现
	return condition
}

// 使用示例
func (s *sSysFinanceRsi) Rsi(ctx context.Context, code string, scale int) (err error) {
	// 配置RSI参数
	config := RSIConfig{
		Length:              14,
		Source:              "close",
		CalculateDivergence: true,
		MAType:              "SMA",
		MALength:            14,
		BBMultiplier:        2.0,
		LookbackRight:       5,
		LookbackLeft:        5,
		RangeUpper:          60,
		RangeLower:          5,
	}

	// 创建RSI计算器
	rsi := NewRSICalculator(config)

	// 模拟价格数据
	var financeList []PriceData
	err = dao.FinanceKline.Ctx(ctx).Fields("open_price as open, close_price as close, high_price as high, low_price as low, Volume, day").Where(dao.FinanceKline.Columns().Code, code).Where(dao.FinanceKline.Columns().Scale, scale).Scan(&financeList)
	if err != nil {
		return
	}

	// 计算RSI
	var list []RSIData
	var rsiList []*entity.FinanceRsi
	for _, price := range financeList {
		result := rsi.Update(price)

		fmt.Printf("Close: %.2f, RSI: %.2f, 时间: %s", price.Close, result.RSI, price.Day)
		if result.SmoothingMA > 0 {
			fmt.Printf(", SmoothingMA: %.2f", result.SmoothingMA)
		}
		if result.IsBB {
			fmt.Printf(", BB Upper: %.2f, BB Lower: %.2f", result.BBUpper, result.BBLower)
		}
		if result.BullDivergence {
			fmt.Printf(" <- 牛市背离!")
		}
		if result.BearDivergence {
			fmt.Printf(" <- 熊市背离!")
		}
		fmt.Println()
		result.Day = price.Day
		list = append(list, result)
		rsiList = append(rsiList, &entity.FinanceRsi{
			Code:           code,
			Rsi:            result.RSI,
			SmoothingMa:    result.SmoothingMA,
			BollUpper:      gconv.Int(result.BBUpper),
			BollLower:      gconv.Int(result.BBLower),
			IsBoll:         utility.BoolTrueInt(result.IsBB),
			BullDivergence: utility.BoolTrueInt(result.BullDivergence),
			BearDivergence: utility.BoolTrueInt(result.BearDivergence),
			MaLength:       14,
			Key:            fmt.Sprintf("%s%s:%d", code, price.Day, scale),
			Day:            price.Day,
			Scale:          scale,
		})
	}

	_, err = s.Model(ctx).InsertIgnore(rsiList)

	return
}

func createPlot(ctx context.Context, data []RSIData, code string) {

	p := plot.New()
	p.Title.Text = "RSI 分析图"
	p.X.Label.Text = "日期"
	p.Y.Label.Text = "数值"

	// 使用简单的索引作为x轴
	rsiPoints := make(plotter.XYs, len(data))
	smaPoints := make(plotter.XYs, len(data))

	for i := range data {
		rsiPoints[i].X = float64(i)
		rsiPoints[i].Y = data[i].RSI

		smaPoints[i].X = float64(i)
		smaPoints[i].Y = data[i].SmoothingMA
	}

	// 创建线条
	rsiLine, err := plotter.NewLine(rsiPoints)
	if err != nil {
		log.Fatal(err)
	}
	rsiLine.Color = plotutil.Color(0)

	smaLine, err := plotter.NewLine(smaPoints)
	if err != nil {
		log.Fatal(err)
	}
	smaLine.Color = plotutil.Color(1)

	p.Add(rsiLine, smaLine)
	p.Legend.Add("RSI", rsiLine)
	p.Legend.Add("SmoothingMA", smaLine)

	// 设置x轴标签
	p.NominalX(data[0].Day, data[1].Day, data[2].Day, data[3].Day, data[4].Day)

	// 保存为PNG文件
	saveDir := "D:\\go\\src\\finance_crawlerV2\\hotgo\\server\\internal\\logic\\finance\\image\\"
	savePath := saveDir + fmt.Sprintf("%s_finance_trend.png", code)
	if _, err = os.Stat(saveDir); os.IsNotExist(err) {
		if err = os.MkdirAll(saveDir, 0755); err != nil {
			return
		}
		log.Printf("创建文件夹: %s", saveDir)
	}
	if err = p.Save(10*vg.Inch, 6*vg.Inch, savePath); err != nil {
		return
	}
	fmt.Println("图表已生成: " + savePath)
}
