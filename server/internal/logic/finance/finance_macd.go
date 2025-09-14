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
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/format"

	"github.com/gogf/gf/v2/database/gdb"
)

type sSysFinanceMacd struct{}

func NewSysFinanceMacd() *sSysFinanceMacd {
	return &sSysFinanceMacd{}
}

func init() {
	service.RegisterSysFinanceMacd(NewSysFinanceMacd())
}

// Model macd线ORM模型
func (s *sSysFinanceMacd) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceMacd.Ctx(ctx), option...)
}

// Macd Macd计算
func (s *sSysFinanceMacd) Macd(ctx context.Context, data []*entity.FinanceKline, slowPeriod int, fastPeriod int, signalPeriod int) (results []*entity.FinanceMacd) {
	if len(data) < fastPeriod {
		return nil
	}
	klineNum := len(data)
	codeInfo := data[0]
	// 设置默认参数
	if slowPeriod == 0 {
		slowPeriod = 12
	}
	if fastPeriod == 0 {
		fastPeriod = 26
	}
	if signalPeriod == 0 {
		signalPeriod = 9
	}

	// 计算短期和长期EMA
	shortEMA := calculateEMA(data, slowPeriod)
	longEMA := calculateEMA(data, fastPeriod)

	// 计算DIF (差离值)
	dif := make([]float64, len(data))
	for i := 0; i < len(data); i++ {
		dif[i] = shortEMA[i] - longEMA[i]
	}

	// 将DIF转换为StockData格式，以便计算DEM
	difData := make([]*entity.FinanceKline, len(data))
	for i := 0; i < len(data); i++ {
		difData[i] = &entity.FinanceKline{
			Day:        data[i].Day,
			ClosePrice: dif[i],
		}
	}

	// 计算DEM (信号线)
	dem := calculateEMA(difData, signalPeriod)

	// 计算MACD柱状图并组装结果
	results = make([]*entity.FinanceMacd, len(data))
	for i := 0; i < len(data); i++ {
		results[i] = &entity.FinanceMacd{
			Code:         codeInfo.Code,
			FastPeriod:   fastPeriod,
			SlowPeriod:   slowPeriod,
			SignalPeriod: signalPeriod,
			Dif:          dif[i],
			Dea:          dem[i],
			Macd:         (dif[i] - dem[i]) * 2,
			Timestamp:    format.DayStrToTimestamp(codeInfo.Day),
			ClosePrice:   data[i].ClosePrice,
			KlineNum:     klineNum,
			Day:          data[i].Day,
			Key:          fmt.Sprintf("%s%s", data[i].Code, data[i].Day),
			Scale:        consts.ScaleDay,
		}
	}

	if len(results) > 0 {
		_, _ = dao.FinanceMacd.Ctx(ctx).InsertIgnore(results)
	}
	return
}

// 计算指数移动平均线 (EMA)
func calculateEMA(data []*entity.FinanceKline, period int) []float64 {
	if len(data) == 0 {
		return nil
	}

	ema := make([]float64, len(data))
	// 第一个EMA值等于第一个收盘价
	ema[0] = data[0].ClosePrice

	// 计算平滑系数
	multiplier := 2.0 / float64(period+1)

	// 计算后续EMA值
	for i := 1; i < len(data); i++ {
		ema[i] = (data[i].ClosePrice * multiplier) + (ema[i-1] * (1 - multiplier))
	}

	return ema
}

// 计算EMA
func calculateEMAV3(prices []float64, period int) []float64 {
	ema := make([]float64, len(prices))
	multiplier := 2.0 / float64(period+1)
	ema[0] = prices[0] // 初始值为第一个价格
	for i := 1; i < len(prices); i++ {
		ema[i] = (prices[i]-ema[i-1])*multiplier + ema[i-1]
	}
	return ema
}

// 计算MACD
func CalculateMACDV3(prices []float64) ([]float64, []float64, []float64) {
	shortPeriod := 12
	longPeriod := 26
	signalPeriod := 9

	shortEMA := calculateEMAV3(prices, shortPeriod)
	longEMA := calculateEMAV3(prices, longPeriod)

	macdLine := make([]float64, len(prices))
	for i := range prices {
		macdLine[i] = shortEMA[i] - longEMA[i]
	}

	signalLine := calculateEMAV3(macdLine, signalPeriod)

	histogram := make([]float64, len(prices))
	for i := range prices {
		histogram[i] = macdLine[i] - signalLine[i]
	}
	return macdLine, signalLine, histogram
}
