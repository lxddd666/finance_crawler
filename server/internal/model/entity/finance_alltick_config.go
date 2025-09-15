// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinanceAlltickConfig is the golang structure for table finance_alltick_config.
type FinanceAlltickConfig struct {
	Id           int64  `json:"id"           orm:"id"            description:"配置ID"`
	AlltickToken string `json:"alltickToken" orm:"alltick_token" description:"配置分组"`
}
