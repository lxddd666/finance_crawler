// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceCode is the golang structure of table hg_finance_code for DAO operations like Where/Data.
type FinanceCode struct {
	g.Meta       `orm:"table:hg_finance_code, do:true"`
	Id           interface{} // 分类ID
	Code         interface{} // 代码
	Name         interface{} // 名称
	Exchange     interface{} // 交易所
	CompleteCode interface{} // 完整code
}
