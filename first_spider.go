package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/spider"
)

type FirstSpider struct {
}

type FirstPageSpider struct {
}

type FirstPipeSpider struct {
}

func NewFirstSpider() *FirstSpider {
	return &FirstSpider{}
}

func (this *FirstSpider) Run() {
	firstSpider := spider.NewSpider(&FirstPageSpider{}, "FirstSpider")
	firstSpider.AddPipeline(&FirstPipeSpider{})
	firstSpider.AddUrl("http://xi-qu.com/pxx", "html")
	firstSpider.SetThreadnum(3)
	firstSpider.Run()
}

func (this *FirstPageSpider) Process(p *page.Page) {
	query := p.GetHtmlParser()
	query.Find("div[class='pdcolumn'] h2 a").Each(func(i int, s *goquery.Selection) {
		url := "http://xi-qu.com" + s.AttrOr("href", "")
		if s.Text() == "MORE" {
			return
		}
		PageModel[url] = s.Text()
	})
}

func (this *FirstPageSpider) Finish() {

}

func (this *FirstPipeSpider) Process(items *page_items.PageItems, t com_interfaces.Task) {
}
