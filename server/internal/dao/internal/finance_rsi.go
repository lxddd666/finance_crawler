// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceRsiDao is the data access object for the table hg_finance_rsi.
type FinanceRsiDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FinanceRsiColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FinanceRsiColumns defines and stores column names for the table hg_finance_rsi.
type FinanceRsiColumns struct {
	Id             string // 分类ID
	Code           string // code
	Rsi            string // 相对强弱指数的数值
	SmoothingMa    string // 平滑移动平均线
	BollUpper      string // 布林带上轨
	BollLower      string // 布林带下轨
	IsBoll         string // 是否触及布林带边界
	BullDivergence string // 看涨背离信号
	BearDivergence string // 看跌背离信号
	MaLength       string // 参数
	CreatedAt      string // 创建时间
	Key            string // timestamp和code组合:scalse
	Day            string // 日期
	Scale          string // 分钟一条k线
}

// financeRsiColumns holds the columns for the table hg_finance_rsi.
var financeRsiColumns = FinanceRsiColumns{
	Id:             "id",
	Code:           "code",
	Rsi:            "rsi",
	SmoothingMa:    "smoothing_ma",
	BollUpper:      "boll_upper",
	BollLower:      "boll_lower",
	IsBoll:         "is_boll",
	BullDivergence: "bull_divergence",
	BearDivergence: "bear_divergence",
	MaLength:       "ma_length",
	CreatedAt:      "created_at",
	Key:            "key",
	Day:            "day",
	Scale:          "scale",
}

// NewFinanceRsiDao creates and returns a new DAO object for table data access.
func NewFinanceRsiDao(handlers ...gdb.ModelHandler) *FinanceRsiDao {
	return &FinanceRsiDao{
		group:    "default",
		table:    "hg_finance_rsi",
		columns:  financeRsiColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceRsiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceRsiDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceRsiDao) Columns() FinanceRsiColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceRsiDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceRsiDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceRsiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
