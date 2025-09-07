package httpReq

type SinaHttpReq struct {
	Symbol  string `json:"symbol"  description:"股票代码"`
	Scale   int    `json:"scale"   description:"分钟，240为一天"`
	Ma      string `json:"ma"      description:"平均线"`
	Datalen int    `json:"datalen" description:"k线数"`
}
