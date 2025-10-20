// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceZigzagMd is the golang structure of table hg_finance_zigzag_md for DAO operations like Where/Data.
type FinanceZigzagMd struct {
	g.Meta    `orm:"table:hg_finance_zigzag_md, do:true"`
	Id        interface{} // 分类ID
	Code      interface{} // 代码
	Md5Slope  interface{} // 5日均线所在zigzag斜率
	Md10Slope interface{} // 10日均线所在zigzag斜率
	Md20Slope interface{} // 20日均线所在zigzag斜率
	Md30Slope interface{} // 30日均线所在zigzag斜率
	Md60Slope interface{} // 60日均线所在zigzag斜率
	Day       interface{} // 日期(yyyy-MM-dd)
}
