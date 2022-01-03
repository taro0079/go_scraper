package scraper

import (
	"regexp"
	"strconv"
	"strings"
)

type TextFormatting struct {
	RawData map[string]string
}

func NewTextFormatting(raw_data map[string]string) *TextFormatting {
	text_formatting := new(TextFormatting)
	text_formatting.RawData = raw_data
	return text_formatting
}

func (tf TextFormatting) RentFee() int {
	text_data := tf.RawData["賃料（管理費等）"]
	formatted_text := tf.DeleteSpace(text_data)
	reg := regexp.MustCompile(`\d.*万円`)
	rent_fee := reg.FindAllStringSubmatch(formatted_text, -1)[0][0]
	deleted_rent_fee := strings.Replace(rent_fee, "万円", "", -1)

	rent_fee_data, err := strconv.ParseFloat(deleted_rent_fee, 64)
	ten_thousand_rent_fee := rent_fee_data * 10000

	if err != nil {
		panic("error while converting string to float !")
	}
	return int(ten_thousand_rent_fee)
}

func (tf TextFormatting) DeleteSpace(text string) string {
	reg := regexp.MustCompile(`\s`)
	return reg.ReplaceAllString(text, "")
}
