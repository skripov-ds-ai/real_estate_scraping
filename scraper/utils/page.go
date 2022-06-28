package utils

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var EXTRACTORS map[string]interface{} = map[string]interface{}{
	"general_info": "//table[contains(@class, \"general-info\")]",
	// div[@class='property-info']/h3
	// (//div[@class='property-info'])[1]/h3[3]/following-sibling::*[not(preceding-sibling::h3[4]) and not(self::h3)]
}

//func GenerateBetweenInfoXpath() {
//
//}

func ExtractUrls(doc *html.Node) (paginateUrls, itemUrls []string) {
	for _, itemNode := range htmlquery.Find(doc, "//div[@data-category='Search']//h3/a/@href") {
		itemUrls = append(itemUrls, itemNode.Data)
	}
	for _, paginateNode := range htmlquery.Find(doc, "//li[contains(@class, 'page-item') and not(contains(@class, 'disabled'))]//a/@href") {
		paginateUrls = append(paginateUrls, paginateNode.Data)
	}
	return
}
