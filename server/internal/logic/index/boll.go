package index

import (
	"context"
	"hotgo/internal/model/entity"
	"math"
)

// BollResult 布林带计算结果
type BollResult struct {
	MiddleBand        float64 // 中轨（移动平均线）
	UpperBand         float64 // 上轨（移动平均线 + 两倍标准差）
	LowerBand         float64 // 下轨（移动平均线 - 两倍标准差）
	StandardDeviation float64 // 标准差
	NowPrice          float64 // 当前时间
}

func Boll(ctx context.Context, data []entity.StockKlineData) (result *BollResult, err error) {
	if len(data) == 0 {
		return nil, nil
	}

	// 移动平均线
	sma := SimpleMovingAverage(data)

	// 计算标准差
	var sumSquaredDeviations float64
	for _, kline := range data {
		deviation := kline.ClosePrice - sma
		sumSquaredDeviations += deviation * deviation
	}

	variance := sumSquaredDeviations / float64(len(data))
	standardDeviation := math.Sqrt(variance)

	// 两倍标准差
	twoStandardDeviations := 2 * standardDeviation

	// 构造布林带结果
	result = &BollResult{
		MiddleBand:        sma,
		UpperBand:         sma + twoStandardDeviations,
		LowerBand:         sma - twoStandardDeviations,
		StandardDeviation: standardDeviation,
		NowPrice:          data[len(data)-1].ClosePrice,
	}

	return result, nil
}

// SimpleMovingAverage 简单移动平均线
func SimpleMovingAverage(data []entity.StockKlineData) float64 {
	var total float64
	for _, kline := range data {
		total += kline.ClosePrice
	}
	return total / float64(len(data))
}
