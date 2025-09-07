// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceKline is the golang structure of table hg_finance_kline for DAO operations like Where/Data.
type FinanceKline struct {
	g.Meta     `orm:"table:hg_finance_kline, do:true"`
	Id         interface{} // 分类ID
	Code       interface{} // code
	KlineType  interface{} // k线类型
	Timestamp  interface{} // 时间戳
	OpenPrice  interface{} // 该K线开盘价
	ClosePrice interface{} // 该K线收盘价
	HighPrice  interface{} // 该K线最高价
	LowPrice   interface{} // 该K线最低价
	Volume     interface{} // 该K线成交数量
	Turnover   interface{} // 该K线成交金额
	CreatedAt  *gtime.Time // 创建时间
	Key        interface{} // 唯一key，timestamp和code组合
	Scale      interface{} // 多少分钟一根K线
	Day        interface{} // 日期
}
