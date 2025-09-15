// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TgstatChannelCrawlerUrlDao is the data access object for the table hg_tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrlDao struct {
	table    string                         // table is the underlying table name of the DAO.
	group    string                         // group is the database configuration group name of the current DAO.
	columns  TgstatChannelCrawlerUrlColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler             // handlers for customized model modification.
}

// TgstatChannelCrawlerUrlColumns defines and stores column names for the table hg_tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrlColumns struct {
	Id        string // id
	Url       string // url
	Status    string // 采集状态0未开始1执行完成-1执行失败
	Type      string // 采集类型
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string //
	Comment   string //
}

// tgstatChannelCrawlerUrlColumns holds the columns for the table hg_tgstat_channel_crawler_url.
var tgstatChannelCrawlerUrlColumns = TgstatChannelCrawlerUrlColumns{
	Id:        "id",
	Url:       "url",
	Status:    "status",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Comment:   "comment",
}

// NewTgstatChannelCrawlerUrlDao creates and returns a new DAO object for table data access.
func NewTgstatChannelCrawlerUrlDao(handlers ...gdb.ModelHandler) *TgstatChannelCrawlerUrlDao {
	return &TgstatChannelCrawlerUrlDao{
		group:    "default",
		table:    "hg_tgstat_channel_crawler_url",
		columns:  tgstatChannelCrawlerUrlColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TgstatChannelCrawlerUrlDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TgstatChannelCrawlerUrlDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TgstatChannelCrawlerUrlDao) Columns() TgstatChannelCrawlerUrlColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TgstatChannelCrawlerUrlDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TgstatChannelCrawlerUrlDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TgstatChannelCrawlerUrlDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
