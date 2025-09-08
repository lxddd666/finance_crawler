package sina

import (
	"context"
	"math/rand"
	"time"

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

func randomProxy() string {
	proxyList := []string{
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-zOoySxTi-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-XtiwhGyN-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-8BvtGGNn-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-n0N0oCLk-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-77Q9AYKj-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-hOIO4zC7-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-Swud9ZUo-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-DF5aHcp5-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-cbmKL9KW-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-mAUNmd0V-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-ZzHMxol5-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-JoXwb6Sf-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-rMeIbruo-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-iVUqdCKa-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-EwYRLqEg-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-ifUNP4jH-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-mlp5oKEd-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-c3BB78FN-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-x8O9SpCg-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-BB48oQ2M-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-VTQu1kAh-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-aSrPlV1f-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-4fNb76Oj-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-UHX7CwgK-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-GKVjoYiu-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-FpAq0GTW-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-N2MoFfeI-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-bcvKC0BH-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-7Dp3uLkm-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-Vyf57PmV-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-qp7tkHFd-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-adcfpPXf-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-EI0OJY1G-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-iSNODiul-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-cKI6ajSG-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-qibi0axb-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-0PTMC7fS-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-BFF4jKon-sessTime-15:TGMMlg2H",
		"na.proxys5.net:6200:81794984-zone-custom-region-US-state-alaska-sessid-zqg0eHap-sessTime-15:TGMMlg2H",
	}

	// 使用当前时间作为随机种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择一个代理
	randomIndex := rand.Intn(len(proxyList))
	return proxyList[randomIndex]
}
