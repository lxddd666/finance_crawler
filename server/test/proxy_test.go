package test

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

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

// parseProxy 解析动态代理格式
func parseProxy(proxyStr string) (*url.URL, error) {
	// 解析格式: host:port:username:password
	parts := strings.Split(proxyStr, ":")
	if len(parts) < 4 {
		return nil, fmt.Errorf("代理格式错误，应为 host:port:username:password")
	}

	// 提取各部分
	host := parts[0]
	port := parts[1]
	username := parts[2]
	password := parts[3]

	// 构建代理URL
	proxyURL := &url.URL{
		Scheme: "http", // 根据代理类型调整，可能是http、https或socks5
		Host:   fmt.Sprintf("%s:%s", host, port),
		User:   url.UserPassword(username, password),
	}

	return proxyURL, nil
}

// RequestWithRandomProxy 使用随机代理发送HTTP请求
func RequestWithRandomProxy(targetURL string) (string, error) {
	// 获取随机代理
	proxyStr := randomProxy()

	// 解析代理字符串
	proxyURL, err := parseProxy(proxyStr)
	if err != nil {
		return "", fmt.Errorf("解析代理失败: %v", err)
	}

	fmt.Printf("使用代理: %s\n", proxyURL.String())

	// 创建自定义Transport
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 仅用于测试，生产环境应设置为false
		},
		// 添加其他可能需要的配置
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	// 创建HTTP客户端
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
		// 设置不自动重定向，以便我们可以检查响应
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// 创建请求
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}

	// 设置请求头，模拟浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v (代理: %s)", err, proxyStr)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("请求失败，状态码: %d (代理: %s)", resp.StatusCode, proxyStr)
	}

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// RequestWithRetry 带重试机制的请求函数
func RequestWithRetry(targetURL string, maxRetries int) (string, error) {
	var lastError error

	for i := 0; i < maxRetries; i++ {
		fmt.Printf("尝试第 %d 次请求...\n", i+1)

		result, err := RequestWithRandomProxy(targetURL)
		if err == nil {
			return result, nil
		}

		lastError = err
		fmt.Printf("请求失败: %v\n", err)

		// 随机等待一段时间再重试
		waitTime := time.Duration(rand.Intn(5)+2) * time.Second
		fmt.Printf("等待 %.0f 秒后重试...\n", waitTime.Seconds())
		time.Sleep(waitTime)
	}

	return "", fmt.Errorf("经过 %d 次尝试后仍然失败: %v", maxRetries, lastError)
}

// TestProxyConnection 测试代理连接
func ProxyConnectionTest() error {
	proxyStr := randomProxy()
	proxyURL, err := parseProxy(proxyStr)
	if err != nil {
		return err
	}

	// 创建一个简单的请求测试代理
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// 使用一个简单的测试URL
	testURL := "http://httpbin.org/ip"
	req, err := http.NewRequest("GET", testURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("代理测试失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("代理测试返回非200状态码: %d", resp.StatusCode)
	}

	fmt.Println("代理连接测试成功!")
	return nil
}

func TestProxy(t *testing.T) {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 测试代理连接
	fmt.Println("测试代理连接...")
	err := ProxyConnectionTest()
	if err != nil {
		fmt.Printf("代理测试失败: %v\n", err)
	} else {
		fmt.Println("代理测试成功!")
	}

	// 要访问的URL
	targetURL := "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sz002095&scale=240&ma=no&datalen=1023" // 这个URL会返回当前使用的IP

	// 使用随机代理发送请求
	result, err := RequestWithRetry(targetURL, 3)
	if err != nil {
		fmt.Printf("最终请求失败: %v\n", err)
		return
	}

	fmt.Printf("请求成功! 响应内容:\n%s\n", result)

	// 示例：多次请求展示不同代理的效果
	fmt.Println("\n=== 展示不同代理的效果 ===")
	for i := 0; i < 3; i++ {
		result, err := RequestWithRandomProxy(targetURL)
		if err != nil {
			fmt.Printf("请求 %d 失败: %v\n", i+1, err)
		} else {
			fmt.Printf("请求 %d 成功: %s\n", i+1, result)
		}

		// 添加随机延迟，避免请求过于频繁
		delay := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(delay)
	}
}

func TestProxyV2(t *testing.T) {
	// 创建SOCKS5拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:7890", nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

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
	resp, err := client.Get("http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol=sz002095&scale=240&ma=no&datalen=1023")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("响应内容:\n%s\n", string(body))
}
