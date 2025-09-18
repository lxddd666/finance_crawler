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
	"hotgo/internal/logic/sina"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/httpReq"
	"hotgo/internal/model/result"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysFinanceBoll struct{}

func NewSysFinanceBoll() *sSysFinanceBoll {
	return &sSysFinanceBoll{}
}

func init() {
	service.RegisterSysFinanceBoll(NewSysFinanceBoll())
}

// Model boll带ORM模型
func (s *sSysFinanceBoll) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceBoll.Ctx(ctx), option...)
}

// Boll 用股票k线计算boll
func (s *sSysFinanceBoll) Boll(ctx context.Context, data []*entity.FinanceKline, multiple int) (resp *result.BollResult, lastKline *entity.FinanceKline, err error) {
	if len(data) == 0 {
		return
	}
	if multiple == 0 {
		multiple = 2
	}

	dataInfo := data[len(data)-1]

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
	resp.Degree = normalizeDegree(resp.UpperBand, resp.LowerBand, resp.ClosePrice)
	lastKline = data[0]
	boll := &entity.FinanceBoll{
		Code:              lastKline.Code,
		Multiple:          multiple,
		Timestamp:         resp.Timestamp,
		MiddleBand:        resp.MiddleBand,
		UpperBand:         resp.UpperBand,
		LowerBand:         resp.LowerBand,
		StandardDeviation: resp.StandardDeviation,
		ClosePrice:        resp.ClosePrice,
		Degree:            resp.Degree,
		Day:               dataInfo.Day,
		Scale:             consts.ScaleDay,
		Key:               lastKline.Key,
	}
	_, err = dao.FinanceBoll.Ctx(ctx).InsertIgnore(boll)

	return
}

// bollReq 请求
func (s *sSysFinanceBoll) bollReq(ctx context.Context, code, ma string, scale, datalen int) (respData []*entity.FinanceKline, err error) {

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

func normalizeDegree(top, bottom, point float64) float64 {
	if top <= bottom {
		return 0
	}
	return (point - bottom) / (top - bottom)
}
