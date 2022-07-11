package main

import (
	"encoding/json"
	"fmt"
	"real_estate_scraping/scraper/utils"
	"strings"

	"github.com/antchfx/htmlquery"
	"net/http"
	"time"
)

func main() {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"

	//url := "https://www.hongkonghomes.com/en/hong-kong-property/for-sale/western-kennedy-town/tai-pak-terrace/185325?"
	url := "https://www.hongkonghomes.com/en/hong-kong-property/for-rent/tsim-sha-tsui/the-arch-block-2a-moon-tower/64427?"
	request, _ := http.NewRequest("GET", url, nil)
	headers := utils.GetOrdinaryPageHeaders(userAgent)
	utils.SetHeaders(request, &headers)

	client := &http.Client{Timeout: time.Second * 10}

	response, _ := client.Do(request)
	defer response.Body.Close()

	fmt.Println(response.Header)

	encoding := utils.GetEncoding(response)
	data, _ := utils.GetDataBytes(response, encoding)

	finalString := string(data)
	//fmt.Println(finalString)

	doc, _ := htmlquery.Parse(strings.NewReader(finalString))
	list := htmlquery.Find(doc, "//script[@type=\"application/ld+json\"]/text()")

	fmt.Println()

	for _, l := range list {
		res := (*l).Data

		var ares map[string]interface{}
		json.Unmarshal([]byte(res), &ares)
		fmt.Println(ares)
		fmt.Println()

		//fmt.Println(res)
	}
	//fmt.Println(list)

	//if err != nil {
	//	fmt.Printf("The HTTP request failed with error %s\n", err)
	//}

	// it is not for all items; only for vr
	// https://www.hongkonghomes.com/en/hong-kong-property/for-sale/happy-valley/fung-fai-terrace-18-19a/83889?
	//graphqlUrl := "https://my.matterport.com/api/mp/models/graph"
	//
	//userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	//referer := "https://my.matterport.com/show/?play=1&m=yFHoSPfUWZF"
	//modelId, _ := utils.GetModelId(referer)
	//
	//headers := utils.CreateHeaders(userAgent, referer)
	//payload := utils.CreatePayload(*modelId)
	//
	//payloadValue, _ := json.Marshal(payload)
	//
	//request, err := http.NewRequest("POST", graphqlUrl, bytes.NewBuffer(payloadValue))
	//for k, v := range headers {
	//	request.Header.Set(k, v)
	//}
	//
	//client := &http.Client{Timeout: time.Second * 10}
	//
	//response, err := client.Do(request)
	//defer response.Body.Close()
	//if err != nil {
	//	fmt.Printf("The HTTP request failed with error %s\n", err)
	//}
	//
	//reader, err := gzip.NewReader(response.Body)
	//defer reader.Close()
	//
	//data, _ := ioutil.ReadAll(reader)
	//mp := make(map[string]interface{})
	//json.Unmarshal([]byte(data), &mp)
	//fmt.Println(mp)
	////fmt.Println(data)

}
