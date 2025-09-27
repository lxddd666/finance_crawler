// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/testfinance"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
)

var (
	TestFinance = cTestFinance{}
)

type cTestFinance struct{}

// List 查看测试分类列表
func (c *cTestFinance) List(ctx context.Context, req *testfinance.ListReq) (res *testfinance.ListRes, err error) {
	list, totalCount, err := service.SysTestFinance().List(ctx, &req.TestFinanceListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.TestFinanceListModel{}
	}

	res = new(testfinance.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出测试分类列表
func (c *cTestFinance) Export(ctx context.Context, req *testfinance.ExportReq) (res *testfinance.ExportRes, err error) {
	err = service.SysTestFinance().Export(ctx, &req.TestFinanceListInp)
	return
}

// Edit 更新测试分类
func (c *cTestFinance) Edit(ctx context.Context, req *testfinance.EditReq) (res *testfinance.EditRes, err error) {
	err = service.SysTestFinance().Edit(ctx, &req.TestFinanceEditInp)
	return
}

// MaxSort 获取测试分类最大排序
func (c *cTestFinance) MaxSort(ctx context.Context, req *testfinance.MaxSortReq) (res *testfinance.MaxSortRes, err error) {
	data, err := service.SysTestFinance().MaxSort(ctx, &req.TestFinanceMaxSortInp)
	if err != nil {
		return
	}

	res = new(testfinance.MaxSortRes)
	res.TestFinanceMaxSortModel = data
	return
}

// View 获取指定测试分类信息
func (c *cTestFinance) View(ctx context.Context, req *testfinance.ViewReq) (res *testfinance.ViewRes, err error) {
	data, err := service.SysTestFinance().View(ctx, &req.TestFinanceViewInp)
	if err != nil {
		return
	}

	res = new(testfinance.ViewRes)
	res.TestFinanceViewModel = data
	return
}

// Delete 删除测试分类
func (c *cTestFinance) Delete(ctx context.Context, req *testfinance.DeleteReq) (res *testfinance.DeleteRes, err error) {
	err = service.SysTestFinance().Delete(ctx, &req.TestFinanceDeleteInp)
	return
}

// Status 更新测试分类状态
func (c *cTestFinance) Status(ctx context.Context, req *testfinance.StatusReq) (res *testfinance.StatusRes, err error) {
	err = service.SysTestFinance().Status(ctx, &req.TestFinanceStatusInp)
	return
}

// Start 开始测试
func (c *cTestFinance) Start(ctx context.Context, req *testfinance.StartReq) (res *testfinance.StartRes, err error) {
	//err = service.SysTestFinance().Start(ctx)
	service.SysTestFinance().MovingAverageLaboratory(ctx)
	return
}
