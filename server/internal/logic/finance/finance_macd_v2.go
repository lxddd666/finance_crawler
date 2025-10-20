package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/utility/format"
)

// MACDConfig MACD配置参数
type MACDConfig struct {
	FastLength   int
	SlowLength   int
	SignalLength int
	SmaSource    string // "SMA" 或 "EMA"
	SmaSignal    string // "SMA" 或 "EMA"
}

// MACDData MACD计算结果
type MACDData struct {
	MACD   float64
	Signal float64
	Hist   float64
	FastMA float64
	SlowMA float64
}

// MACDCalculator MACD计算器
type MACDCalculator struct {
	config     MACDConfig
	fastEMA    *EMA
	slowEMA    *EMA
	signalEMA  *EMA
	fastSMA    *SMA
	slowSMA    *SMA
	signalSMA  *SMA
	macdValues []float64
}

// NewMACDCalculator 创建新的MACD计算器
func NewMACDCalculator(config MACDConfig) *MACDCalculator {
	return &MACDCalculator{
		config:     config,
		fastEMA:    NewEMA(config.FastLength),
		slowEMA:    NewEMA(config.SlowLength),
		signalEMA:  NewEMA(config.SignalLength),
		fastSMA:    NewSMA(config.FastLength),
		slowSMA:    NewSMA(config.SlowLength),
		signalSMA:  NewSMA(config.SignalLength),
		macdValues: make([]float64, 0),
	}
}

// Update 更新MACD计算（单根K线）
func (m *MACDCalculator) Update(price entity.FinanceKline) MACDData {
	var fastMA, slowMA, macd, signal float64

	// 计算快慢线
	if m.config.SmaSource == "SMA" {
		fastMA = m.fastSMA.Update(price.ClosePrice)
		slowMA = m.slowSMA.Update(price.ClosePrice)
	} else {
		fastMA = m.fastEMA.Update(price.ClosePrice)
		slowMA = m.slowEMA.Update(price.ClosePrice)
	}

	// 计算MACD
	macd = fastMA - slowMA

	// 保存MACD值用于信号线计算
	m.macdValues = append(m.macdValues, macd)
	if len(m.macdValues) > 1000 { // 限制历史数据长度
		m.macdValues = m.macdValues[1:]
	}

	// 计算信号线
	if len(m.macdValues) >= m.config.SignalLength {
		if m.config.SmaSignal == "SMA" {
			signal = m.signalSMA.Update(macd)
		} else {
			signal = m.signalEMA.Update(macd)
		}
	}

	// 计算柱状图
	hist := macd - signal

	return MACDData{
		MACD:   macd,
		Signal: signal,
		Hist:   hist,
		FastMA: fastMA,
		SlowMA: slowMA,
	}
}

// IsReady 检查MACD是否已准备好（有足够数据）
func (m *MACDCalculator) IsReady() bool {
	return len(m.macdValues) >= m.config.SlowLength
}

// EMA 指数移动平均计算器
type EMA struct {
	length  int
	alpha   float64
	current float64
	isReady bool
}

func NewEMA(length int) *EMA {
	return &EMA{
		length:  length,
		alpha:   2.0 / (float64(length) + 1.0),
		current: 0,
		isReady: false,
	}
}

func (e *EMA) Update(price float64) float64 {
	if !e.isReady {
		e.current = price
		e.isReady = true
		return e.current
	}

	e.current = (price * e.alpha) + (e.current * (1 - e.alpha))
	return e.current
}

// SMA 简单移动平均计算器
type SMA struct {
	length int
	prices []float64
	sum    float64
}

func NewSMA(length int) *SMA {
	return &SMA{
		length: length,
		prices: make([]float64, 0),
		sum:    0,
	}
}

func (s *SMA) Update(price float64) float64 {
	s.prices = append(s.prices, price)
	s.sum += price

	if len(s.prices) > s.length {
		// 移除最旧的价格
		oldest := s.prices[0]
		s.prices = s.prices[1:]
		s.sum -= oldest
	}

	if len(s.prices) == 0 {
		return 0
	}

	return s.sum / float64(len(s.prices))
}

// 辅助函数：检测MACD柱状图状态变化
type HistogramState int

const (
	Rising HistogramState = iota
	Falling
	Unchanged
)

// GetHistogramStateChange 检测柱状图状态变化
func GetHistogramStateChange(prevHist, currHist float64) HistogramState {
	if prevHist >= 0 && currHist < 0 {
		return Falling // 从上升转为下降
	} else if prevHist <= 0 && currHist > 0 {
		return Rising // 从下降转为上升
	}
	return Unchanged
}

// v
func (s *sSysFinanceMacd) MacdV2(ctx context.Context, code string, scale int) (err error) {
	// 创建MACD计算器（使用默认参数）
	config := MACDConfig{
		FastLength:   12,
		SlowLength:   26,
		SignalLength: 9,
		SmaSource:    "EMA",
		SmaSignal:    "EMA",
	}

	macd := NewMACDCalculator(config)

	// 模拟K线数据
	var financeList []entity.FinanceKline
	err = dao.FinanceKline.Ctx(ctx).Where(dao.FinanceKline.Columns().Code, code).Where(dao.FinanceKline.Columns().Scale, scale).Scan(&financeList)
	if err != nil {
		return
	}

	var prevHist float64
	var macdList []*entity.FinanceMacd
	for i, price := range financeList {
		data := macd.Update(price)

		if i > 0 {
			state := GetHistogramStateChange(prevHist, data.Hist)
			switch state {
			case Rising:
				println("MACD柱状图从下降转为上升")
			case Falling:
				println("MACD柱状图从上升转为下降")
			}
		}

		prevHist = data.Hist

		fmt.Println(fmt.Sprintf("价格: %.2f, MACD: %.4f, 信号: %.4f, 柱状图: %.4f 时间：%s\n",
			price.ClosePrice, data.MACD, data.Signal, data.Hist, price.Day))

		macdRes := &entity.FinanceMacd{
			Code:         price.Code,
			FastPeriod:   12,
			SlowPeriod:   26,
			SignalPeriod: 9,
			Dif:          data.FastMA,
			Dea:          data.SlowMA,
			Macd:         data.Hist,
			Timestamp:    format.DayStrToTimestamp(price.Day),
			ClosePrice:   price.ClosePrice,
			Day:          price.Day,
			Key:          fmt.Sprintf("%s%s:%d", price.Code, price.Day, scale),
			Scale:        consts.ScaleFiveDay,
		}
		macdList = append(macdList, macdRes)
	}
	_, err = dao.FinanceMacd.Ctx(ctx).InsertIgnore(macdList)
	return
}
