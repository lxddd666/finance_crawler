// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceCodeDao is the data access object for the table hg_finance_code.
type FinanceCodeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns FinanceCodeColumns // columns contains all the column names of Table for convenient usage.
}

// FinanceCodeColumns defines and stores column names for the table hg_finance_code.
type FinanceCodeColumns struct {
	Id           string // 分类ID
	Code         string // 代码
	Name         string // 名称
	Exchange     string // 交易所
	CompleteCode string // 完整code
}

// financeCodeColumns holds the columns for the table hg_finance_code.
var financeCodeColumns = FinanceCodeColumns{
	Id:           "id",
	Code:         "code",
	Name:         "name",
	Exchange:     "exchange",
	CompleteCode: "complete_code",
}

// NewFinanceCodeDao creates and returns a new DAO object for table data access.
func NewFinanceCodeDao() *FinanceCodeDao {
	return &FinanceCodeDao{
		group:   "default",
		table:   "hg_finance_code",
		columns: financeCodeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceCodeDao) Columns() FinanceCodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceCodeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinanceCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
