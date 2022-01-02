// homesのスクレイピング

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func getPage(url string) *goquery.Document {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; MAFSJS; rv:11.0) like Gecko")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		panic("リクエスト作成中にエラー")
	}
	defer resp.Body.Close()
	if err != nil {
		panic("Getリクエスト中にエラー発生")
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic("Error occurs during getting html")
	}
	return doc
}

type Scraper struct {
	url string
	doc *goquery.Document
}

func NewScraper(url string) *Scraper {
	scraper := new(Scraper)
	scraper.url = url
	scraper.doc = getPage(url)
	return scraper
}

func (s *Scraper) getTitle() string {
	// selector := "title"
	title := s.doc.Find("title")
	return title.Text()
}

// 物件情報のテーブルからスクレイピング
func (s *Scraper) getDetailTable() map[string]string {
	selector := "div.mod-bukkenSpecDetail > table"
	table := s.doc.Find(selector)         // 物件詳細テーブルを取得
	table_data := make(map[string]string) // rubyでいうHash的なものを作る
	table.Find("th").Each(func(index int, s *goquery.Selection) {
		s.Next().Find("div").Remove() // tr内のdivを削除
		table_data[s.Text()] = s.Next().Text()
	})
	return table_data

}

// func (s *Scraper) getRentFee() string {
// 	rent_fee := s.doc.Find("span.price > span.num")
// 	return rent_fee.Text()
// }

func main() {
	url := "https://www.homes.co.jp/chintai/room/a3ee0fd00c2434192dea235ae73d489a2d568608/?bid=37031820006952"
	scraper := NewScraper(url)
	fmt.Println(scraper.getDetailTable())
}
