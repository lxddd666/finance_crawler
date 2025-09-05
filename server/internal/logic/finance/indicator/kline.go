package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/logic/alltick"
	"hotgo/internal/model/entity"
)

// Kline k线
func (s *sSysStockIndicator) Kline(ctx context.Context, code string, klineType, klineNum int) (klineList []*entity.FinanceKline, err error) {
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
	klineListResp := response.Data.KlineList

	for _, kline := range klineListResp {
		klineList = append(klineList, &entity.FinanceKline{
			Code:       response.Data.Code,
			Timestamp:  gconv.Int64(kline.Timestamp),
			OpenPrice:  gconv.Float64(kline.OpenPrice),
			ClosePrice: gconv.Float64(kline.ClosePrice),
			HighPrice:  gconv.Float64(kline.HighPrice),
			LowPrice:   gconv.Float64(kline.LowPrice),
			Volume:     gconv.Int64(kline.Volume),
			Turnover:   gconv.Float64(kline.Turnover),
			Key:        fmt.Sprintf("%s%s", code, kline.Timestamp),
		})
	}
	_, err = dao.FinanceKline.Ctx(ctx).InsertIgnore(klineList)
	return
}
