// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// FinanceKlineZigzagTrend is the golang structure for table finance_kline_zigzag_trend.
type FinanceKlineZigzagTrend struct {
	Id               int64   `json:"id"               orm:"id"                 description:"主键ID"`
	Code             string  `json:"code"             orm:"code"               description:"股票代码"`
	KlineId          int64   `json:"klineId"          orm:"kline_id"           description:"K线id"`
	Key              string  `json:"key"              orm:"key"                description:"唯一键: timestamp_code_scale"`
	MinChangePercent float64 `json:"minChangePercent" orm:"min_change_percent" description:"最小浮动点"`
	Day              string  `json:"day"              orm:"day"                description:"日期(yyyy-MM-dd)"`
}
