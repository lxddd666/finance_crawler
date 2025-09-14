// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financekline"
)

var (
	FinanceKline = cFinanceKline{}
)

type cFinanceKline struct{}

// List 查看k线列表
func (c *cFinanceKline) List(ctx context.Context, req *financekline.ListReq) (res *financekline.ListRes, err error) {
	//list, totalCount, err := service.SysFinanceKline().List(ctx, &req.FinanceKlineListInp)
	//if err != nil {
	//	return
	//}
	//
	//if list == nil {
	//	list = []*sysin.FinanceKlineListModel{}
	//}
	//
	//res = new(financekline.ListRes)
	//res.List = list
	//res.PageRes.Pack(req, totalCount)
	//return
	return
}
