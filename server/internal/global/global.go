// Package global
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package global

import "hotgo/internal/model/entity"

var (
	FinanceConfig *entity.FinanceConfig
	Day           string
	Timestamp     int64
	ProxyList     *SafeProxyList
)
