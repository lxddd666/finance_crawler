package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/utility/format"
)

// Macd Macd计算
func (s *sSysStockIndicator) Macd(ctx context.Context, data []*entity.FinanceKline, slowPeriod int, fastPeriod int, signalPeriod int) (results []*entity.FinanceMacd) {
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
