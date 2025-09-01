// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AlltickStockModelDao is the data access object for the table hg_alltick_stock_model.
type AlltickStockModelDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  AlltickStockModelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// AlltickStockModelColumns defines and stores column names for the table hg_alltick_stock_model.
type AlltickStockModelColumns struct {
	Id                string // 主键
	Code              string // 股票代码
	KlineType         string // k线类型 1、1是1分钟K，2是5分钟K，3是15分钟K，4是30分钟K，5是小时K，6是2小时K(股票不支持2小时)，7是4小时K(股票不支持4小时)，8是日K，9是周K，10是月K （注：股票不支持2小时K、4小时K）
	KlineTimestampEnd string // 从指定时间往前查询K线 1、传0表示从当前最新的交易日往前查k线 2、指定时间请传时间戳，传时间戳表示从该时间戳往前查k线 3、只有外汇贵金属加密货币支持传时间戳，股票类的code不支持
	AdjustType        string // 复权类型,对于股票类的code才有效，例如：0:除权,1:前复权，目前仅支持0
	QueryKlineNum     string // 表示多少根K线，每次最多500根
	SymbolList        string // code参数list
	Status            string // 状态(1未验证,2已验证)
	CreatedAt         string // 创建时间
	UpdatedAt         string // 更新时间
}

// alltickStockModelColumns holds the columns for the table hg_alltick_stock_model.
var alltickStockModelColumns = AlltickStockModelColumns{
	Id:                "id",
	Code:              "code",
	KlineType:         "kline_type",
	KlineTimestampEnd: "kline_timestamp_end",
	AdjustType:        "adjust_type",
	QueryKlineNum:     "query_kline_num",
	SymbolList:        "symbol_list",
	Status:            "status",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
}

// NewAlltickStockModelDao creates and returns a new DAO object for table data access.
func NewAlltickStockModelDao(handlers ...gdb.ModelHandler) *AlltickStockModelDao {
	return &AlltickStockModelDao{
		group:    "default",
		table:    "hg_alltick_stock_model",
		columns:  alltickStockModelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AlltickStockModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AlltickStockModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AlltickStockModelDao) Columns() AlltickStockModelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AlltickStockModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AlltickStockModelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AlltickStockModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
