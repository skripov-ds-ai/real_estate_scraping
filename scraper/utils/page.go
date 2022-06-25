package utils

var EXTRACTORS map[string]interface{} = map[string]interface{}{
	"general_info": "//table[contains(@class, \"general-info\")]",
	// div[@class='property-info']/h3
	// (//div[@class='property-info'])[1]/h3[3]/following-sibling::*[not(preceding-sibling::h3[4]) and not(self::h3)]
}

func GenerateBetweenInfoXpath() {

}
