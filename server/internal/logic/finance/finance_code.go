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
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/xuri/excelize/v2"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
	"hotgo/utility/stock"
	"io"
	"mime/multipart"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysFinanceCode struct{}

func NewSysFinanceCode() *sSysFinanceCode {
	return &sSysFinanceCode{}
}

func init() {
	service.RegisterSysFinanceCode(NewSysFinanceCode())
}

// Model 股票代码ORM模型
func (s *sSysFinanceCode) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.FinanceCode.Ctx(ctx), option...)
}

// List 获取股票代码列表
func (s *sSysFinanceCode) List(ctx context.Context, in *sysin.FinanceCodeListInp) (list []*sysin.FinanceCodeListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.FinanceCodeListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.FinanceCode.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.FinanceCode.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取股票代码列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出股票代码
func (s *sSysFinanceCode) Export(ctx context.Context, in *sysin.FinanceCodeListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.FinanceCodeExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出股票代码-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.FinanceCodeExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增股票代码
func (s *sSysFinanceCode) Edit(ctx context.Context, in *sysin.FinanceCodeEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.FinanceCodeUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改股票代码失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.FinanceCodeInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增股票代码失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除股票代码
func (s *sSysFinanceCode) Delete(ctx context.Context, in *sysin.FinanceCodeDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除股票代码失败，请稍后重试！")
		return
	}
	return
}

// View 获取股票代码指定信息
func (s *sSysFinanceCode) View(ctx context.Context, in *sysin.FinanceCodeViewInp) (res *sysin.FinanceCodeViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取股票代码信息，请稍后重试！")
		return
	}
	return
}

// ImportCode 导入股票代码
func (s *sSysFinanceCode) ImportCode(ctx context.Context, inp sysin.FinanceImportCodeInp) (err error) {
	file, err := inp.File.Open()
	if err != nil {
		return gerror.Wrap(err, "获取上传文件失败，请检查文件格式")
	}
	defer func(file multipart.File) { _ = file.Close() }(file)
	fileBytes, err := io.ReadAll(file)
	data, err := parseExcelFile(fileBytes)
	list := make([]entity.FinanceCode, 0)
	for i, aStock := range data[0].Rows {
		if i == 0 {
			continue
		}
		stock := entity.FinanceCode{}
		if len(aStock) == 0 {
			break
		}
		stockCode := aStock[0].Value
		codeSplit := gstr.Split(stockCode, ".")
		stock.Code = codeSplit[0]
		stock.Exchange = codeSplit[1]
		stock.Name = aStock[1].Value
		list = append(list, stock)
	}
	_, err = s.Model(ctx).Insert(list)
	return
}

// 解析Excel文件
func parseExcelFile(fileBytes []byte) ([]SheetData, error) {
	// 从字节切片打开Excel文件
	f, err := excelize.OpenReader(&buffer{data: fileBytes})
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 获取所有工作表
	sheets := f.GetSheetList()
	var result []SheetData

	// 遍历每个工作表
	for _, sheet := range sheets {
		// 获取工作表中的所有行
		rows, err := f.GetRows(sheet)
		if err != nil {
			return nil, err
		}

		// 将数据转换为SheetData结构
		var sheetData SheetData
		sheetData.Name = sheet

		for rowIdx, row := range rows {
			var rowData []CellData
			for colIdx, colCell := range row {
				cellName, _ := excelize.CoordinatesToCellName(colIdx+1, rowIdx+1)
				rowData = append(rowData, CellData{
					Name:  cellName,
					Value: colCell,
				})
			}
			sheetData.Rows = append(sheetData.Rows, rowData)
		}

		result = append(result, sheetData)
	}

	return result, nil
}

// SheetData 表示工作表数据
type SheetData struct {
	Name string
	Rows [][]CellData
}

// CellData 表示单元格数据
type CellData struct {
	Name  string
	Value string
}

// 实现io.ReadSeeker接口用于excelize
type buffer struct {
	data []byte
	pos  int
}

func (b *buffer) Read(p []byte) (n int, err error) {
	if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	n = copy(p, b.data[b.pos:])
	b.pos += n
	return n, nil
}

func (b *buffer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		b.pos = int(offset)
	case io.SeekCurrent:
		b.pos += int(offset)
	case io.SeekEnd:
		b.pos = len(b.data) + int(offset)
	}
	return int64(b.pos), nil
}

func (b *buffer) Close() error {
	return nil
}

func (s *sSysFinanceCode) GetAllCode(ctx context.Context) (list []*entity.FinanceCode, err error) {
	// 获取每天boll
	codeList := make([]*entity.FinanceCode, 0)
	err = dao.FinanceKline.Ctx(ctx).Scan(&codeList)
	if err != nil {
		return
	}
	if len(codeList) == 0 {
		err = gerror.New("获取code失败")
		return
	}
	return
}

func (s *sSysFinanceCode) GetCodeKline(ctx context.Context, code string, KlineNum int) (list []*entity.FinanceKline, err error) {
	if KlineNum == 0 {
		KlineNum = 50
	}
	err = dao.FinanceKline.Ctx(ctx).Where(dao.FinanceKline.Columns().Code, code).OrderDesc(dao.FinanceKline.Columns().Day).Limit(KlineNum).Scan(&list)
	if err != nil {
		return
	}
	stock.ReverseKline(list)
	return
}
