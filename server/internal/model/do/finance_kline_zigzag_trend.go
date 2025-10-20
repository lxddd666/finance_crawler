// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceKlineZigzagTrend is the golang structure of table hg_finance_kline_zigzag_trend for DAO operations like Where/Data.
type FinanceKlineZigzagTrend struct {
	g.Meta           `orm:"table:hg_finance_kline_zigzag_trend, do:true"`
	Id               interface{} // 主键ID
	Code             interface{} // 股票代码
	KlineId          interface{} // K线id
	Key              interface{} // 唯一键: timestamp_code_scale
	MinChangePercent interface{} // 最小浮动点
	Day              interface{} // 日期(yyyy-MM-dd)
}
