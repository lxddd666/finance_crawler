// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// StockIndicator is the golang structure for table stock_indicator.
type StockIndicator struct {
	Id        int64       `json:"id"        orm:"id"         description:"分类ID"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
}
