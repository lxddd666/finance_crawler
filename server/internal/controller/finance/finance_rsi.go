// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financersi"
)

var (
	FinanceRsi = cFinanceRsi{}
)

type cFinanceRsi struct{}

// List 查看rsi线列表
func (c *cFinanceRsi) List(ctx context.Context, req *financersi.ListReq) (res *financersi.ListRes, err error) {

	return
}
