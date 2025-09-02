// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financealltickresponse"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
)

var (
	FinanceAlltickResponse = cFinanceAlltickResponse{}
)

type cFinanceAlltickResponse struct{}

// List 查看alltick返回值列表
func (c *cFinanceAlltickResponse) List(ctx context.Context, req *financealltickresponse.ListReq) (res *financealltickresponse.ListRes, err error) {
	list, totalCount, err := service.SysFinanceAlltickResponse().List(ctx, &req.FinanceAlltickResponseListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.FinanceAlltickResponseListModel{}
	}

	res = new(financealltickresponse.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出alltick返回值列表
func (c *cFinanceAlltickResponse) Export(ctx context.Context, req *financealltickresponse.ExportReq) (res *financealltickresponse.ExportRes, err error) {
	err = service.SysFinanceAlltickResponse().Export(ctx, &req.FinanceAlltickResponseListInp)
	return
}

// Edit 更新alltick返回值
func (c *cFinanceAlltickResponse) Edit(ctx context.Context, req *financealltickresponse.EditReq) (res *financealltickresponse.EditRes, err error) {
	err = service.SysFinanceAlltickResponse().Edit(ctx, &req.FinanceAlltickResponseEditInp)
	return
}

// MaxSort 获取alltick返回值最大排序
func (c *cFinanceAlltickResponse) MaxSort(ctx context.Context, req *financealltickresponse.MaxSortReq) (res *financealltickresponse.MaxSortRes, err error) {
	data, err := service.SysFinanceAlltickResponse().MaxSort(ctx, &req.FinanceAlltickResponseMaxSortInp)
	if err != nil {
		return
	}

	res = new(financealltickresponse.MaxSortRes)
	res.FinanceAlltickResponseMaxSortModel = data
	return
}

// View 获取指定alltick返回值信息
func (c *cFinanceAlltickResponse) View(ctx context.Context, req *financealltickresponse.ViewReq) (res *financealltickresponse.ViewRes, err error) {
	data, err := service.SysFinanceAlltickResponse().View(ctx, &req.FinanceAlltickResponseViewInp)
	if err != nil {
		return
	}

	res = new(financealltickresponse.ViewRes)
	res.FinanceAlltickResponseViewModel = data
	return
}

// Delete 删除alltick返回值
func (c *cFinanceAlltickResponse) Delete(ctx context.Context, req *financealltickresponse.DeleteReq) (res *financealltickresponse.DeleteRes, err error) {
	err = service.SysFinanceAlltickResponse().Delete(ctx, &req.FinanceAlltickResponseDeleteInp)
	return
}
