package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/logic/sina"
	"hotgo/internal/model/httpReq"
	"hotgo/internal/model/result"
	"hotgo/utility/format"

	"hotgo/internal/model/entity"
	"math"
)

// Boll boll带
func (s *sSysStockIndicator) Boll(ctx context.Context, code, ma string, scale, datalen, multiple int) (boll *entity.FinanceBoll, err error) {
	klineList, err := s.bollReq(ctx, code, ma, scale, datalen)
	if err != nil {
		return
	}
	result, lastKline, err := s.CalculateBoll(klineList, multiple)
	if err != nil {
		return
	}
	id, err := dao.FinanceKline.Ctx(ctx).InsertAndGetId(lastKline)
	if err != nil {
		return
	}
	// 插入数据库
	boll = &entity.FinanceBoll{
		Code:              code,
		KlineId:           id,
		Multiple:          multiple,
		Timestamp:         result.Timestamp,
		MiddleBand:        result.MiddleBand,
		UpperBand:         result.UpperBand,
		LowerBand:         result.LowerBand,
		StandardDeviation: result.StandardDeviation,
		ClosePrice:        result.ClosePrice,
		Scale:             scale,
		Key:               lastKline.Key,
	}
	_, _ = dao.FinanceBoll.Ctx(ctx).InsertIgnore(boll)
	return
}

func (s *sSysStockIndicator) CalculateBoll(data []*entity.FinanceKline, multiple int) (resp *result.BollResult, lastKline *entity.FinanceKline, err error) {
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
	resp = &result.BollResult{
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
func (s *sSysStockIndicator) bollReq(ctx context.Context, code, ma string, scale, datalen int) (respData []*entity.FinanceKline, err error) {

	KlineList, err := sina.GetKlineData(ctx, &httpReq.SinaHttpReq{
		Symbol:  code,
		Scale:   scale,
		Ma:      ma,
		Datalen: datalen,
	})

	//KlineList, err = s.Kline(ctx, code, ma, scale, datalen)
	if err != nil {
		return
	}

	respData = make([]*entity.FinanceKline, 0)
	for _, kline := range KlineList {
		respData = append(respData, &entity.FinanceKline{
			Code:       code,
			Timestamp:  format.DayStrToTimestamp(kline.Day),
			OpenPrice:  gconv.Float64(kline.Open),
			ClosePrice: gconv.Float64(kline.Close),
			HighPrice:  gconv.Float64(kline.High),
			LowPrice:   gconv.Float64(kline.Low),
			Volume:     gconv.Int64(kline.Volume),
			Key:        fmt.Sprintf("%s%d", code, format.DayStrToTimestamp(kline.Day)),
		})
	}
	if len(respData) == 0 {
		err = gerror.New("获取数据为空")
		return
	}
	return
}

// SimpleMovingAverage 简单移动平均线
func SimpleMovingAverage(data []*entity.FinanceKline) float64 {
	var total float64
	for _, kline := range data {
		total += kline.ClosePrice
	}
	return total / float64(len(data))
}
