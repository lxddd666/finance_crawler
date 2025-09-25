// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceDailyKline is the golang structure for table finance_daily_kline.
type FinanceDailyKline struct {
	Id         int64       `json:"id"         orm:"id"          description:"主键ID"`
	Code       string      `json:"code"       orm:"code"        description:"股票代码"`
	KlineType  int         `json:"klineType"  orm:"kline_type"  description:"K线类型: 0-普通 1-复权"`
	Timestamp  int64       `json:"timestamp"  orm:"timestamp"   description:"时间戳"`
	OpenPrice  float64     `json:"openPrice"  orm:"open_price"  description:"开盘价"`
	ClosePrice float64     `json:"closePrice" orm:"close_price" description:"收盘价"`
	HighPrice  float64     `json:"highPrice"  orm:"high_price"  description:"最高价"`
	LowPrice   float64     `json:"lowPrice"   orm:"low_price"   description:"最低价"`
	Volume     int64       `json:"volume"     orm:"volume"      description:"成交量"`
	Turnover   float64     `json:"turnover"   orm:"turnover"    description:"成交额"`
	Md5        float64     `json:"md5"        orm:"md_5"        description:"5日均线"`
	Md10       float64     `json:"md10"       orm:"md_10"       description:"10日均线"`
	Md20       float64     `json:"md20"       orm:"md_20"       description:"20日均线"`
	Md30       float64     `json:"md30"       orm:"md_30"       description:"30日均线"`
	Md60       float64     `json:"md60"       orm:"md_60"       description:"60日均线"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	Key        string      `json:"key"        orm:"key"         description:"唯一键: timestamp_code_scale"`
	Scale      int         `json:"scale"      orm:"scale"       description:"K线周期(分钟)"`
	Day        string      `json:"day"        orm:"day"         description:"日期(yyyy-MM-dd)"`
	Md50       float64     `json:"md50"       orm:"md_50"       description:"50日均线"`
}
