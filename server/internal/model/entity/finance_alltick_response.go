// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceAlltickResponse is the golang structure for table finance_alltick_response.
type FinanceAlltickResponse struct {
	Id         int64       `json:"id"         orm:"id"          description:"分类ID"`
	Msg        string      `json:"msg"        orm:"msg"         description:""`
	Trace      string      `json:"trace"      orm:"trace"       description:""`
	Data       *gjson.Json `json:"data"       orm:"data"        description:"data"`
	Code       string      `json:"code"       orm:"code"        description:"code"`
	Sort       int         `json:"sort"       orm:"sort"        description:"排序"`
	KlineType  int         `json:"klineType"  orm:"kline_type"  description:"k线类型"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"   description:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  orm:"open_price"  description:"该K线开盘价"`
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"该K线收盘价"`
	HighPrice  float64     `json:"highPrice"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   orm:"low_price"   description:"该K线最低价"`
	Volume     float64     `json:"volume"     orm:"volume"      description:"该K线成交数量"`
	Turnover   float64     `json:"turnover"   orm:"turnover"    description:"该K线成交金额"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:"删除时间"`
}
