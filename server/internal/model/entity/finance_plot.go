// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinancePlot is the golang structure for table finance_plot.
type FinancePlot struct {
	Id        int64  `json:"id"        orm:"id"        description:"分类ID"`
	Code      string `json:"code"      orm:"code"      description:"code"`
	Indicator string `json:"indicator" orm:"indicator" description:"指标"`
	Day       string `json:"day"       orm:"day"       description:"日期"`
	Path      string `json:"path"      orm:"path"      description:"path"`
}
