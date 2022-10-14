package main

import (
	"fmt"
	"os"
)

type Article struct {
	Title string
	Body  string
}

var (
	UrlModel     = make(map[string][]string)
	PageModel    = make(map[string]string)
	ArticleModel = make(map[string]Article)
)

var FileHeaderTemplate = `---
title: %s
description:
published: true
date: 2022-10-14T00:18:39.501Z
tags:
editor: markdown
dateCreated: 2022-10-14T00:14:14.108Z
---`

var FileBodyTemplate = `
%s


`

func main() {
	NewFirstSpider().Run()
	NewSecondSpider().Run()
	NewThirdSpider().Run()
	for key, _ := range PageModel {
		saveToFile(key)
	}
}

func saveToFile(rootUrl string) {
	header := PageModel[rootUrl]

	saveFile, _ := os.OpenFile(header+".md", os.O_RDWR|os.O_CREATE, 0666)

	defer saveFile.Close()

	var articles []Article
	for _, url := range UrlModel[rootUrl] {
		articles = append(articles, ArticleModel[url])
	}

	saveFile.WriteString(fmt.Sprintf(FileHeaderTemplate, header))
	for _, article := range articles {
		saveFile.WriteString(fmt.Sprintf(FileBodyTemplate, article.Body))
	}
}
