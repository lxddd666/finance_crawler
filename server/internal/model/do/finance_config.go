// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FinanceConfig is the golang structure of table hg_finance_config for DAO operations like Where/Data.
type FinanceConfig struct {
	g.Meta       `orm:"table:hg_finance_config, do:true"`
	Id           interface{} // 配置ID
	AlltickToken interface{} // 配置分组
}
