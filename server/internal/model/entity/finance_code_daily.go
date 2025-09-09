// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinanceCodeDaily is the golang structure for table finance_code_daily.
type FinanceCodeDaily struct {
	Id        int64  `json:"id"        orm:"id"        description:"分类ID"`
	Code      string `json:"code"      orm:"code"      description:"代码"`
	Name      string `json:"name"      orm:"name"      description:"名称"`
	Exchange  string `json:"exchange"  orm:"exchange"  description:"交易所"`
	Day       string `json:"day"       orm:"day"       description:"日期"`
	Timestamp int64  `json:"timestamp" orm:"timestamp" description:"时间戳"`
	Status    int    `json:"status"    orm:"status"    description:"0未开始 1完成 -1失败"`
}
