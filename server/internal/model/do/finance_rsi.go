// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceRsi is the golang structure of table hg_finance_rsi for DAO operations like Where/Data.
type FinanceRsi struct {
	g.Meta         `orm:"table:hg_finance_rsi, do:true"`
	Id             interface{} // 分类ID
	Code           interface{} // code
	Rsi            interface{} // 相对强弱指数的数值
	SmoothingMa    interface{} // 平滑移动平均线
	BollUpper      interface{} // 布林带上轨
	BollLower      interface{} // 布林带下轨
	IsBoll         interface{} // 是否触及布林带边界
	BullDivergence interface{} // 看涨背离信号
	BearDivergence interface{} // 看跌背离信号
	MaLength       interface{} // 参数
	CreatedAt      *gtime.Time // 创建时间
	Key            interface{} // timestamp和code组合:scalse
	Day            interface{} // 日期
	Scale          interface{} // 分钟一条k线
}
