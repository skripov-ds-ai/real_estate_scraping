package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"real_estate_scraping/scraper/utils"
	"time"
)

func main() {
	graphqlUrl := "https://my.matterport.com/api/mp/models/graph"

	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36"
	referer := "https://my.matterport.com/show/?play=1&m=yFHoSPfUWZF"
	modelId := utils.GetModelId(referer)

	headers := utils.CreateHeaders(userAgent, referer)
	payload := utils.CreatePayload(modelId)

	payloadValue, _ := json.Marshal(payload)

	request, err := http.NewRequest("POST", graphqlUrl, bytes.NewBuffer(payloadValue))
	for k, v := range headers {
		request.Header.Set(k, v)
	}

	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	reader, err := gzip.NewReader(response.Body)
	defer reader.Close()

	data, _ := ioutil.ReadAll(reader)
	mp := make(map[string]interface{})
	json.Unmarshal([]byte(data), &mp)
	fmt.Println(mp)
	fmt.Println(data)

}
