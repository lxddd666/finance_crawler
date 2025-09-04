package entity

type StockKlineData struct {
	Code       string  `json:"code"       orm:"code"        description:"code"`
	Timestamp  int64   `json:"timestamp"`
	OpenPrice  float64 `json:"open_price"  orm:"open_price"  description:"该K线开盘价"`
	ClosePrice float64 `json:"close_price" orm:"close_price" description:"该K线收盘价"`
	HighPrice  float64 `json:"high_price"  orm:"high_price"  description:"该K线最高价"`
	LowPrice   float64 `json:"low_price"   orm:"low_price"   description:"该K线最低价"`
	Volume     int64   `json:"volume"     orm:"volume"      description:"该K线成交数量"`
	Turnover   float64 `json:"turnover"   orm:"turnover"    description:"该K线成交金额"`
}
