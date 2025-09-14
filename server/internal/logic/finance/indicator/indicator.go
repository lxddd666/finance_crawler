// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/service"
)

type sSysStockIndicator struct{}

func NewSysStockIndicator() *sSysStockIndicator {
	return &sSysStockIndicator{}
}

func init() {
	service.RegisterSysStockIndicator(NewSysStockIndicator())
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

func (s *sSysStockIndicator) Screening(ctx context.Context) {
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)
	if err != nil {
		return
	}
	for _, code := range codeList {
		// 获取 满足boll
		fmt.Println(code)
	}
}
