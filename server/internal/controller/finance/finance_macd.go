// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financemacd"
)

var (
	FinanceMacd = cFinanceMacd{}
)

type cFinanceMacd struct{}

// List 查看macd线列表
func (c *cFinanceMacd) List(ctx context.Context, req *financemacd.ListReq) (res *financemacd.ListRes, err error) {
	//list, totalCount, err := service.SysFinanceMacd().List(ctx, &req.FinanceMacdListInp)
	//if err != nil {
	//	return
	//}
	//
	//if list == nil {
	//	list = []*sysin.FinanceMacdListModel{}
	//}
	//
	//res = new(financemacd.ListRes)
	//res.List = list
	//res.PageRes.Pack(req, totalCount)
	return
}
