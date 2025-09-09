package sys

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"hotgo/utility/simple"
	"time"
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
		for _, code := range codeDaily {
			// sz002095
			stockCode := fmt.Sprintf("%s%s", gstr.ToLower(code.Exchange), code.Code)
			_, gErr := service.SysStockIndicator().Kline(ctx, stockCode, consts.MaNo, consts.ScaleDay, 50, proxyFlag)
			if gErr != nil {
				_, _ = dao.FinanceCodeDaily.Ctx(ctx).Where(dao.FinanceCodeDaily.Columns().Code, code.Code).Update(do.FinanceCodeDaily{Status: consts.TaskFail})

				if !proxyFlag {
					// 切换代理爬
					proxyFlag = true
				} else {
					time.Sleep(6 * time.Second)
				}
			} else {
				// 爬取数据成功
				_, _ = dao.FinanceCodeDaily.Ctx(ctx).Where(dao.FinanceCodeDaily.Columns().Code, code.Code).Update(do.FinanceCodeDaily{Status: consts.TaskComplete})
			}
		}
	})
	return
}
