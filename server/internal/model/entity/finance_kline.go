// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceKline is the golang structure for table finance_kline.
type FinanceKline struct {
	Id         int64       `json:"id"         orm:"id"          description:"分类ID"`
	Code       string      `json:"code"       orm:"code"        description:"code"`
	KlineType  int         `json:"klineType"  orm:"kline_type"  description:"k线类型"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"   description:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  orm:"open_price"  description:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   orm:"low_price"   description:"该K线最低价"`
	Volume     int64       `json:"volume"     orm:"volume"      description:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   orm:"turnover"    description:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	Key        string      `json:"key"        orm:"key"         description:"唯一key，timestamp和code组合"`
	Scale      int         `json:"scale"      orm:"scale"       description:"多少分钟一根K线"`
	Day        string      `json:"day"        orm:"day"         description:"日期"`
}
