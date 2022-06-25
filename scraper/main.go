package main

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/google/brotli/go/cbrotli"
	"io/ioutil"
	"net/http"
	"time"
)

type ReaderType interface {
	Close() error
	Read(p []byte) (n int, err error)
}

//func GetDataBytes(r io.ReadCloser, encoding string) (data []byte, err error) {
func GetDataBytes(resp *http.Response, encoding string) (data []byte, err error) {
	var reader ReaderType = resp.Body
	if resp.Uncompressed {
		switch encoding {
		case "br":
			reader = cbrotli.NewReader(reader)
		case "gzip":
			reader, err = gzip.NewReader(reader)
		case "deflate":
			reader = flate.NewReader(reader)
		}
		if err != nil {
			return
		}
		defer reader.Close()
	}
	data, err = ioutil.ReadAll(reader)
	return
}

func main() {
	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	accept := "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"

	url := "https://www.hongkonghomes.com/en/hong-kong-property/for-sale/western-kennedy-town/tai-pak-terrace/185325?"
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", accept)
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")

	client := &http.Client{Timeout: time.Second * 10}

	response, _ := client.Do(request)
	defer response.Body.Close()

	fmt.Println(response.Header)

	var data []byte

	encoding := response.Header["Content-Encoding"][0]
	data, _ = GetDataBytes(response, encoding)

	finalString := string(data)
	//fmt.Println(finalString)

	doc, _ := htmlquery.Parse(strings.NewReader(finalString))
	list := htmlquery.Find(doc, "//script[@type=\"application/ld+json\"]/text()")

	for _, l := range list {
		res := (*l).Data

		var ares map[string]interface{}
		json.Unmarshal([]byte(res), &ares)
		fmt.Println(ares)

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
