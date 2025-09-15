// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceBollDao is the data access object for the table hg_finance_boll.
type FinanceBollDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FinanceBollColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FinanceBollColumns defines and stores column names for the table hg_finance_boll.
type FinanceBollColumns struct {
	Id                string // 分类ID
	Code              string // code
	KlineId           string // k线id
	KlineType         string // k线类型
	Timestamp         string // 时间戳
	MiddleBand        string // 中轨（移动平均线）
	UpperBand         string // 上轨
	LowerBand         string // 下轨
	StandardDeviation string // 标准差
	ClosePrice        string // 收盘价
	KlineNum          string // k线根数
	Multiple          string // 标准差倍数
	CreatedAt         string // 创建时间
	Key               string // timestamp和code组合
	Day               string // 日期
	Scale             string // 分钟一条k线
}

// financeBollColumns holds the columns for the table hg_finance_boll.
var financeBollColumns = FinanceBollColumns{
	Id:                "id",
	Code:              "code",
	KlineId:           "kline_id",
	KlineType:         "kline_type",
	Timestamp:         "timestamp",
	MiddleBand:        "middle_band",
	UpperBand:         "upper_band",
	LowerBand:         "lower_band",
	StandardDeviation: "standard_deviation",
	ClosePrice:        "close_price",
	KlineNum:          "kline_num",
	Multiple:          "multiple",
	CreatedAt:         "created_at",
	Key:               "key",
	Day:               "day",
	Scale:             "scale",
}

// NewFinanceBollDao creates and returns a new DAO object for table data access.
func NewFinanceBollDao(handlers ...gdb.ModelHandler) *FinanceBollDao {
	return &FinanceBollDao{
		group:    "default",
		table:    "hg_finance_boll",
		columns:  financeBollColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceBollDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceBollDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceBollDao) Columns() FinanceBollColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceBollDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceBollDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceBollDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
