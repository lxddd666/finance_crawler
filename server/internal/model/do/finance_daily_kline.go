// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceDailyKline is the golang structure of table hg_finance_daily_kline for DAO operations like Where/Data.
type FinanceDailyKline struct {
	g.Meta     `orm:"table:hg_finance_daily_kline, do:true"`
	Id         interface{} // 主键ID
	Code       interface{} // 股票代码
	KlineType  interface{} // K线类型: 0-普通 1-复权
	Timestamp  interface{} // 时间戳
	OpenPrice  interface{} // 开盘价
	ClosePrice interface{} // 收盘价
	HighPrice  interface{} // 最高价
	LowPrice   interface{} // 最低价
	Volume     interface{} // 成交量
	Turnover   interface{} // 成交额
	Md5        interface{} // 5日均线
	Md10       interface{} // 10日均线
	Md20       interface{} // 20日均线
	Md30       interface{} // 30日均线
	Md60       interface{} // 60日均线
	CreatedAt  *gtime.Time // 创建时间
	UniqueKey  interface{} // 唯一键: timestamp_code_scale
	Scale      interface{} // K线周期(分钟)
	Day        interface{} // 日期(yyyy-MM-dd)
}
