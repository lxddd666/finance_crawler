// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinanceZigzagMd is the golang structure for table finance_zigzag_md.
type FinanceZigzagMd struct {
	Id        int64   `json:"id"        orm:"id"          description:"分类ID"`
	Code      string  `json:"code"      orm:"code"        description:"代码"`
	Md5Slope  float64 `json:"md5Slope"  orm:"md_5_slope"  description:"5日均线所在zigzag斜率"`
	Md10Slope float64 `json:"md10Slope" orm:"md_10_slope" description:"10日均线所在zigzag斜率"`
	Md20Slope float64 `json:"md20Slope" orm:"md_20_slope" description:"20日均线所在zigzag斜率"`
	Md30Slope float64 `json:"md30Slope" orm:"md_30_slope" description:"30日均线所在zigzag斜率"`
	Md60Slope float64 `json:"md60Slope" orm:"md_60_slope" description:"60日均线所在zigzag斜率"`
	Day       string  `json:"day"       orm:"day"         description:"日期(yyyy-MM-dd)"`
}
