// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FinanceBoll is the golang structure of table hg_finance_boll for DAO operations like Where/Data.
type FinanceBoll struct {
	g.Meta            `orm:"table:hg_finance_boll, do:true"`
	Id                interface{} // 分类ID
	Code              interface{} // code
	KlineId           interface{} // k线id
	KlineType         interface{} // k线类型
	Timestamp         interface{} // 时间戳
	MiddleBand        interface{} // 中轨（移动平均线）
	UpperBand         interface{} // 上轨
	LowerBand         interface{} // 下轨
	StandardDeviation interface{} // 标准差
	ClosePrice        interface{} // 收盘价
	KlineNum          interface{} // k线根数
	Multiple          interface{} // 标准差倍数
	CreatedAt         *gtime.Time // 创建时间
	Key               interface{} // timestamp和code组合
}
