// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceDailyKlineDao is the data access object for the table hg_finance_daily_kline.
type FinanceDailyKlineDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  FinanceDailyKlineColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// FinanceDailyKlineColumns defines and stores column names for the table hg_finance_daily_kline.
type FinanceDailyKlineColumns struct {
	Id         string // 主键ID
	Code       string // 股票代码
	KlineType  string // K线类型: 0-普通 1-复权
	Timestamp  string // 时间戳
	OpenPrice  string // 开盘价
	ClosePrice string // 收盘价
	HighPrice  string // 最高价
	LowPrice   string // 最低价
	Volume     string // 成交量
	Turnover   string // 成交额
	Md5        string // 5日均线
	Md10       string // 10日均线
	Md20       string // 20日均线
	Md30       string // 30日均线
	Md60       string // 60日均线
	CreatedAt  string // 创建时间
	Key        string // 唯一键: timestamp_code_scale
	Scale      string // K线周期(分钟)
	Day        string // 日期(yyyy-MM-dd)
	Md50       string // 50日均线
}

// financeDailyKlineColumns holds the columns for the table hg_finance_daily_kline.
var financeDailyKlineColumns = FinanceDailyKlineColumns{
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
	Md5:        "md_5",
	Md10:       "md_10",
	Md20:       "md_20",
	Md30:       "md_30",
	Md60:       "md_60",
	CreatedAt:  "created_at",
	Key:        "key",
	Scale:      "scale",
	Day:        "day",
	Md50:       "md_50",
}

// NewFinanceDailyKlineDao creates and returns a new DAO object for table data access.
func NewFinanceDailyKlineDao(handlers ...gdb.ModelHandler) *FinanceDailyKlineDao {
	return &FinanceDailyKlineDao{
		group:    "default",
		table:    "hg_finance_daily_kline",
		columns:  financeDailyKlineColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FinanceDailyKlineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FinanceDailyKlineDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FinanceDailyKlineDao) Columns() FinanceDailyKlineColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FinanceDailyKlineDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FinanceDailyKlineDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FinanceDailyKlineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
