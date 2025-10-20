package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/model/do"
	"hotgo/internal/service"
	"hotgo/utility/stock"

	"hotgo/internal/model/entity"

	"hotgo/utility/simple"

	"sync"
)

func (s *sSysFinanceCode) CodeDailyKlineStart(ctx context.Context) (err error) {
	simple.SafeGo(gctx.New(), func(ctx context.Context) {

		defer func() {
			// 递归
			count, _ := dao.FinanceCodeDaily.Ctx(ctx).WhereNot(dao.FinanceCodeDaily.Columns().Status, consts.TaskComplete).Count()
			if count > 0 {
				if global.ProxyList.Size() > 0 {
					_ = s.CodeDailyKlineStart(ctx)
				}
			}

		}()
		// 获取当前未完成任务
		codeDaily := make([]*entity.FinanceCodeDaily, 0)
		err = dao.FinanceCodeDaily.Ctx(ctx).WhereNot(dao.FinanceCodeDaily.Columns().Status, consts.TaskComplete).Scan(&codeDaily)
		if err != nil {
			return
		}
		proxyFlag := true
		// 创建大小为5的并发限制通道
		concurrencyLimit := 30
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
				_, gErr := service.SysFinanceKline().Kline(ctx, stockCode, consts.MaNo, consts.ScaleFiveDay, 300, proxyFlag)
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
	simple.SafeGo(gctx.New(), func(ctx context.Context) {
		defer func() {
			// 递归
			count, _ := dao.FinanceIndicatorDaily.Ctx(ctx).WhereNot(dao.FinanceIndicatorDaily.Columns().Status, consts.TaskComplete).Count()
			if count > 0 {
				if global.ProxyList.Size() > 0 {
					_ = s.DailyIndicator(ctx)
				}
			}

		}()

		codeDaily := make([]*entity.FinanceCodeDaily, 0)
		err = dao.FinanceIndicatorDaily.Ctx(ctx).WhereNot(dao.FinanceIndicatorDaily.Columns().Status, consts.TaskComplete).Scan(&codeDaily)
		if err != nil {
			return
		}
		wg := &sync.WaitGroup{}
		concurrencyLimit := 30
		semaphore := make(chan struct{}, concurrencyLimit)

		for _, financeCode := range codeDaily {
			wg.Add(1)
			semaphore <- struct{}{}

			simple.SafeGo(gctx.New(), func(ctx context.Context) {
				defer wg.Done()
				defer func() {
					if err != nil {
						_, _ = dao.FinanceIndicatorDaily.Ctx(ctx).Where(dao.FinanceIndicatorDaily.Columns().Code, financeCode.Code).Update(do.FinanceCodeDaily{Status: consts.TaskFail})
					} else {
						_, _ = dao.FinanceIndicatorDaily.Ctx(ctx).Where(dao.FinanceIndicatorDaily.Columns().Code, financeCode.Code).Update(do.FinanceCodeDaily{Status: consts.TaskComplete})
					}
					// 释放信号量
					<-semaphore
				}()
				code := stock.GetCode(financeCode.Code, financeCode.Exchange)
				//klineList, gErr := s.GetCodeKline(ctx, code, 0)
				//if gErr != nil {
				//	err = gErr
				//	return
				//}
				//_, _, err = service.SysFinanceBoll().Boll(ctx, klineList, consts.BollDefaultMultiple2)
				//// macd
				//_, err = service.SysFinanceMacd().Macd(ctx, klineList, consts.MacdDefaultSlowPeriod12, consts.MacdDefaultFastPeriod26, consts.MacdDefaultSignalPeriod9)
				////// kdj
				//_, err = service.SysFinanceKdj().Kdj(ctx, klineList, consts.KdjDefaultPeriod9)
				//// 均线
				//stock.ReverseKline(klineList)
				//err = service.SysFinanceDailyKline().MovingAverage(ctx, klineList)
				err = service.SysFinanceMacd().MacdV2(ctx, code, consts.ScaleFiveDay)
				if err != nil {
					return
				}
				err = service.SysFinanceRsi().Rsi(ctx, code, consts.ScaleFiveDay)
				if err != nil {
					return
				}
			})
		}
		wg.Wait()
	})
	return
}
