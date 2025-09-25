// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceIndicatorDailyDao is the data access object for the table hg_finance_indicator_daily.
type FinanceIndicatorDailyDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of the current DAO.
	columns FinanceIndicatorDailyColumns // columns contains all the column names of Table for convenient usage.
}

// FinanceIndicatorDailyColumns defines and stores column names for the table hg_finance_indicator_daily.
type FinanceIndicatorDailyColumns struct {
	Id        string // 分类ID
	Code      string // 代码
	Name      string // 名称
	Exchange  string // 交易所
	Day       string // 日期
	Timestamp string // 时间戳
	Status    string // 0未开始 1完成 -1失败
}

// financeIndicatorDailyColumns holds the columns for the table hg_finance_indicator_daily.
var financeIndicatorDailyColumns = FinanceIndicatorDailyColumns{
	Id:        "id",
	Code:      "code",
	Name:      "name",
	Exchange:  "exchange",
	Day:       "day",
	Timestamp: "timestamp",
	Status:    "status",
}

// NewFinanceIndicatorDailyDao creates and returns a new DAO object for table data access.
func NewFinanceIndicatorDailyDao() *FinanceIndicatorDailyDao {
	return &FinanceIndicatorDailyDao{
		group:   "default",
		table:   "hg_finance_indicator_daily",
		columns: financeIndicatorDailyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceIndicatorDailyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceIndicatorDailyDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceIndicatorDailyDao) Columns() FinanceIndicatorDailyColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceIndicatorDailyDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceIndicatorDailyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinanceIndicatorDailyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
