package main

import (
	"github.com/hu17889/go_spider/core/common/com_interfaces"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/page_items"
	"github.com/hu17889/go_spider/core/spider"
	"regexp"
	"strings"
)

type ThirdSpider struct {
}

type ThirdPageSpider struct {
}

type ThirdPipeSpider struct {
}

func NewThirdSpider() *ThirdSpider {
	return &ThirdSpider{}
}

func (this *ThirdSpider) Run() {
	thirdSpider := spider.NewSpider(&ThirdPageSpider{}, "ThirdSpider")
	thirdSpider.AddPipeline(&ThirdPipeSpider{})
	for _, value := range UrlModel {
		for _, url := range value {
			thirdSpider.AddUrl(url, "html")
		}
	}
	thirdSpider.SetThreadnum(3)
	thirdSpider.Run()
}

func (this *ThirdPageSpider) Process(p *page.Page) {
	query := p.GetHtmlParser()

	url := p.GetRequest().Url
	title := query.Find("div[id='tips'] h1").Text()
	body, _ := query.Find("div[class='leftwrap article']").Html()

	body = strings.ReplaceAll(body, "src=\"/uploads", "src=\"http://xi-qu.com/uploads")
	body = strings.ReplaceAll(body, "href=\"/pxx", "href=\"http://xi-qu.com/pxx")
	footerRegexp, _ := regexp.Compile("<h3>[\\s\\S]*</h4>")
	body = footerRegexp.ReplaceAllString(body, "")

	ArticleModel[url] = Article{
		Title: title,
		Body:  body,
	}
}

func (this *ThirdPageSpider) Finish() {

}

func (this *ThirdPipeSpider) Process(items *page_items.PageItems, t com_interfaces.Task) {
}
