// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceIndicatorDaily is the golang structure of table hg_finance_indicator_daily for DAO operations like Where/Data.
type FinanceIndicatorDaily struct {
	g.Meta    `orm:"table:hg_finance_indicator_daily, do:true"`
	Id        interface{} // 分类ID
	Code      interface{} // 代码
	Name      interface{} // 名称
	Exchange  interface{} // 交易所
	Day       interface{} // 日期
	Timestamp interface{} // 时间戳
	Status    interface{} // 0未开始 1完成 -1失败
}
