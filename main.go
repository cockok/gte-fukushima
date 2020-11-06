package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/go-querystring/query"
)

func main() {

	const baseUrl string = "https://gotoeat-fukushima.jp/shop/?"

	type Options struct {
		Area     string `url:"area[]"`
		Category string `url:"category[]"`
		Search   string `url:"s"`
	}

	area := "会津若松市,喜多方市,取麻郡,河沼郡,大沼郡,西会津町,磐梯町,猪苗代町,会津坂下町,柳津町,三島町,金山町,北塩原村,湯川村,会津美里町,昭和村,耶麻郡北塩原村,耶麻郡西会津町,耶麻郡猪苗代町,耶麻郡磐梯町,河沼郡湯川村,耶麻郡北塩原村大字,河沼郡会津坂下町,河沼郡柳津町,河沼郡柳津町大字,大沼郡昭和村,大沼郡会津美里町"
	// areas := []string{
	// 	"会津若松市",
	// 	"喜多方市",
	// 	"取麻郡",
	// 	"河沼郡",
	// 	"大沼郡",
	// 	"西会津町",
	// 	"磐梯町",
	// 	"猪苗代町",
	// 	"会津坂下町",
	// 	"柳津町",
	// 	"三島町",
	// 	"金山町",
	// 	"北塩原村",
	// 	"湯川村",
	// 	"会津美里町",
	// 	"昭和村",
	// 	"耶麻郡北塩原村",
	// 	"耶麻郡西会津町",
	// 	"耶麻郡猪苗代町",
	// 	"耶麻郡磐梯町",
	// 	"河沼郡湯川村",
	// 	"耶麻郡北塩原村大字",
	// 	"河沼郡会津坂下町",
	// 	"河沼郡柳津町",
	// 	"河沼郡柳津町大字",
	// 	"大沼郡昭和村",
	// }
	categories := map[string]string{
		"cat01": "和食・寿司",
		"cat02": "洋食",
		"cat03": "中華料理",
		"cat04": "ラーメン・餃子",
		"cat05": "うどん・そば・丼",
		"cat06": "フレンチ・イタリアン",
		"cat07": "焼肉・ホルモン・韓国料理",
		"cat08": "すき焼き・しゃぶしゃぶ",
		"cat09": "アジア・エスニック・各国料理",
		"cat10": "ファミリーレストラン・食堂",
		"cat11": "カフェ・スイーツ",
		"cat12": "居酒屋",
		"cat13": "バー・ダイニングバー",
		"cat14": "ファーストフード",
	}

	for slug := range categories {
		q, _ := query.Values(Options{area, slug, ""})
		parse(baseUrl + q.Encode())
	}
}

func parse(url string) {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return
	}
	li := doc.Find("ul.list_search-result > li")
	li.Each(func(index int, selection *goquery.Selection) {
		fmt.Println(selection.Find(".result-cat").Text() + "," + selection.Find(".result-name").Text() + "," + selection.Find(".result-address").Text())
	})
	nextUrl, exists := doc.Find("a.nextpostslink").First().Attr("href")
	if exists {
		parse(nextUrl)
	}
}
