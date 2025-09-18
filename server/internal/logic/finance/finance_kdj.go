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
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
)

type sSysFinanceKdj struct{}

func NewSysFinanceKdj() *sSysFinanceKdj {
	return &sSysFinanceKdj{}
}

func init() {
	service.RegisterSysFinanceKdj(NewSysFinanceKdj())
}

// Model kdjORM模型
func (s *sSysFinanceKdj) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceKdj.Ctx(ctx), option...)
}

// Kdj Kdj计算
func (s *sSysFinanceKdj) Kdj(ctx context.Context, data []*entity.FinanceKline, period int) (results []*entity.FinanceKdj, err error) {
	if len(data) < period {
		return nil, nil
	}

	// 设置默认周期为9
	if period == 0 {
		period = 9
	}

	results = make([]*entity.FinanceKdj, len(data))

	// 初始化K、D值（通常设为50）
	prevK := 50.0
	prevD := 50.0

	for i := 0; i < len(data); i++ {
		// 如果是前period-1天，无法计算完整的KDJ，只记录日期
		if i < period-1 {
			results[i] = &entity.FinanceKdj{
				Day: data[i].Day,
				K:   0,
				D:   0,
				J:   0,
			}
			continue
		}

		// 计算最近period天内的最高价和最低价
		highestHigh := data[i].HighPrice
		lowestLow := data[i].LowPrice
		for j := i - period + 1; j <= i; j++ {
			if data[j].HighPrice > highestHigh {
				highestHigh = data[j].HighPrice
			}
			if data[j].LowPrice < lowestLow {
				lowestLow = data[j].LowPrice
			}
		}

		// 计算RSV值
		var rsv float64
		if math.Abs(highestHigh-lowestLow) < 1e-6 { // 避免除以0
			rsv = 50.0 // 如果最高价等于最低价，设为50
		} else {
			rsv = (data[i].ClosePrice - lowestLow) / (highestHigh - lowestLow) * 100
		}

		// 计算K值 (当日K值 = 2/3 * 前一日K值 + 1/3 * 当日RSV)
		k := (2.0/3.0)*prevK + (1.0/3.0)*rsv

		// 计算D值 (当日D值 = 2/3 * 前一日D值 + 1/3 * 当日K值)
		d := (2.0/3.0)*prevD + (1.0/3.0)*k

		// 计算J值 (J = 3*K - 2*D)
		j := 3*k - 2*d

		// 更新前一日K、D值
		prevK = k
		prevD = d

		// 保存结果
		results[i] = &entity.FinanceKdj{
			Day:        data[i].Day,
			Code:       data[i].Code,
			K:          k,
			D:          d,
			J:          j,
			ClosePrice: data[i].ClosePrice,
			HighPrice:  data[i].HighPrice,
			LowPrice:   data[i].LowPrice,
			Scale:      consts.ScaleDay,
			Timestamp:  format.DayStrToTimestamp(data[i].Day),
			Key:        fmt.Sprintf("%s%s", data[i].Code, data[i].Day),
		}
	}
	if len(results) > 0 {
		_, err = dao.FinanceKdj.Ctx(ctx).InsertIgnore(results)
	}

	return
}

// CheckKDJBuySignal 判断KDJ买入信号
func (s *sSysFinanceKdj) CheckKDJBuySignal(kdjValues []*entity.FinanceKdj, currentIndex int) sysin.BuySignal {
	if currentIndex < 1 || currentIndex >= len(kdjValues) {
		return sysin.BuySignal{IsValid: false}
	}

	current := kdjValues[currentIndex]
	previous := kdjValues[currentIndex-1]

	var signals []sysin.BuySignal

	// 1. 超卖区域金叉判断 (K从下向上穿越D，且在20以下)
	if previous.K < previous.D && current.K > current.D && current.K < 20 && current.D < 20 {
		strength := "强"
		if current.K < 10 && current.D < 10 {
			strength = "极强"
		}
		signals = append(signals, sysin.BuySignal{
			IsValid:      true,
			SignalType:   "超卖金叉",
			Strength:     strength,
			Description:  fmt.Sprintf("K值(%.2f)从下向上穿越D值(%.2f)，处于超卖区域", current.K, current.D),
			CurrentValue: current,
		})
	}

	// 2. 底部背离判断 (价格创新低，但KDJ指标未创新低)
	// 注意：这里需要价格数据，需要额外传入价格参数

	// 3. J值极低判断 (J值小于0)
	if current.J < 0 {
		strength := "中"
		if current.J < -10 {
			strength = "强"
		}
		signals = append(signals, sysin.BuySignal{
			IsValid:      true,
			SignalType:   "J值超卖",
			Strength:     strength,
			Description:  fmt.Sprintf("J值(%.2f)极低，表明超卖严重", current.J),
			CurrentValue: current,
		})
	}

	// 4. KDJ三线均低于20
	if current.K < 20 && current.D < 20 && current.J < 20 {
		signals = append(signals, sysin.BuySignal{
			IsValid:      true,
			SignalType:   "三线超卖",
			Strength:     "中",
			Description:  "KDJ三线均低于20，处于超卖区域",
			CurrentValue: current,
		})
	}

	// 5. K值从低位向上拐头
	if current.K > previous.K && previous.K < 20 {
		signals = append(signals, sysin.BuySignal{
			IsValid:      true,
			SignalType:   "K值拐头",
			Strength:     "弱",
			Description:  fmt.Sprintf("K值从%.2f上升至%.2f，显示反弹迹象", previous.K, current.K),
			CurrentValue: current,
		})
	}

	// 如果没有找到信号，返回无效信号
	if len(signals) == 0 {
		return sysin.BuySignal{IsValid: false}
	}

	// 找出最强的信号
	strongestSignal := signals[0]
	for _, signal := range signals {
		if getSignalStrengthValue(signal.Strength) > getSignalStrengthValue(strongestSignal.Strength) {
			strongestSignal = signal
		}
	}

	return strongestSignal
}

// 辅助函数：将信号强度转换为可比较的值
func getSignalStrengthValue(strength string) int {
	switch strength {
	case "极强":
		return 4
	case "强":
		return 3
	case "中":
		return 2
	case "弱":
		return 1
	default:
		return 0
	}
}
