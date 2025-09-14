// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceKdj is the golang structure of table hg_finance_kdj for DAO operations like Where/Data.
type FinanceKdj struct {
	g.Meta     `orm:"table:hg_finance_kdj, do:true"`
	Id         interface{} // 分类ID
	Code       interface{} // code
	K          interface{} // k值
	D          interface{} // D值
	J          interface{} // J值
	ClosePrice interface{} // 收盘价
	HighPrice  interface{} // 该K线最高价
	LowPrice   interface{} // 该K线最低价
	Timestamp  interface{} // 时间戳
	CreatedAt  *gtime.Time // 创建时间
	Key        interface{} // timestamp和code组合
	Day        interface{} // 日期
	Scale      interface{} // 分钟一条k线
}
