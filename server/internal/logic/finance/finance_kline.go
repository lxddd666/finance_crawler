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
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/logic/sina"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/httpReq"
	"hotgo/internal/service"
	"hotgo/utility/format"
	"log"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysFinanceKline struct{}

func NewSysFinanceKline() *sSysFinanceKline {
	return &sSysFinanceKline{}
}

func init() {
	service.RegisterSysFinanceKline(NewSysFinanceKline())
}

// Model k线ORM模型
func (s *sSysFinanceKline) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceKline.Ctx(ctx), option...)
}

// Kline k线
func (s *sSysFinanceKline) Kline(ctx context.Context, code, ma string, scale, datalen int, proxyFlag bool) (klineList []*entity.FinanceKline, err error) {
	var KlineList []*sina.SinaResult
	if !proxyFlag {
		KlineList, err = sina.GetKlineData(ctx, &httpReq.SinaHttpReq{
			Symbol:  code,
			Scale:   scale,
			Ma:      ma,
			Datalen: datalen,
		})
		if err != nil {
			return
		}
	} else {
		KlineList, err = sina.GetKlineDataProxy(ctx, &httpReq.SinaHttpReq{
			Symbol:  code,
			Scale:   scale,
			Ma:      ma,
			Datalen: datalen,
		})
		if err != nil {
			return
		}
	}

	for _, kline := range KlineList {
		klineList = append(klineList, &entity.FinanceKline{
			Code:       code,
			Timestamp:  format.DayStrToTimestamp(kline.Day),
			OpenPrice:  gconv.Float64(kline.Open),
			ClosePrice: gconv.Float64(kline.Close),
			HighPrice:  gconv.Float64(kline.High),
			LowPrice:   gconv.Float64(kline.Low),
			Volume:     gconv.Int64(kline.Volume),
			Scale:      scale,
			Day:        kline.Day,
			Key:        fmt.Sprintf("%s%s", code, kline.Day),
		})
	}
	if len(klineList) == 0 {
		err = gerror.New("获取数据为空，请查看是否ip被禁止")
		return
	}
	err = BatchInsertKline(ctx, klineList, code)
	return
}

func BatchInsertKline(ctx context.Context, klineList []*entity.FinanceKline, code string) error {
	batchSize := 300
	total := len(klineList)

	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := klineList[i:end]
		_, err := dao.FinanceKline.Ctx(ctx).InsertIgnore(batch)
		if err != nil {
			return fmt.Errorf("批量插入失败，批次 %d-%d: %v", i, end-1, err)
		}

		// 可选：添加日志输出进度
		log.Printf("已插入[%s] %d-%d 条数据，总共 %d 条", code, i, end-1, total)
	}

	return nil
}
