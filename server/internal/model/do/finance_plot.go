// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinancePlot is the golang structure of table hg_finance_plot for DAO operations like Where/Data.
type FinancePlot struct {
	g.Meta    `orm:"table:hg_finance_plot, do:true"`
	Id        interface{} // 分类ID
	Code      interface{} // code
	Indicator interface{} // 指标
	Day       interface{} // 日期
	Path      interface{} // path
}
