// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceScreening is the golang structure of table hg_finance_screening for DAO operations like Where/Data.
type FinanceScreening struct {
	g.Meta     `orm:"table:hg_finance_screening, do:true"`
	Id         interface{} // 分类ID
	Code       interface{} // code
	Boll       interface{} // 满足boll
	Macd       interface{} // 满足macd
	Kdj        interface{} // 满足kdj
	Rsi        interface{} // 满足rsi
	ClosePrice interface{} // 收盘价
	Timestamp  interface{} // 时间戳
	CreatedAt  *gtime.Time // 创建时间
	Key        interface{} // timestamp和code组合
	Day        interface{} // 日期
	Scale      interface{} // 分钟一条k线
	MatchCount interface{} // 符合条件数量
}
