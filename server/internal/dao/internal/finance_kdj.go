// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceKdjDao is the data access object for the table hg_finance_kdj.
type FinanceKdjDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of the current DAO.
	columns FinanceKdjColumns // columns contains all the column names of Table for convenient usage.
}

// FinanceKdjColumns defines and stores column names for the table hg_finance_kdj.
type FinanceKdjColumns struct {
	Id         string // 分类ID
	Code       string // code
	K          string // k值
	D          string // D值
	J          string // J值
	ClosePrice string // 收盘价
	HighPrice  string // 该K线最高价
	LowPrice   string // 该K线最低价
	Timestamp  string // 时间戳
	CreatedAt  string // 创建时间
	Key        string // timestamp和code组合
	Day        string // 日期
	Scale      string // 分钟一条k线
}

// financeKdjColumns holds the columns for the table hg_finance_kdj.
var financeKdjColumns = FinanceKdjColumns{
	Id:         "id",
	Code:       "code",
	K:          "k",
	D:          "d",
	J:          "j",
	ClosePrice: "close_price",
	HighPrice:  "high_price",
	LowPrice:   "low_price",
	Timestamp:  "timestamp",
	CreatedAt:  "created_at",
	Key:        "key",
	Day:        "day",
	Scale:      "scale",
}

// NewFinanceKdjDao creates and returns a new DAO object for table data access.
func NewFinanceKdjDao() *FinanceKdjDao {
	return &FinanceKdjDao{
		group:   "default",
		table:   "hg_finance_kdj",
		columns: financeKdjColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceKdjDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceKdjDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceKdjDao) Columns() FinanceKdjColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceKdjDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceKdjDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *FinanceKdjDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
