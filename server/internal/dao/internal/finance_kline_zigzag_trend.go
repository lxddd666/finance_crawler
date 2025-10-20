// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceKlineZigzagTrendDao is the data access object for the table hg_finance_kline_zigzag_trend.
type FinanceKlineZigzagTrendDao struct {
	table    string                         // table is the underlying table name of the DAO.
	group    string                         // group is the database configuration group name of the current DAO.
	columns  FinanceKlineZigzagTrendColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler             // handlers for customized model modification.
}

// FinanceKlineZigzagTrendColumns defines and stores column names for the table hg_finance_kline_zigzag_trend.
type FinanceKlineZigzagTrendColumns struct {
	Id               string // 主键ID
	Code             string // 股票代码
	KlineId          string // K线id
	Key              string // 唯一键: timestamp_code_scale
	MinChangePercent string // 最小浮动点
	Day              string // 日期(yyyy-MM-dd)
}

// financeKlineZigzagTrendColumns holds the columns for the table hg_finance_kline_zigzag_trend.
var financeKlineZigzagTrendColumns = FinanceKlineZigzagTrendColumns{
	Id:               "id",
	Code:             "code",
	KlineId:          "kline_id",
	Key:              "key",
	MinChangePercent: "min_change_percent",
	Day:              "day",
}

// NewFinanceKlineZigzagTrendDao creates and returns a new DAO object for table data access.
func NewFinanceKlineZigzagTrendDao(handlers ...gdb.ModelHandler) *FinanceKlineZigzagTrendDao {
	return &FinanceKlineZigzagTrendDao{
		group:    "default",
		table:    "hg_finance_kline_zigzag_trend",
		columns:  financeKlineZigzagTrendColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceKlineZigzagTrendDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceKlineZigzagTrendDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceKlineZigzagTrendDao) Columns() FinanceKlineZigzagTrendColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceKlineZigzagTrendDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceKlineZigzagTrendDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinanceKlineZigzagTrendDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
