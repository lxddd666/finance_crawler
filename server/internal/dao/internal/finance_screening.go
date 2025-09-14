// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceScreeningDao is the data access object for the table hg_finance_screening.
type FinanceScreeningDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  FinanceScreeningColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// FinanceScreeningColumns defines and stores column names for the table hg_finance_screening.
type FinanceScreeningColumns struct {
	Id         string // 分类ID
	Code       string // code
	Boll       string // 满足boll
	Macd       string // 满足macd
	Kdj        string // 满足kdj
	Rsi        string // 满足rsi
	ClosePrice string // 收盘价
	Timestamp  string // 时间戳
	CreatedAt  string // 创建时间
	Key        string // timestamp和code组合
	Day        string // 日期
	Scale      string // 分钟一条k线
	MatchCount string // 符合条件数量
}

// financeScreeningColumns holds the columns for the table hg_finance_screening.
var financeScreeningColumns = FinanceScreeningColumns{
	Id:         "id",
	Code:       "code",
	Boll:       "boll",
	Macd:       "macd",
	Kdj:        "kdj",
	Rsi:        "rsi",
	ClosePrice: "close_price",
	Timestamp:  "timestamp",
	CreatedAt:  "created_at",
	Key:        "key",
	Day:        "day",
	Scale:      "scale",
	MatchCount: "match_count",
}

// NewFinanceScreeningDao creates and returns a new DAO object for table data access.
func NewFinanceScreeningDao(handlers ...gdb.ModelHandler) *FinanceScreeningDao {
	return &FinanceScreeningDao{
		group:    "default",
		table:    "hg_finance_screening",
		columns:  financeScreeningColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceScreeningDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceScreeningDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceScreeningDao) Columns() FinanceScreeningColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceScreeningDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceScreeningDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceScreeningDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
