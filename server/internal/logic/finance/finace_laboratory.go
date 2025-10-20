package sys

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"time"
)

// MovingAverageLaboratory 移动平均线试验
func (s *sSysTestFinance) MovingAverageLaboratory(ctx context.Context) {
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	if err != nil {
		return
	}

	for _, code := range codeList {
		klineList, _ := service.SysFinanceCode().GetCodeKline(ctx, code.CompleteCode, 100)
		var trendList []*entity.FinanceKlineZigzagTrend
		var trendListScan []*entity.FinanceKlineZigzagTrend
		// 100条k线所有时间的zigzag标记
		trendList = append(trendList, &entity.FinanceKlineZigzagTrend{Code: klineList[0].Code, KlineId: klineList[0].Id, Key: klineList[0].Key, MinChangePercent: 1, Day: klineList[0].Day})
		_ = dao.FinanceKlineZigzagTrend.Ctx(ctx).Where(dao.FinanceKlineZigzagTrend.Columns().Code, code.CompleteCode).Scan(&trendListScan)
		trendList = append(trendList, trendListScan...)
		trendList = append(trendList, &entity.FinanceKlineZigzagTrend{Code: klineList[len(klineList)-1].Code, KlineId: klineList[len(klineList)-1].Id, Key: klineList[len(klineList)-1].Key, MinChangePercent: 1, Day: klineList[len(klineList)-1].Day})

		//trendDayMap := make(map[string]float64)
		//for _, kline := range klineList {
		//
		//}
	}
}

// 判断目标日期是否在起始日期和结束日期之间（包含边界）
func isDateInRange(startDateStr, endDateStr, targetDateStr string) (bool, error) {
	// 解析日期字符串，格式为 "2006-01-02"
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return false, fmt.Errorf("解析起始日期失败: %v", err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return false, fmt.Errorf("解析结束日期失败: %v", err)
	}

	targetDate, err := time.Parse("2006-01-02", targetDateStr)
	if err != nil {
		return false, fmt.Errorf("解析目标日期失败: %v", err)
	}

	// 判断目标日期是否在区间内（包含边界）
	if (targetDate.Equal(startDate) || targetDate.After(startDate)) &&
		(targetDate.Equal(endDate) || targetDate.Before(endDate)) {
		return true, nil
	}

	return false, nil
}
