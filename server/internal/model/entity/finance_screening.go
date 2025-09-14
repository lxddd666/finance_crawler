// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceScreening is the golang structure for table finance_screening.
type FinanceScreening struct {
	Id         int64       `json:"id"         orm:"id"          description:"分类ID"`
	Code       string      `json:"code"       orm:"code"        description:"code"`
	Boll       int         `json:"boll"       orm:"boll"        description:"满足boll"`
	Macd       int         `json:"macd"       orm:"macd"        description:"满足macd"`
	Kdj        int         `json:"kdj"        orm:"kdj"         description:"满足kdj"`
	Rsi        int         `json:"rsi"        orm:"rsi"         description:"满足rsi"`
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"收盘价"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"   description:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	Key        string      `json:"key"        orm:"key"         description:"timestamp和code组合"`
	Day        string      `json:"day"        orm:"day"         description:"日期"`
	Scale      int         `json:"scale"      orm:"scale"       description:"分钟一条k线"`
	MatchCount int         `json:"matchCount" orm:"match_count" description:"符合条件数量"`
}
