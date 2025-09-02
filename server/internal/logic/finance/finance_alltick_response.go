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
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysFinanceAlltickResponse struct{}

func NewSysFinanceAlltickResponse() *sSysFinanceAlltickResponse {
	return &sSysFinanceAlltickResponse{}
}

func init() {
	service.RegisterSysFinanceAlltickResponse(NewSysFinanceAlltickResponse())
}

// Model alltick返回值ORM模型
func (s *sSysFinanceAlltickResponse) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceAlltickResponse.Ctx(ctx), option...)
}

// List 获取alltick返回值列表
func (s *sSysFinanceAlltickResponse) List(ctx context.Context, in *sysin.FinanceAlltickResponseListInp) (list []*sysin.FinanceAlltickResponseListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.FinanceAlltickResponseListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.FinanceAlltickResponse.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.FinanceAlltickResponse.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.FinanceAlltickResponse.Columns().Sort).OrderDesc(dao.FinanceAlltickResponse.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取alltick返回值列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出alltick返回值
func (s *sSysFinanceAlltickResponse) Export(ctx context.Context, in *sysin.FinanceAlltickResponseListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.FinanceAlltickResponseExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出alltick返回值-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.FinanceAlltickResponseExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增alltick返回值
func (s *sSysFinanceAlltickResponse) Edit(ctx context.Context, in *sysin.FinanceAlltickResponseEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.FinanceAlltickResponseUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改alltick返回值失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.FinanceAlltickResponseInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增alltick返回值失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除alltick返回值
func (s *sSysFinanceAlltickResponse) Delete(ctx context.Context, in *sysin.FinanceAlltickResponseDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除alltick返回值失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取alltick返回值最大排序
func (s *sSysFinanceAlltickResponse) MaxSort(ctx context.Context, in *sysin.FinanceAlltickResponseMaxSortInp) (res *sysin.FinanceAlltickResponseMaxSortModel, err error) {
	if err = dao.FinanceAlltickResponse.Ctx(ctx).Fields(dao.FinanceAlltickResponse.Columns().Sort).OrderDesc(dao.FinanceAlltickResponse.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取alltick返回值最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.FinanceAlltickResponseMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取alltick返回值指定信息
func (s *sSysFinanceAlltickResponse) View(ctx context.Context, in *sysin.FinanceAlltickResponseViewInp) (res *sysin.FinanceAlltickResponseViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取alltick返回值信息，请稍后重试！")
		return
	}
	return
}
