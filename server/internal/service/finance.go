// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/result"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysFinanceCode interface {
		CodeDailyKlineStart(ctx context.Context) (err error)
		// DailyIndicator 获取每日指标
		DailyIndicator(ctx context.Context) (err error)
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
		// GetAllCode 获取所有code
		GetAllCode(ctx context.Context) (codeList []*entity.FinanceCode, err error)
		// GetCodeKline 获取股票k线
		GetCodeKline(ctx context.Context, code string, KlineNum int) (list []*entity.FinanceKline, err error)
	}
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
	ISysFinanceBoll interface {
		// Model boll带ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Boll 用股票k线计算boll
		Boll(ctx context.Context, data []*entity.FinanceKline, multiple int) (resp *result.BollResult, lastKline *entity.FinanceKline, err error)
	}
	ISysFinanceDailyKline interface {
		// Model 股票日K线数据表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// MovingAverage 均线计算
		MovingAverage(ctx context.Context, klineList []*entity.FinanceKline) (err error)
		// CalculateMAOptimized 计算每天均线
		CalculateMAOptimized(klines []*entity.FinanceKline) []entity.FinanceDailyKline
	}
	ISysFinanceKdj interface {
		// Model kdjORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Kdj Kdj计算
		Kdj(ctx context.Context, data []*entity.FinanceKline, period int) (results []*entity.FinanceKdj, err error)
		// CheckKDJBuySignal 判断KDJ买入信号
		CheckKDJBuySignal(kdjValues []*entity.FinanceKdj, currentIndex int) sysin.BuySignal
	}
	ISysFinanceKline interface {
		// Model k线ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Kline k线
		Kline(ctx context.Context, code string, ma string, scale int, datalen int, proxyFlag bool) (klineList []*entity.FinanceKline, err error)
		GetCodeAllKline(ctx context.Context, code string) (klineList []*entity.FinanceKline, err error)
	}
	ISysFinanceMacd interface {
		// Model macd线ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// Macd Macd计算
		Macd(ctx context.Context, data []*entity.FinanceKline, slowPeriod int, fastPeriod int, signalPeriod int) (results []*entity.FinanceMacd, err error)
	}
	ISysFinancePlot interface {
		// Model plot图ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取plot图列表
		List(ctx context.Context, in *sysin.FinancePlotListInp) (list []*sysin.FinancePlotListModel, totalCount int, err error)
		// CreatePlot 指标图创建
		CreatePlot(ctx context.Context, in sysin.FinancePlotCreate) (err error)
		KlinePlot(ctx context.Context, code string)
		MacdPlot(ctx context.Context, code string)
		KdjPlot(ctx context.Context, code string)
		BollPlot(ctx context.Context, code string)
		RsiPlot(ctx context.Context, code string)
		TrendPlot(ctx context.Context, klineList []*entity.FinanceKline, trendList []*entity.FinanceKline) (err error)
	}
	ISysFinanceScreening interface {
		// Model 筛股ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// ScreeningDaily 日常筛股
		ScreeningDaily(ctx context.Context) (err error)
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
		Start(ctx context.Context) (err error)
	}
)

var (
	localSysFinanceCode            ISysFinanceCode
	localSysFinanceAlltickResponse ISysFinanceAlltickResponse
	localSysFinanceBoll            ISysFinanceBoll
	localSysFinanceDailyKline      ISysFinanceDailyKline
	localSysFinanceKdj             ISysFinanceKdj
	localSysFinanceKline           ISysFinanceKline
	localSysFinanceMacd            ISysFinanceMacd
	localSysFinancePlot            ISysFinancePlot
	localSysFinanceScreening       ISysFinanceScreening
	localSysTestFinance            ISysTestFinance
)

func SysFinanceCode() ISysFinanceCode {
	if localSysFinanceCode == nil {
		panic("implement not found for interface ISysFinanceCode, forgot register?")
	}
	return localSysFinanceCode
}

func RegisterSysFinanceCode(i ISysFinanceCode) {
	localSysFinanceCode = i
}

func SysFinanceAlltickResponse() ISysFinanceAlltickResponse {
	if localSysFinanceAlltickResponse == nil {
		panic("implement not found for interface ISysFinanceAlltickResponse, forgot register?")
	}
	return localSysFinanceAlltickResponse
}

func RegisterSysFinanceAlltickResponse(i ISysFinanceAlltickResponse) {
	localSysFinanceAlltickResponse = i
}

func SysFinanceBoll() ISysFinanceBoll {
	if localSysFinanceBoll == nil {
		panic("implement not found for interface ISysFinanceBoll, forgot register?")
	}
	return localSysFinanceBoll
}

func RegisterSysFinanceBoll(i ISysFinanceBoll) {
	localSysFinanceBoll = i
}

func SysFinanceDailyKline() ISysFinanceDailyKline {
	if localSysFinanceDailyKline == nil {
		panic("implement not found for interface ISysFinanceDailyKline, forgot register?")
	}
	return localSysFinanceDailyKline
}

func RegisterSysFinanceDailyKline(i ISysFinanceDailyKline) {
	localSysFinanceDailyKline = i
}

func SysFinanceKdj() ISysFinanceKdj {
	if localSysFinanceKdj == nil {
		panic("implement not found for interface ISysFinanceKdj, forgot register?")
	}
	return localSysFinanceKdj
}

func RegisterSysFinanceKdj(i ISysFinanceKdj) {
	localSysFinanceKdj = i
}

func SysFinanceKline() ISysFinanceKline {
	if localSysFinanceKline == nil {
		panic("implement not found for interface ISysFinanceKline, forgot register?")
	}
	return localSysFinanceKline
}

func RegisterSysFinanceKline(i ISysFinanceKline) {
	localSysFinanceKline = i
}

func SysFinanceMacd() ISysFinanceMacd {
	if localSysFinanceMacd == nil {
		panic("implement not found for interface ISysFinanceMacd, forgot register?")
	}
	return localSysFinanceMacd
}

func RegisterSysFinanceMacd(i ISysFinanceMacd) {
	localSysFinanceMacd = i
}

func SysFinancePlot() ISysFinancePlot {
	if localSysFinancePlot == nil {
		panic("implement not found for interface ISysFinancePlot, forgot register?")
	}
	return localSysFinancePlot
}

func RegisterSysFinancePlot(i ISysFinancePlot) {
	localSysFinancePlot = i
}

func SysFinanceScreening() ISysFinanceScreening {
	if localSysFinanceScreening == nil {
		panic("implement not found for interface ISysFinanceScreening, forgot register?")
	}
	return localSysFinanceScreening
}

func RegisterSysFinanceScreening(i ISysFinanceScreening) {
	localSysFinanceScreening = i
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
