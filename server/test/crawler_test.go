package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGoldAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1) // 等待一个异步操作完成

	// 创建默认收集器
	c := colly.NewCollector()

	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})

	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		defer wg.Done() // 请求完成后通知WaitGroup
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			t.Errorf("解析文档失败: %v", err)
			return
		}

		firstHead := doc.Find("div.header").First()
		nowPrice := firstHead.Find("strong.text-xlg").Text()
		firstHead.Find("div.mo").Each(func(i int, s *goquery.Selection) {
			hightestPrice := s.Text()
			fmt.Println(hightestPrice)
		})
		fmt.Println(nowPrice)
	})

	c.OnError(func(r *colly.Response, err error) {
		defer wg.Done() // 错误情况下也要结束等待
		t.Errorf("请求失败: %v", err)
	})

	c.SetRequestTimeout(20 * time.Second)

	// 访问起始页面
	err := c.Visit("https://xau.today/zh/gold-price-ounce/#google_vignette")
	if err != nil {
		t.Errorf("访问URL失败: %v", err)
		return
	}

	wg.Wait() // 等待异步操作完成
}

func Test10BondAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1) // 等待一个异步操作完成

	// 创建默认收集器
	c := colly.NewCollector()

	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})

	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		defer wg.Done() // 请求完成后通知WaitGroup
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			fmt.Printf("解析文档失败: %v", err)
			return
		}

		// 查找目标 <div> 元素
		doc.Find("div.mb-3.flex.flex-wrap.items-center.gap-x-4.gap-y-2.md\\:mb-0\\.5.md\\:gap-6").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Html()
			fmt.Println(a)
			b := s.Text()
			fmt.Println(b)
			price := s.Find("div[data-test='instrument-price-last']").Text()
			priceChange := s.Find("div[data-test='instrument-price-change']").Text()
			changePercent := s.Find("div[data-test='instrument-price-change-percent']").Text()
			fmt.Println(price, priceChange, changePercent)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		defer wg.Done() // 错误情况下也要结束等待
		t.Errorf("请求失败: %v", err)
	})

	c.SetRequestTimeout(20 * time.Second)

	// 访问起始页面
	err := c.Visit("https://cn.investing.com/rates-bonds/u.s.-10-year-bond-yield")
	if err != nil {
		t.Errorf("访问URL失败: %v", err)
		return
	}

	wg.Wait() // 等待异步操作完成
}

func TestUsFedFundInterestRate(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1) // 等待一个异步操作完成

	// 创建默认收集器
	c := colly.NewCollector()

	// 向 API 发送请求
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("发送请求到：", r.URL)
	})

	// 处理 API 返回的数据
	c.OnResponse(func(r *colly.Response) {
		defer wg.Done() // 请求完成后通知WaitGroup
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			fmt.Printf("解析文档失败: %v", err)
			return
		}
		fmt.Println(doc)
	})

	c.OnError(func(r *colly.Response, err error) {
		defer wg.Done() // 错误情况下也要结束等待
		t.Errorf("请求失败: %v", err)
	})

	c.SetRequestTimeout(20 * time.Second)

	// 访问起始页面
	err := c.Visit("https://fred.stlouisfed.org/graph/api/series/?id=IORB%2CSOFR%2CEFFR&width=1030")
	if err != nil {
		t.Errorf("访问URL失败: %v", err)
		return
	}

	wg.Wait() // 等待异步操作完成
}
