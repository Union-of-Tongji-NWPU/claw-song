package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
	"os"
	"time"
)

type Movie struct {
	idx    string
	title  string
	year   string
	info   string
	rating string
	url    string
}

func main() {
	// 存储文件名
	fName := "test.txt"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("创建文件失败 %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// 起始Url
	//startUrl := "https://movie.douban.com/top250"
	startUrl := "https://noobnotes.net/zombie-the-cranberries?solfege=false"

	// 创建Collector
	collector := colly.NewCollector()
	extensions.RandomUserAgent(collector)

	// 设置抓取频率限制
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 5 * time.Second, // 随机延迟
	})

	// 异常处理
	collector.OnError(func(response *colly.Response, err error) {
		log.Println(err.Error())
	})

	collector.OnRequest(func(request *colly.Request) {
		log.Println("start visit: ", request.URL.String())
	})

	// 解析列表

	collector.OnHTML("body", func(element *colly.HTMLElement) {

		selection := element.DOM.Find("div.post-content")
		// 依次遍历所有的p节点
		selection.Find("p").Each(func(i int, selection *goquery.Selection) {
			text := selection.Text()
			// 进一步处理
			fmt.Println(text)
		})
	})
	// 起始入口
	collector.Visit(startUrl)
}
