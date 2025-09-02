// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceAlltickResponse is the golang structure of table hg_finance_alltick_response for DAO operations like Where/Data.
type FinanceAlltickResponse struct {
	g.Meta     `orm:"table:hg_finance_alltick_response, do:true"`
	Id         interface{} // 分类ID
	Msg        interface{} //
	Trace      interface{} //
	Data       *gjson.Json // data
	Code       interface{} // code
	Sort       interface{} // 排序
	KlineType  interface{} // k线类型
	Timestamp  interface{} // 时间戳
	OpenPrice  interface{} // 该K线开盘价
	ClosePrice interface{} // 该K线收盘价
	HighPrice  interface{} // 该K线最高价
	LowPrice   interface{} // 该K线最低价
	Volume     interface{} // 该K线成交数量
	Turnover   interface{} // 该K线成交金额
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
	DeletedAt  *gtime.Time // 删除时间
}
