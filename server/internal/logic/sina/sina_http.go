package sina

import (
	"context"

	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"

	"hotgo/internal/model/httpReq"
	"io/ioutil"
	"log"
	"net/http"
)

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

type SinaResult struct {
	Day    string `json:"day"`
	Open   string `json:"open"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Close  string `json:"close"`
	Volume string `json:"volume"`
}

func GetKlineData(ctx context.Context, sinaReq *httpReq.SinaHttpReq) (response []*SinaResult, err error) {
	//url := "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sz002095&scale=240&ma=no&datalen=1023"
	url := "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData"

	log.Println("请求内容：", url)
	queryStr, err := BuildSinaURL(sinaReq)
	url = fmt.Sprintf("%s?%s", url, queryStr)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	q := req.URL.Query()
	q.Add("", queryStr)

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

func BuildSinaURL(req *httpReq.SinaHttpReq) (string, error) {
	// 如果 trace 为空，使用默认值
	if req.Symbol == "" {
		return "", gerror.New("股票代码为空")

	}
	if req.Ma == "" {
		req.Ma = "no"
	}
	if req.Datalen == 0 {
		req.Datalen = 20
	}
	if req.Scale == 0 {
		// 一天
		req.Scale = 240
	}

	// 创建URL参数
	params := url.Values{}

	if req.Symbol != "" {
		params.Add("symbol", req.Symbol)
	}
	if req.Scale != 0 {
		params.Add("scale", gconv.String(req.Scale))
	}
	if req.Ma != "" {
		params.Add("ma", req.Ma)
	}
	if req.Datalen != 0 {
		params.Add("datalen", gconv.String(req.Datalen))
	}

	// 拼接URL
	if len(params) > 0 {
		return params.Encode(), nil
	}

	return "", gerror.New("拼接url失败")
}
