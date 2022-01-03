package scraper

import (
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

func (s *Scraper) GetTitle() string {
	// selector := "title"
	title := s.doc.Find("title")
	return title.Text()
}

// 物件情報のテーブルからスクレイピング
func (s *Scraper) GetDetailTable() map[string]string {
	selector := "div.mod-bukkenSpecDetail > table"
	table := s.doc.Find(selector)         // 物件詳細テーブルを取得
	table_data := make(map[string]string) // rubyでいうHash的なものを作る
	table.Find("th").Each(func(index int, s *goquery.Selection) {
		s.Next().Find("div.text").Remove() // tr内のdivを削除
		s.Next().Find("div.inquire").Remove() // tr内のdivを削除
		table_data[s.Text()] = s.Next().Text()
	})
	return table_data

}

// func (s *Scraper) getRentFee() string {
// 	rent_fee := s.doc.Find("span.price > span.num")
// 	return rent_fee.Text()
// }
