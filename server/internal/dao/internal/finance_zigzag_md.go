// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceZigzagMdDao is the data access object for the table hg_finance_zigzag_md.
type FinanceZigzagMdDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  FinanceZigzagMdColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// FinanceZigzagMdColumns defines and stores column names for the table hg_finance_zigzag_md.
type FinanceZigzagMdColumns struct {
	Id        string // 分类ID
	Code      string // 代码
	Md5Slope  string // 5日均线所在zigzag斜率
	Md10Slope string // 10日均线所在zigzag斜率
	Md20Slope string // 20日均线所在zigzag斜率
	Md30Slope string // 30日均线所在zigzag斜率
	Md60Slope string // 60日均线所在zigzag斜率
	Day       string // 日期(yyyy-MM-dd)
}

// financeZigzagMdColumns holds the columns for the table hg_finance_zigzag_md.
var financeZigzagMdColumns = FinanceZigzagMdColumns{
	Id:        "id",
	Code:      "code",
	Md5Slope:  "md_5_slope",
	Md10Slope: "md_10_slope",
	Md20Slope: "md_20_slope",
	Md30Slope: "md_30_slope",
	Md60Slope: "md_60_slope",
	Day:       "day",
}

// NewFinanceZigzagMdDao creates and returns a new DAO object for table data access.
func NewFinanceZigzagMdDao(handlers ...gdb.ModelHandler) *FinanceZigzagMdDao {
	return &FinanceZigzagMdDao{
		group:    "default",
		table:    "hg_finance_zigzag_md",
		columns:  financeZigzagMdColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceZigzagMdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceZigzagMdDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceZigzagMdDao) Columns() FinanceZigzagMdColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceZigzagMdDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceZigzagMdDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceZigzagMdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
