// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2025 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.17.8
package sys

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/do"
	"hotgo/internal/model/entity"
	sysin "hotgo/internal/model/input/financein"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"
)

type sSysTestFinance struct{}

// FinanceAlltickRequestData 用于JSON序列化的数据结构
type FinanceAlltickRequestData struct {
	Code              string `json:"code"`
	KlineType         int    `json:"kline_type"`
	KlineTimestampEnd int    `json:"kline_timestamp_end"`
	AdjustType        int    `json:"adjust_type"`
	QueryKlineNum     int    `json:"query_kline_num"`
}

// FinanceAlltickRequestQuery 用于JSON序列化的完整查询结构
type FinanceAlltickRequestQuery struct {
	Trace string                    `json:"trace"`
	Data  FinanceAlltickRequestData `json:"data"`
}

func NewSysTestFinance() *sSysTestFinance {
	return &sSysTestFinance{}
}

func init() {
	service.RegisterSysTestFinance(NewSysTestFinance())
}

// ConvertToQueryString 将 FinanceAlltickRequest 结构体转换为 JSON 查询字符串
func (s *sSysTestFinance) ConvertToQueryString(req *entity.FinanceAlltickRequest) (string, error) {
	// 如果 trace 为空，使用默认值

	trace := "1111111111111111111111111"

	query := FinanceAlltickRequestQuery{
		Trace: trace,
		Data: FinanceAlltickRequestData{
			Code:              req.Code,
			KlineType:         req.KlineType,
			KlineTimestampEnd: req.KlineTimestampEnd,
			AdjustType:        req.AdjustType,
			QueryKlineNum:     req.QueryKlineNum,
		},
	}

	jsonBytes, err := json.Marshal(query)
	if err != nil {
		return "", fmt.Errorf("JSON序列化失败: %v", err)
	}

	return string(jsonBytes), nil
}

// Model 测试分类ORM模型
func (s *sSysTestFinance) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.TestFinance.Ctx(ctx), option...)
}

// List 获取测试分类列表
func (s *sSysTestFinance) List(ctx context.Context, in *sysin.TestFinanceListInp) (list []*sysin.TestFinanceListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.TestFinanceListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.TestFinance.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.TestFinance.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.TestFinance.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.TestFinance.Columns().Sort).OrderDesc(dao.TestFinance.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取测试分类列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出测试分类
func (s *sSysTestFinance) Export(ctx context.Context, in *sysin.TestFinanceListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.TestFinanceExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出测试分类-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.TestFinanceExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增测试分类
func (s *sSysTestFinance) Edit(ctx context.Context, in *sysin.TestFinanceEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.TestFinanceUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改测试分类失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.TestFinanceInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增测试分类失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除测试分类
func (s *sSysTestFinance) Delete(ctx context.Context, in *sysin.TestFinanceDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除测试分类失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取测试分类最大排序
func (s *sSysTestFinance) MaxSort(ctx context.Context, in *sysin.TestFinanceMaxSortInp) (res *sysin.TestFinanceMaxSortModel, err error) {
	if err = dao.TestFinance.Ctx(ctx).Fields(dao.TestFinance.Columns().Sort).OrderDesc(dao.TestFinance.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(sysin.TestFinanceMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取测试分类指定信息
func (s *sSysTestFinance) View(ctx context.Context, in *sysin.TestFinanceViewInp) (res *sysin.TestFinanceViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取测试分类信息，请稍后重试！")
		return
	}
	return
}

// Status 更新测试分类状态
func (s *sSysTestFinance) Status(ctx context.Context, in *sysin.TestFinanceStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.TestFinance.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新测试分类状态失败，请稍后重试！")
		return
	}
	return
}

// Start 更新测试分类状态
func (s *sSysTestFinance) Start(ctx context.Context) (err error) {
	//"sh000001"
	codeList, err := service.SysFinanceCode().GetAllCode(ctx)

	for _, code := range codeList {
		if code.CompleteCode != "" {
			continue
		}
		_, _ = dao.FinanceCode.Ctx(ctx).Fields(dao.FinanceCode.Columns().CompleteCode).Where(dao.FinanceCode.Columns().Code, code.Code).Data(do.FinanceCode{CompleteCode: fmt.Sprintf("%s%s", gstr.ToLower(code.Exchange), code.Code)}).Update()

	}

	return
}

//
//func example (ctx context.Context) (err error) {
//	//AToken := "1b48ab3a3f318e1db193f5de915d4583-c-app"
//	/*
//		将如下JSON进行url的encode，复制到http的查询字符串的query字段里
//		{"trace" : "go_http_test1","data" : {"code" : "700.HK","kline_type" : 1,"kline_timestamp_end" : 0,"query_kline_num" : 2,"adjust_type": 0}}
//
//		特别注意：
//		github: https://github.com/alltick/realtime-forex-crypto-stock-tick-finance-websocket-api
//		token申请：https://alltick.co
//		把下面url中的testtoken替换为您自己的token
//		外汇，加密货币（数字币），贵金属的api址：
//		https://quote.alltick.io/quote-b-api
//		股票api地址:
//		https://quote.alltick.io/quote-stock-b-api
//	*/
//	url := "https://quote.alltick.io/quote-stock-b-api/kline"
//	log.Println("请求内容：", url)
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		fmt.Println("Error creating request:", err)
//		return nil
//	}
//
//	q := req.URL.Query()
//	q.Add("token", global.FinanceConfig.AlltickToken)
//
//	// 创建 FinanceAlltickRequest 实例
//	financeRequest := &entity.FinanceAlltickRequest{
//		Code:              "AAPL.US",
//		KlineType:         8,
//		KlineTimestampEnd: 0,
//		QueryKlineNum:     20,
//		AdjustType:        0,
//	}
//
//	// 使用转换方法生成查询字符串
//	queryStr, err := s.ConvertToQueryString(financeRequest)
//	if err != nil {
//		fmt.Println("Error converting request to query string:", err)
//		return nil
//	}
//	q.Add("query", queryStr)
//	req.URL.RawQuery = q.Encode()
//	// 发送请求
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		fmt.Println("Error sending request:", err)
//		return nil
//	}
//	defer resp.Body.Close()
//
//	body2, err := ioutil.ReadAll(resp.Body)
//
//	if err != nil {
//
//		log.Println("读取响应失败：", err)
//
//		return nil
//
//	}
//	var response entity.FinanceAlltickResponse
//	err = gconv.Scan(body2, &response)
//
//	if err != nil {
//		return err
//	}
//	return
//}
