// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// StockIndicator is the golang structure of table hg_stock_indicator for DAO operations like Where/Data.
type StockIndicator struct {
	g.Meta    `orm:"table:hg_stock_indicator, do:true"`
	Id        interface{} // 分类ID
	CreatedAt *gtime.Time // 创建时间
}
