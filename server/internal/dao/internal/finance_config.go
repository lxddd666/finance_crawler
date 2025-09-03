// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceConfigDao is the data access object for table hg_finance_config.
type FinanceConfigDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns FinanceConfigColumns // columns contains all the column names of Table for convenient usage.
}

// FinanceConfigColumns defines and stores column names for table hg_finance_config.
type FinanceConfigColumns struct {
	Id           string // 配置ID
	AlltickToken string // 配置分组
}

// financeConfigColumns holds the columns for table hg_finance_config.
var financeConfigColumns = FinanceConfigColumns{
	Id:           "id",
	AlltickToken: "alltick_token",
}

// NewFinanceConfigDao creates and returns a new DAO object for table data access.
func NewFinanceConfigDao() *FinanceConfigDao {
	return &FinanceConfigDao{
		group:   "default",
		table:   "hg_finance_config",
		columns: financeConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FinanceConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FinanceConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FinanceConfigDao) Columns() FinanceConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FinanceConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FinanceConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FinanceConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
