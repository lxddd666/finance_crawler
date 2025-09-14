// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceKdj is the golang structure for table finance_kdj.
type FinanceKdj struct {
	Id         int64       `json:"id"         orm:"id"          description:"分类ID"`
	Code       string      `json:"code"       orm:"code"        description:"code"`
	K          float64     `json:"k"          orm:"k"           description:"k值"`
	D          float64     `json:"d"          orm:"d"           description:"D值"`
	J          float64     `json:"j"          orm:"j"           description:"J值"`
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"收盘价"`
	HighPrice  float64     `json:"highPrice"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   float64     `json:"lowPrice"   orm:"low_price"   description:"该K线最低价"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"   description:"时间戳"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	Key        string      `json:"key"        orm:"key"         description:"timestamp和code组合"`
	Day        string      `json:"day"        orm:"day"         description:"日期"`
	Scale      int         `json:"scale"      orm:"scale"       description:"分钟一条k线"`
}
