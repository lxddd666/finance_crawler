package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/do"
	"hotgo/internal/service"
	"hotgo/utility/stock"

	"hotgo/internal/model/entity"

	"hotgo/utility/simple"

	"sync"
)

func (s *sSysFinanceCode) CodeDailyKlineStart(ctx context.Context) (err error) {
	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		// 获取当前未完成任务
		codeDaily := make([]*entity.FinanceCodeDaily, 0)
		err = dao.FinanceCodeDaily.Ctx(ctx).WhereNot(dao.FinanceCodeDaily.Columns().Status, consts.TaskComplete).Scan(&codeDaily)
		if err != nil {
			return
		}
		proxyFlag := true
		// 创建大小为5的并发限制通道
		concurrencyLimit := 15
		semaphore := make(chan struct{}, concurrencyLimit)
		wg := sync.WaitGroup{}
		for _, code := range codeDaily {
			wg.Add(1)
			// 获取信号量，控制并发数
			semaphore <- struct{}{}
			simple.SafeGo(gctx.New(), func(ctx context.Context) {
				defer wg.Done()
				defer func() {
					// 释放信号量
					<-semaphore
				}()
				// sz002095
				stockCode := fmt.Sprintf("%s%s", gstr.ToLower(code.Exchange), code.Code)
				_, gErr := service.SysFinanceKline().Kline(ctx, stockCode, consts.MaNo, consts.ScaleDay, 2, proxyFlag)
				if gErr != nil {
					_, _ = dao.FinanceCodeDaily.Ctx(ctx).Where(dao.FinanceCodeDaily.Columns().Code, code.Code).Update(do.FinanceCodeDaily{Status: consts.TaskFail})
				} else {
					// 爬取数据成功
					_, _ = dao.FinanceCodeDaily.Ctx(ctx).Where(dao.FinanceCodeDaily.Columns().Code, code.Code).Update(do.FinanceCodeDaily{Status: consts.TaskComplete})
				}
			})
		}
		wg.Wait()
	})
	return
}

// DailyIndicator 获取每日指标
func (s *sSysFinanceCode) DailyIndicator(ctx context.Context) (err error) {
	codeList, err := s.GetAllCode(ctx)
	if err != nil {
		return
	}
	wg := &sync.WaitGroup{}
	concurrencyLimit := 50
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
			klineList, gErr := s.GetCodeKline(ctx, code, 50)
			if gErr != nil {
				err = gErr
				return
			}
			_, _, _ = service.SysFinanceBoll().Boll(ctx, klineList, consts.BollDefaultMultiple2)
			// macd
			_ = service.SysFinanceMacd().Macd(ctx, klineList, consts.MacdDefaultSlowPeriod12, consts.MacdDefaultFastPeriod26, consts.MacdDefaultSignalPeriod9)
			//// kdj
			_ = service.SysFinanceKdj().Kdj(ctx, klineList, consts.KdjDefaultPeriod9)
		})
	}
	wg.Wait()
	return
}
