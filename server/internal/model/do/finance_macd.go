// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceMacd is the golang structure of table hg_finance_macd for DAO operations like Where/Data.
type FinanceMacd struct {
	g.Meta       `orm:"table:hg_finance_macd, do:true"`
	Id           interface{} // 分类ID
	Code         interface{} // code
	FastPeriod   interface{} // 快线周期
	SlowPeriod   interface{} // 慢线周期
	SignalPeriod interface{} // 信号周期
	Dif          interface{} // EMA12 - EMA26 快线
	Dea          interface{} // DIF的9日EMA（26+9-1)慢线
	Macd         interface{} // macd柱子
	Timestamp    interface{} // 时间戳
	ClosePrice   interface{} // 收盘价
	KlineNum     interface{} // k线根数
	CreatedAt    *gtime.Time // 创建时间
	Key          interface{} // timestamp和code组合
	Day          interface{} // 日期
	Scale        interface{} // 分钟一条k线
}
