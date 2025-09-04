package alltick

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"hotgo/internal/global"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"
	"io/ioutil"
	"log"
	"net/http"
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

// FinanceAlltickResponse is the golang structure for table finance_alltick_response.
type FinanceAlltickResponse struct {
	Id    int64         `json:"id"         orm:"id"          description:"分类ID"`
	Msg   string        `json:"msg"        orm:"msg"         description:""`
	Ret   int           `json:"ret"        `
	Trace string        `json:"trace"      orm:"trace"       description:""`
	Data  StockDataResp `json:"data"       orm:"data"        description:"data"`
}

type StockDataResp struct {
	Code      string       `json:"code"       orm:"code"        description:"code"`
	Sort      int          `json:"sort"       orm:"sort"        description:"排序"`
	KlineType int          `json:"kline_type"  orm:"kline_type"  description:"k线类型"`
	KlineList []StockKline `json:"kline_list"`
}

type StockKline struct {
	Timestamp  string `json:"timestamp"`
	OpenPrice  string `json:"open_price"  orm:"open_price"  description:"该K线开盘价"`
	ClosePrice string `json:"close_price" orm:"close_price" description:"该K线收盘价"`
	HighPrice  string `json:"high_price"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   string `json:"low_price"   orm:"low_price"   description:"该K线最低价"`
	Volume     string `json:"volume"     orm:"volume"      description:"该K线成交数量"`
	Turnover   string `json:"turnover"   orm:"turnover"    description:"该K线成交金额"`
}

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

func GetKlineData(ctx context.Context, financeRequest *entity.FinanceAlltickRequest) (response *FinanceAlltickResponse, err error) {
	url := "https://quote.alltick.io/quote-stock-b-api/kline"

	log.Println("请求内容：", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("token", global.FinanceConfig.AlltickToken)

	// 使用转换方法生成查询字符串
	queryStr, err := convertToQueryString(financeRequest)
	if err != nil {

		return
	}
	q.Add("query", queryStr)
	req.URL.RawQuery = q.Encode()
	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return

	}
	err = gconv.Scan(body2, &response)

	if err != nil {
		return
	}

	return
}

// convertToQueryString 将 FinanceAlltickRequest 结构体转换为 JSON 查询字符串
func convertToQueryString(req *entity.FinanceAlltickRequest) (string, error) {
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
