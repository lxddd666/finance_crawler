// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financescreening"
	"hotgo/internal/service"
)

var (
	FinanceScreening = cFinanceScreening{}
)

type cFinanceScreening struct{}

// List 查看筛股列表
func (c *cFinanceScreening) List(ctx context.Context, req *financescreening.ListReq) (res *financescreening.ListRes, err error) {

	return
}

// ScreeningDaily 日常筛股
func (c *cFinanceScreening) ScreeningDaily(ctx context.Context, req *financescreening.ScreeningDailyReq) (res *financescreening.ScreeningDailyRes, err error) {
	err = service.SysFinanceScreening().ScreeningDaily(ctx)
	return
}
