// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceBoll is the golang structure for table finance_boll.
type FinanceBoll struct {
	Id                int64       `json:"id"                orm:"id"                 description:"分类ID"`
	Code              string      `json:"code"              orm:"code"               description:"code"`
	KlineId           int64       `json:"klineId"           orm:"kline_id"           description:"k线id"`
	KlineType         int         `json:"klineType"         orm:"kline_type"         description:"k线类型"`
	Timestamp         int64       `json:"timestamp"         orm:"timestamp"          description:"时间戳"`
	MiddleBand        float64     `json:"middleBand"        orm:"middle_band"        description:"中轨（移动平均线）"`
	UpperBand         float64     `json:"upperBand"         orm:"upper_band"         description:"上轨"`
	LowerBand         float64     `json:"lowerBand"         orm:"lower_band"         description:"下轨"`
	StandardDeviation float64     `json:"standardDeviation" orm:"standard_deviation" description:"标准差"`
	ClosePrice        float64     `json:"closePrice"        orm:"close_price"        description:"收盘价"`
	KlineNum          int         `json:"klineNum"          orm:"kline_num"          description:"k线根数"`
	Multiple          int         `json:"multiple"          orm:"multiple"           description:"标准差倍数"`
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"         description:"创建时间"`
}
