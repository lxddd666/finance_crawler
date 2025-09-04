package sys

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/logic/alltick"

	"hotgo/internal/model/entity"
	"math"
)

// BollResult 布林带计算结果
type BollResult struct {
	Timestamp         int64   `json:"timestamp"`
	MiddleBand        float64 // 中轨（移动平均线）
	UpperBand         float64 // 上轨（移动平均线 + 两倍标准差）
	LowerBand         float64 // 下轨（移动平均线 - 两倍标准差）
	StandardDeviation float64 // 标准差
	ClosePrice        float64 // 当前时间
}

// Bool
func (s *sSysStockIndicator) Boll(ctx context.Context, code string, klineType, klineNum, multiple int) (result *BollResult, err error) {
	klineList, err := s.bollReq(ctx, code, klineType, klineNum)
	if err != nil {
		return
	}
	result, lastKline, err := s.calculateBoll(klineList, multiple)
	if err != nil {
		return
	}
	id, err := dao.FinanceKline.Ctx(ctx).InsertAndGetId(lastKline)
	if err != nil {
		return
	}
	// 插入数据库
	_, _ = dao.FinanceBoll.Ctx(ctx).InsertIgnore(entity.FinanceBoll{
		Code:              code,
		KlineId:           id,
		KlineType:         klineType,
		KlineNum:          klineNum,
		Multiple:          multiple,
		Timestamp:         result.Timestamp,
		MiddleBand:        result.MiddleBand,
		UpperBand:         result.UpperBand,
		LowerBand:         result.LowerBand,
		StandardDeviation: result.StandardDeviation,
		ClosePrice:        result.ClosePrice,
	})
	return
}

func (s *sSysStockIndicator) calculateBoll(data []*entity.StockKlineData, multiple int) (result *BollResult, lastKline *entity.StockKlineData, err error) {
	if len(data) == 0 {
		return
	}
	if multiple == 0 {
		multiple = 2
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
	multipleStandardDeviations := float64(multiple) * standardDeviation

	// 构造布林带结果
	result = &BollResult{
		MiddleBand:        sma,
		UpperBand:         sma + multipleStandardDeviations,
		LowerBand:         sma - multipleStandardDeviations,
		StandardDeviation: standardDeviation,
		ClosePrice:        data[len(data)-1].ClosePrice,
		Timestamp:         data[0].Timestamp,
	}
	lastKline = data[0]
	return
}

// bollReq 请求
func (s *sSysStockIndicator) bollReq(ctx context.Context, code string, klineType, klineNum int) (respData []*entity.StockKlineData, err error) {
	////Code:              code,
	////		KlineType:         8,
	////		KlineTimestampEnd: 0,
	////		QueryKlineNum:     20,
	////		AdjustType:        0,
	//
	response, err := alltick.GetKlineData(ctx, &entity.FinanceAlltickRequest{
		Code:              code,
		KlineType:         klineType,
		KlineTimestampEnd: 0,
		QueryKlineNum:     klineNum,
		AdjustType:        0,
	})
	if err != nil {
		return
	}

	respData = make([]*entity.StockKlineData, 0)
	for _, kline := range response.Data.KlineList {
		respData = append(respData, &entity.StockKlineData{
			Code:       response.Data.Code,
			Timestamp:  gconv.Int64(kline.Timestamp),
			OpenPrice:  gconv.Float64(kline.OpenPrice),
			ClosePrice: gconv.Float64(kline.ClosePrice),
			HighPrice:  gconv.Float64(kline.HighPrice),
			LowPrice:   gconv.Float64(kline.LowPrice),
			Volume:     gconv.Int64(kline.Volume),
			Turnover:   gconv.Float64(kline.Turnover),
		})
	}
	if len(respData) == 0 {
		err = gerror.New("获取数据为空")
		return
	}
	return
}

// SimpleMovingAverage 简单移动平均线
func SimpleMovingAverage(data []*entity.StockKlineData) float64 {
	var total float64
	for _, kline := range data {
		total += kline.ClosePrice
	}
	return total / float64(len(data))
}
