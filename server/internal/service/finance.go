// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/result"
)

type (
	ISysFinanceAlltickResponse interface {
		// Model alltick返回值ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取alltick返回值列表
		List(ctx context.Context, in *sysin.FinanceAlltickResponseListInp) (list []*sysin.FinanceAlltickResponseListModel, totalCount int, err error)
		// Export 导出alltick返回值
		Export(ctx context.Context, in *sysin.FinanceAlltickResponseListInp) (err error)
		// Edit 修改/新增alltick返回值
		Edit(ctx context.Context, in *sysin.FinanceAlltickResponseEditInp) (err error)
		// Delete 删除alltick返回值
		Delete(ctx context.Context, in *sysin.FinanceAlltickResponseDeleteInp) (err error)
		// MaxSort 获取alltick返回值最大排序
		MaxSort(ctx context.Context, in *sysin.FinanceAlltickResponseMaxSortInp) (res *sysin.FinanceAlltickResponseMaxSortModel, err error)
		// View 获取alltick返回值指定信息
		View(ctx context.Context, in *sysin.FinanceAlltickResponseViewInp) (res *sysin.FinanceAlltickResponseViewModel, err error)
	}
	ISysFinanceCode interface {
		// Model 股票代码ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取股票代码列表
		List(ctx context.Context, in *sysin.FinanceCodeListInp) (list []*sysin.FinanceCodeListModel, totalCount int, err error)
		// Export 导出股票代码
		Export(ctx context.Context, in *sysin.FinanceCodeListInp) (err error)
		// Edit 修改/新增股票代码
		Edit(ctx context.Context, in *sysin.FinanceCodeEditInp) (err error)
		// Delete 删除股票代码
		Delete(ctx context.Context, in *sysin.FinanceCodeDeleteInp) (err error)
		// View 获取股票代码指定信息
		View(ctx context.Context, in *sysin.FinanceCodeViewInp) (res *sysin.FinanceCodeViewModel, err error)
		// ImportCode 导入股票代码
		ImportCode(ctx context.Context, inp sysin.FinanceImportCodeInp) (err error)
	}
	ISysStockIndicator interface {
		// Kline K线
		Kline(ctx context.Context, code, ma string, scale, datalen int) (klineList []*entity.FinanceKline, err error)
		// Boll boll带
		Boll(ctx context.Context, code, ma string, scale, datalen, multiple int) (result *entity.FinanceBoll, err error)
		CalculateBoll(data []*entity.FinanceKline, multiple int) (resp *result.BollResult, lastKline *entity.FinanceKline, err error)
	}
	ISysTestFinance interface {
		// ConvertToQueryString 将 FinanceAlltickRequest 结构体转换为 JSON 查询字符串
		ConvertToQueryString(req *entity.FinanceAlltickRequest) (string, error)
		// Model 测试分类ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取测试分类列表
		List(ctx context.Context, in *sysin.TestFinanceListInp) (list []*sysin.TestFinanceListModel, totalCount int, err error)
		// Export 导出测试分类
		Export(ctx context.Context, in *sysin.TestFinanceListInp) (err error)
		// Edit 修改/新增测试分类
		Edit(ctx context.Context, in *sysin.TestFinanceEditInp) (err error)
		// Delete 删除测试分类
		Delete(ctx context.Context, in *sysin.TestFinanceDeleteInp) (err error)
		// MaxSort 获取测试分类最大排序
		MaxSort(ctx context.Context, in *sysin.TestFinanceMaxSortInp) (res *sysin.TestFinanceMaxSortModel, err error)
		// View 获取测试分类指定信息
		View(ctx context.Context, in *sysin.TestFinanceViewInp) (res *sysin.TestFinanceViewModel, err error)
		// Status 更新测试分类状态
		Status(ctx context.Context, in *sysin.TestFinanceStatusInp) (err error)
		// Start 更新测试分类状态
		Start(ctx context.Context) error
	}
)

var (
	localSysFinanceAlltickResponse ISysFinanceAlltickResponse
	localSysFinanceCode            ISysFinanceCode
	localSysStockIndicator         ISysStockIndicator
	localSysTestFinance            ISysTestFinance
)

func SysFinanceAlltickResponse() ISysFinanceAlltickResponse {
	if localSysFinanceAlltickResponse == nil {
		panic("implement not found for interface ISysFinanceAlltickResponse, forgot register?")
	}
	return localSysFinanceAlltickResponse
}

func RegisterSysFinanceAlltickResponse(i ISysFinanceAlltickResponse) {
	localSysFinanceAlltickResponse = i
}

func SysFinanceCode() ISysFinanceCode {
	if localSysFinanceCode == nil {
		panic("implement not found for interface ISysFinanceCode, forgot register?")
	}
	return localSysFinanceCode
}

func RegisterSysFinanceCode(i ISysFinanceCode) {
	localSysFinanceCode = i
}

func SysStockIndicator() ISysStockIndicator {
	if localSysStockIndicator == nil {
		panic("implement not found for interface ISysStockIndicator, forgot register?")
	}
	return localSysStockIndicator
}

func RegisterSysStockIndicator(i ISysStockIndicator) {
	localSysStockIndicator = i
}

func SysTestFinance() ISysTestFinance {
	if localSysTestFinance == nil {
		panic("implement not found for interface ISysTestFinance, forgot register?")
	}
	return localSysTestFinance
}

func RegisterSysTestFinance(i ISysTestFinance) {
	localSysTestFinance = i
}
