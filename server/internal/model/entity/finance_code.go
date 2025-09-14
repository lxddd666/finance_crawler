// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinanceCode is the golang structure for table finance_code.
type FinanceCode struct {
	Id           int64  `json:"id"           orm:"id"            description:"分类ID"`
	Code         string `json:"code"         orm:"code"          description:"代码"`
	Name         string `json:"name"         orm:"name"          description:"名称"`
	Exchange     string `json:"exchange"     orm:"exchange"      description:"交易所"`
	CompleteCode string `json:"completeCode" orm:"complete_code" description:"完整code"`
}
