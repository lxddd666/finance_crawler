// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceRsi is the golang structure for table finance_rsi.
type FinanceRsi struct {
	Id             int64       `json:"id"             orm:"id"              description:"分类ID"`
	Code           string      `json:"code"           orm:"code"            description:"code"`
	Rsi            float64     `json:"rsi"            orm:"rsi"             description:"相对强弱指数的数值"`
	SmoothingMa    float64     `json:"smoothingMa"    orm:"smoothing_ma"    description:"平滑移动平均线"`
	BollUpper      int         `json:"bollUpper"      orm:"boll_upper"      description:"布林带上轨"`
	BollLower      int         `json:"bollLower"      orm:"boll_lower"      description:"布林带下轨"`
	IsBoll         int         `json:"isBoll"         orm:"is_boll"         description:"是否触及布林带边界"`
	BullDivergence int         `json:"bullDivergence" orm:"bull_divergence" description:"看涨背离信号"`
	BearDivergence int         `json:"bearDivergence" orm:"bear_divergence" description:"看跌背离信号"`
	MaLength       int         `json:"maLength"       orm:"ma_length"       description:"参数"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	Key            string      `json:"key"            orm:"key"             description:"timestamp和code组合:scalse"`
	Day            string      `json:"day"            orm:"day"             description:"日期"`
	Scale          int         `json:"scale"          orm:"scale"           description:"分钟一条k线"`
}
