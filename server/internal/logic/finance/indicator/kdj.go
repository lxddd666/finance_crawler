package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/utility/format"
	"math"
)

// Kdj Kdj计算
func (s *sSysStockIndicator) Kdj(ctx context.Context, data []*entity.FinanceKline, period int) []*entity.FinanceKdj {
	if len(data) < period {
		return nil
	}

	// 设置默认周期为9
	if period == 0 {
		period = 9
	}

	results := make([]*entity.FinanceKdj, len(data))

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
		_, _ = dao.FinanceKdj.Ctx(ctx).InsertIgnore(results)
	}

	return results
}
