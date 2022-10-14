package main

import (
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/pipeline"
	"github.com/hu17889/go_spider/core/spider"
	"testing"
)

type TestSpider struct {
}

func (this *TestSpider) Process(p *page.Page) {
	query := p.GetHtmlParser()

	key := query.Find("div[id='tips'] h1").Text()
	value := query.Find("div[class='leftwrap article'] p").Text()

	p.AddField(key, value)
}

func (this *TestSpider) Finish() {

}

func Test(t *testing.T) {
	spider.NewSpider(&TestSpider{}, "TestSpider").
		AddUrl("http://xi-qu.com/pxx/zs/10878.html", "html").
		AddPipeline(pipeline.NewPipelineConsole()).
		Run()

}
