// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financekdj"
)

var (
	FinanceKdj = cFinanceKdj{}
)

type cFinanceKdj struct{}

// List 查看kdj列表
func (c *cFinanceKdj) List(ctx context.Context, req *financekdj.ListReq) (res *financekdj.ListRes, err error) {
	return
}
