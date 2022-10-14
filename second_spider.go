package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/spider"
)

type SecondSpider struct {
}

type SecondPageSpider struct {
}

type SecondPipeSpider struct {
}

func NewSecondSpider() *SecondSpider {
	return &SecondSpider{}
}

func (this *SecondSpider) Run() {
	secondSpider := spider.NewSpider(&SecondPageSpider{}, "SecondSpider")
	secondSpider.AddPipeline(&SecondPipeSpider{})
	for key, _ := range PageModel {
		//if strings.Contains(key, "ds") {
		secondSpider.AddUrl(key, "html")
		//}
	}
	secondSpider.SetThreadnum(3)
	secondSpider.Run()
}

func (this *SecondPageSpider) Process(p *page.Page) {
	query := p.GetHtmlParser()
	url := p.GetRequest().GetUrl()
	query.Find("div[class='leftwrap list'] h2 a").Each(func(i int, selection *goquery.Selection) {
		val, _ := selection.Attr("href")
		val = "http://xi-qu.com" + val
		UrlModel[url] = append(UrlModel[url], val)
	})
}

func (this *SecondPageSpider) Finish() {

}

func (this *SecondPipeSpider) Process(items *page_items.PageItems, t com_interfaces.Task) {

}
