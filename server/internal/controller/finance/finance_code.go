// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"hotgo/api/admin/financecode"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/service"
)

var (
	FinanceCode = cFinanceCode{}
)

type cFinanceCode struct{}

// List 查看股票代码列表
func (c *cFinanceCode) List(ctx context.Context, req *financecode.ListReq) (res *financecode.ListRes, err error) {
	list, totalCount, err := service.SysFinanceCode().List(ctx, &req.FinanceCodeListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.FinanceCodeListModel{}
	}

	res = new(financecode.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出股票代码列表
func (c *cFinanceCode) Export(ctx context.Context, req *financecode.ExportReq) (res *financecode.ExportRes, err error) {
	err = service.SysFinanceCode().Export(ctx, &req.FinanceCodeListInp)
	return
}

// Edit 更新股票代码
func (c *cFinanceCode) Edit(ctx context.Context, req *financecode.EditReq) (res *financecode.EditRes, err error) {
	err = service.SysFinanceCode().Edit(ctx, &req.FinanceCodeEditInp)
	return
}

// View 获取指定股票代码信息
func (c *cFinanceCode) View(ctx context.Context, req *financecode.ViewReq) (res *financecode.ViewRes, err error) {
	data, err := service.SysFinanceCode().View(ctx, &req.FinanceCodeViewInp)
	if err != nil {
		return
	}

	res = new(financecode.ViewRes)
	res.FinanceCodeViewModel = data
	return
}

// Delete 删除股票代码
func (c *cFinanceCode) Delete(ctx context.Context, req *financecode.DeleteReq) (res *financecode.DeleteRes, err error) {
	err = service.SysFinanceCode().Delete(ctx, &req.FinanceCodeDeleteInp)
	return
}

// ImportCode 导入股票代码
func (c *cFinanceCode) ImportCode(ctx context.Context, req *financecode.ImportCodeReq) (res *financecode.ImportCodeRes, err error) {
	err = service.SysFinanceCode().ImportCode(ctx, req.FinanceImportCodeInp)
	return
}

func (c *cFinanceCode) CodeDailyKlineStart(ctx context.Context, req *financecode.CodeDailyKlineStartReq) (res *financecode.CodeDailyKlineStartRes, err error) {
	err = service.SysFinanceCode().CodeDailyKlineStart(ctx)
	return
}
