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
	"github.com/gogf/gf/v2/text/gstr"
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

func (s *sSysFinanceDailyKline) MovingAverage(ctx context.Context) (err error) {
	// 获取所有均线数据
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	if err != nil {
		return
	}
	for _, co := range codeList {
		fmt.Println(co)
		stockCode := fmt.Sprintf("%s%s", gstr.ToLower(co.Exchange), co.Code)
		// 获取股票
		list, err := service.SysFinanceKline().GetCodeAllKline(ctx, stockCode)
		if err != nil {
			return
		}
		fmt.Println(list)
	}
	return
}

// 5日 10日 20日 30日
func calculateMovingAverage(klineList []*entity.FinanceKline, days int) (average float64, err error) {
	if len(klineList) > days {

	}
	return
}
