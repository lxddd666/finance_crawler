// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceMacd is the golang structure for table finance_macd.
type FinanceMacd struct {
	Id           int64       `json:"id"           orm:"id"            description:"分类ID"`
	Code         string      `json:"code"         orm:"code"          description:"code"`
	FastPeriod   int         `json:"fastPeriod"   orm:"fast_period"   description:"快线周期"`
	SlowPeriod   int         `json:"slowPeriod"   orm:"slow_period"   description:"慢线周期"`
	SignalPeriod int         `json:"signalPeriod" orm:"signal_period" description:"信号周期"`
	Dif          float64     `json:"dif"          orm:"dif"           description:"EMA12 - EMA26 快线"`
	Dea          float64     `json:"dea"          orm:"dea"           description:"DIF的9日EMA（26+9-1)慢线"`
	Macd         float64     `json:"macd"         orm:"macd"          description:"macd柱子"`
	Timestamp    int64       `json:"timestamp"    orm:"timestamp"     description:"时间戳"`
	ClosePrice   float64     `json:"closePrice"   orm:"close_price"   description:"收盘价"`
	KlineNum     int         `json:"klineNum"     orm:"kline_num"     description:"k线根数"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	Key          string      `json:"key"          orm:"key"           description:"timestamp和code组合"`
	Day          string      `json:"day"          orm:"day"           description:"日期"`
	Scale        int         `json:"scale"        orm:"scale"         description:"分钟一条k线"`
}
