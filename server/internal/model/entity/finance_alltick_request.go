// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceAlltickRequest is the golang structure for table finance_alltick_request.
type FinanceAlltickRequest struct {
	Id                int64       `json:"id"                orm:"id"                  description:"主键"`
	Code              string      `json:"code"              orm:"code"                description:"股票代码"`
	KlineType         int         `json:"klineType"         orm:"kline_type"          description:"k线类型 1、1是1分钟K，2是5分钟K，3是15分钟K，4是30分钟K，5是小时K，6是2小时K(股票不支持2小时)，7是4小时K(股票不支持4小时)，8是日K，9是周K，10是月K （注：股票不支持2小时K、4小时K）"`
	KlineTimestampEnd int         `json:"klineTimestampEnd" orm:"kline_timestamp_end" description:"从指定时间往前查询K线 1、传0表示从当前最新的交易日往前查k线 2、指定时间请传时间戳，传时间戳表示从该时间戳往前查k线 3、只有外汇贵金属加密货币支持传时间戳，股票类的code不支持"`
	AdjustType        int         `json:"adjustType"        orm:"adjust_type"         description:"复权类型,对于股票类的code才有效，例如：0:除权,1:前复权，目前仅支持0"`
	QueryKlineNum     int         `json:"queryKlineNum"     orm:"query_kline_num"     description:"表示多少根K线，每次最多500根"`
	SymbolList        *gjson.Json `json:"symbolList"        orm:"symbol_list"         description:"code参数list"`
	Status            int         `json:"status"            orm:"status"              description:"状态(1未验证,2已验证)"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:"更新时间"`
}
