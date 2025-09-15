// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financedailykline"
)

var (
	FinanceDailyKline = cFinanceDailyKline{}
)

type cFinanceDailyKline struct{}

// List 查看股票日K线数据表列表
func (c *cFinanceDailyKline) List(ctx context.Context, req *financedailykline.ListReq) (res *financedailykline.ListRes, err error) {

	return
}
