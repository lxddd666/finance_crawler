// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

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
