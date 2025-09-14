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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"hotgo/utility/simple"
	"hotgo/utility/stock"
	"sync"
)

type sSysFinanceScreening struct{}

func NewSysFinanceScreening() *sSysFinanceScreening {
	return &sSysFinanceScreening{}
}

func init() {
	service.RegisterSysFinanceScreening(NewSysFinanceScreening())
}

// Model 筛股ORM模型
func (s *sSysFinanceScreening) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceScreening.Ctx(ctx), option...)
}

// ScreeningDaily 日常筛股
func (s *sSysFinanceScreening) ScreeningDaily(ctx context.Context) (err error) {
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	if err != nil {
		return
	}
	wg := &sync.WaitGroup{}
	concurrencyLimit := 100
	semaphore := make(chan struct{}, concurrencyLimit)

	for _, financeCode := range codeList {
		wg.Add(1)
		semaphore <- struct{}{}

		simple.SafeGo(gctx.New(), func(ctx context.Context) {
			defer wg.Done()
			defer func() {
				// 释放信号量
				<-semaphore
			}()
			code := stock.GetCode(financeCode.Code, financeCode.Exchange)
			screening := &entity.FinanceScreening{
				Code:      code,
				Scale:     240,
				Day:       "2025-09-15",
				Timestamp: format.DayStrToTimestamp("2025-09-15"),
				Key:       fmt.Sprintf("%s%s", code, "2025-09-15"),
			}

			var boll *entity.FinanceBoll
			err = dao.FinanceBoll.Ctx(ctx).Where(dao.FinanceBoll.Columns().Code, code).OrderDesc(dao.FinanceBoll.Columns().Day).Limit(1).Scan(&boll)
			if err != nil {
				return
			}
			if boll == nil {
				return
			}
			if boll.Degree < 0.15 {
				screening.Boll = 1
				screening.MatchCount++
			}
			var macd *entity.FinanceMacd
			err = dao.FinanceMacd.Ctx(ctx).Where(dao.FinanceMacd.Columns().Code, code).OrderDesc(dao.FinanceBoll.Columns().Day).Limit(1).Scan(&macd)
			if macd == nil {
				return
			}
			if macd.Macd > 0 {
				screening.Macd = 1
				screening.MatchCount++
			}
			var kdj *entity.FinanceKdj
			err = dao.FinanceKdj.Ctx(ctx).Where(dao.FinanceMacd.Columns().Code, code).OrderDesc(dao.FinanceKdj.Columns().Day).Limit(1).Scan(&kdj)
			if kdj == nil {
				return
			}
			if kdj.K < 20 && kdj.D < 20 && kdj.J < 20 {
				screening.Kdj = 1
				screening.MatchCount++
			}
			dao.FinanceScreening.Ctx(ctx).InsertIgnore(screening)
		})
	}
	wg.Wait()
	return
}
