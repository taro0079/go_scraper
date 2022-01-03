// homesのスクレイピング

package main

import (
	"fmt"
	"github.com/taro0079/scraping_homes/scraper"
)



func main() {
	url := "https://www.homes.co.jp/chintai/room/a3ee0fd00c2434192dea235ae73d489a2d568608/?bid=37031820006952"
	s := scraper.NewScraper(url)
	mapdata := s.GetDetailTable()
	formatted := scraper.NewTextFormatting(mapdata)
	fmt.Println(formatted.RentFee())
	

}
