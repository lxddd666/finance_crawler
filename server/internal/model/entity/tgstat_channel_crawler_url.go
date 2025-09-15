// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TgstatChannelCrawlerUrl is the golang structure for table tgstat_channel_crawler_url.
type TgstatChannelCrawlerUrl struct {
	Id        int64       `json:"id"        orm:"id"         description:"id"`
	Url       string      `json:"url"       orm:"url"        description:"url"`
	Status    int         `json:"status"    orm:"status"     description:"采集状态0未开始1执行完成-1执行失败"`
	Type      string      `json:"type"      orm:"type"       description:"采集类型"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
	Comment   string      `json:"comment"   orm:"comment"    description:""`
}
