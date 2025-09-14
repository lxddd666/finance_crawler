// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceKlineDao is the data access object for the table hg_finance_kline.
type FinanceKlineDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  FinanceKlineColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// FinanceKlineColumns defines and stores column names for the table hg_finance_kline.
type FinanceKlineColumns struct {
	Id         string // 分类ID
	Code       string // code
	KlineType  string // k线类型
	Timestamp  string // 时间戳
	OpenPrice  string // 该K线开盘价
	ClosePrice string // 该K线收盘价
	HighPrice  string // 该K线最高价
	LowPrice   string // 该K线最低价
	Volume     string // 该K线成交数量
	Turnover   string // 该K线成交金额
	CreatedAt  string // 创建时间
	Key        string // 唯一key，timestamp和code组合
	Scale      string // 多少分钟一根K线
	Day        string // 日期
}

// financeKlineColumns holds the columns for the table hg_finance_kline.
var financeKlineColumns = FinanceKlineColumns{
	Id:         "id",
	Code:       "code",
	KlineType:  "kline_type",
	Timestamp:  "timestamp",
	OpenPrice:  "open_price",
	ClosePrice: "close_price",
	HighPrice:  "high_price",
	LowPrice:   "low_price",
	Volume:     "volume",
	Turnover:   "turnover",
	CreatedAt:  "created_at",
	Key:        "key",
	Scale:      "scale",
	Day:        "day",
}

// NewFinanceKlineDao creates and returns a new DAO object for table data access.
func NewFinanceKlineDao(handlers ...gdb.ModelHandler) *FinanceKlineDao {
	return &FinanceKlineDao{
		group:    "default",
		table:    "hg_finance_kline",
		columns:  financeKlineColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceKlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceKlineDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceKlineDao) Columns() FinanceKlineColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceKlineDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceKlineDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceKlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
