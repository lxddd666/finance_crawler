package sina

import (
	"context"
	"golang.org/x/net/proxy"
	"io"
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

func GetKlineDataProxy(ctx context.Context, sinaReq *httpReq.SinaHttpReq) (response []*SinaResult, err error) {
	url := "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData"

	log.Println("请求内容 proxy：", url)
	queryStr, err := BuildSinaURL(sinaReq)
	url = fmt.Sprintf("%s?%s", url, queryStr)
	// 创建SOCKS5拨号器
	socks := RandomSocks5()
	dialer, err := proxy.SOCKS5("tcp", socks, nil, proxy.Direct)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			fmt.Println("获取失败:ip为" + socks)
		}
	}()
	// 创建HTTP传输层
	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	// 创建HTTP客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	// 发送请求
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = gconv.Scan(body, &response)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("获取股票[%s]结果数量[%d]", sinaReq.Symbol, len(response)))
	return
}

func RandomSocks5() string {
	proxyList := []string{
		"67.201.58.190:4145",
		"68.71.242.118:4145",
		"199.102.104.70:4145",
		"98.170.57.241:4145",
		"72.37.216.68:4145",
		"142.54.239.1:4145",
		"67.201.35.145:4145",
		"206.220.175.2:4145",
		"199.58.185.9:4145",
		"184.170.249.65:4145",
		"184.170.245.148:4145",
		"199.58.184.97:4145",
		"98.191.0.37:4145",
		"47.238.226.127:1024",
		"72.211.46.99:4145",
		"46.4.88.72:9050",
		"72.223.188.67:4145",
		"72.207.113.97:4145",
		"68.71.249.153:48606",
		"199.102.106.94:4145",
		"199.102.105.242:4145",
		"199.102.107.145:4145",
		"72.206.74.126:4145",
		"98.182.147.97:4145",
		"107.181.168.145:4145",
		"72.37.217.3:4145",
		"192.252.215.2:4145",
		"68.71.249.158:4145",
		"68.71.252.38:4145",
		"85.111.94.98:15833",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"127.0.0.1:7890",
		"192.252.209.158:4145",
		"98.190.239.3:4145",
		"199.116.112.6:4145",
		"184.170.251.30:11288",
		"98.182.171.161:4145",
		"192.111.129.150:4145",
		"192.252.214.17:4145",
		"174.75.211.193:4145",
		"198.177.252.24:4145",
		"142.54.237.38:4145",
		"68.71.245.206:4145",
		"98.175.31.222:4145",
		"68.71.243.14:4145",
		"68.71.240.210:4145",
		"68.71.254.6:4145",
		"72.207.109.5:4145",
		"74.119.144.60:4145",
		"98.181.137.80:4145",
		"68.1.210.163:4145",
		"98.181.137.83:4145",
		"68.71.247.130:4145",
		"74.119.147.209:4145",
		"192.252.216.81:4145",
		"184.178.172.17:4145",
		"72.205.0.67:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"198.177.254.157:4145",
		"192.252.220.92:17328",
		"184.181.217.201:4145",
		"184.181.217.206:4145",
		"184.181.217.210:4145",
		"184.181.217.213:4145",
		"184.181.217.220:4145",
		"184.178.172.11:4145",
		"184.178.172.14:4145",
		"184.178.172.18:15280",
		"184.178.172.23:4145",
		"184.178.172.25:15291",
		"184.178.172.26:4145",
		"184.178.172.28:15294",
		"184.181.217.194:4145",
		"184.178.172.3:4145",
		"184.178.172.5:15303",
		"174.64.199.79:4145",
		"174.77.111.196:4145",
		"198.177.254.131:4145",
		"68.1.210.189:4145",
		"72.205.0.93:4145",
		"192.252.216.86:4145",
		"192.252.211.193:4145",
		"68.71.241.33:4145",
		"67.201.39.14:4145",
		"174.75.211.222:4145",
		"192.252.220.89:4145",
		"107.152.98.5:4145",
		"198.8.84.3:4145",
		"142.54.232.6:4145",
		"142.54.235.9:4145",
		"142.54.229.249:4145",
		"142.54.237.34:4145",
		"142.54.231.38:4145",
		"68.71.251.134:4145",
		"198.177.253.13:4145",
		"199.187.210.54:4145",
		"192.111.135.17:18302",
		"72.214.108.67:4145",
		"103.82.25.14:10001",
		"198.8.94.174:39078",
		"192.252.215.5:16137",
	}
	// 使用当前时间作为随机种子
	rand.Seed(time.Now().UnixNano())

	// 随机选择一个代理
	randomIndex := rand.Intn(len(proxyList))
	return proxyList[randomIndex]
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
