// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/model/entity"

	"hotgo/internal/dao"

	"hotgo/internal/library/hgorm/handler"

	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
)

type sSysFinanceDailyKline struct{}

func NewSysFinanceDailyKline() *sSysFinanceDailyKline {
	return &sSysFinanceDailyKline{}
}

func init() {
	service.RegisterSysFinanceDailyKline(NewSysFinanceDailyKline())
}

// Model 股票日K线数据表ORM模型
func (s *sSysFinanceDailyKline) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceDailyKline.Ctx(ctx), option...)
}

// MovingAverage 均线计算
func (s *sSysFinanceDailyKline) MovingAverage(ctx context.Context, klineList []*entity.FinanceKline) (err error) {
	lastKline := klineList[len(klineList)-1]

	// 获取所有均线数据
	stockCode := lastKline.Code
	// 获取股票
	// md_5
	average5 := calculateMovingAverage(klineList, 5)
	// md_10
	average10 := calculateMovingAverage(klineList, 10)
	// md_20
	average20 := calculateMovingAverage(klineList, 20)
	// md_30
	average30 := calculateMovingAverage(klineList, 30)
	// md_60
	average60 := calculateMovingAverage(klineList, 60)

	_, err = s.Model(ctx).InsertIgnore(entity.FinanceDailyKline{
		Code:       stockCode,
		Timestamp:  lastKline.Timestamp,
		OpenPrice:  lastKline.OpenPrice,
		ClosePrice: lastKline.ClosePrice,
		HighPrice:  lastKline.HighPrice,
		LowPrice:   lastKline.LowPrice,
		Volume:     lastKline.Volume,
		Turnover:   lastKline.Turnover,
		Md5:        average5,
		Md10:       average10,
		Md20:       average20,
		Md30:       average30,
		Md60:       average60,
		Scale:      lastKline.Scale,
		Key:        lastKline.Key,
		Day:        lastKline.Day,
	})

	return
}

// 5日 10日 20日 30日
func calculateMovingAverage(klineList []*entity.FinanceKline, days int) (average float64) {
	sum := 0.0
	if len(klineList) > days {
		for i := 0; i < days; i++ {
			sum += klineList[i].ClosePrice
		}
	}
	average = sum / float64(days)
	return
}

// CalculateMAOptimized 优化版本，使用滑动窗口避免重复求和
func (s *sSysFinanceDailyKline) CalculateMAOptimized(klines []*entity.FinanceKline) []entity.FinanceDailyKline {
	result := make([]entity.FinanceDailyKline, len(klines))

	// 为每个周期维护滑动窗口
	windows := map[int]*MAWindow{
		5:  NewMAWindow(5),
		10: NewMAWindow(10),
		20: NewMAWindow(20),
		30: NewMAWindow(30),
		50: NewMAWindow(50),
		60: NewMAWindow(60),
	}

	for i, kline := range klines {
		var dailyKline entity.FinanceDailyKline
		_ = gconv.Scan(kline, &dailyKline)

		result[i] = dailyKline

		// 更新所有窗口并计算MA
		for period, window := range windows {
			window.Add(klines[i].ClosePrice)
			switch period {
			case 5:
				result[i].Md5 = window.GetMA()
			case 10:
				result[i].Md10 = window.GetMA()
			case 20:
				result[i].Md20 = window.GetMA()
			case 30:
				result[i].Md30 = window.GetMA()
			case 50:
				result[i].Md50 = window.GetMA()
			case 60:
				result[i].Md60 = window.GetMA()
			}
		}
	}

	return result
}

// MAWindow 移动平均窗口
type MAWindow struct {
	period int
	values []float64
	sum    float64
	count  int
}

func NewMAWindow(period int) *MAWindow {
	return &MAWindow{
		period: period,
		values: make([]float64, 0, period),
	}
}

func (w *MAWindow) Add(value float64) {
	if len(w.values) == w.period {
		// 移除最旧的值
		w.sum -= w.values[0]
		w.values = w.values[1:]
	}

	w.values = append(w.values, value)
	w.sum += value
	w.count = len(w.values)
}

func (w *MAWindow) GetMA() float64 {
	if w.count < w.period {
		return 0
	}
	return w.sum / float64(w.period)
}
