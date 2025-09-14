// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financeboll"
)

var (
	FinanceBoll = cFinanceBoll{}
)

type cFinanceBoll struct{}

// List 查看boll带列表
func (c *cFinanceBoll) List(ctx context.Context, req *financeboll.ListReq) (res *financeboll.ListRes, err error) {
	//list, totalCount, err := service.SysFinanceBoll().List(ctx, &req.FinanceBollListInp)
	//if err != nil {
	//	return
	//}
	//
	//if list == nil {
	//	list = []*sysin.FinanceBollListModel{}
	//}
	//
	//res = new(financeboll.ListRes)
	//res.List = list
	//res.PageRes.Pack(req, totalCount)
	//return
	return
}
