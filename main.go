package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	url string
}

func getPage(url string) {
	doc, _ := goquery.NewDocument(url)
	fmt.Println(doc)
}

func main() {
	url := "https://www.homes.co.jp/chintai/room/a3ee0fd00c2434192dea235ae73d489a2d568608/?bid=37031820006952"
	getPage(url)
}


