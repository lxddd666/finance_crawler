package sys

import (
	"context"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
)

func (s *sSysTestFinance) MovingAverageLaboratory(ctx context.Context) {
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	if err != nil {
		return
	}

	for _, code := range codeList {
		klineList, _ := service.SysFinanceCode().GetCodeKline(ctx, code.CompleteCode, 0)
		var trendList []*entity.FinanceKlineZigzagTrend
		_ = dao.FinanceKlineZigzagTrend.Ctx(ctx).Where(dao.FinanceKlineZigzagTrend.Columns().Code, code.CompleteCode).Scan(&trendList)
		klineDailyMap := make(map[string]*entity.FinanceKline)
		for _, kline := range klineList {
			klineDailyMap[kline.Day] = kline
		}

	}
}
