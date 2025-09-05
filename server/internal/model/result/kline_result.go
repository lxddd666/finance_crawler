package result

type Kline struct {
	Timestamp  string `json:"timestamp"`
	OpenPrice  string `json:"open_price"  orm:"open_price"  description:"该K线开盘价"`
	ClosePrice string `json:"close_price" orm:"close_price" description:"该K线收盘价"`
	HighPrice  string `json:"high_price"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   string `json:"low_price"   orm:"low_price"   description:"该K线最低价"`
	Volume     string `json:"volume"     orm:"volume"      description:"该K线成交数量"`
	Turnover   string `json:"turnover"   orm:"turnover"    description:"该K线成交金额"`
}
