package main

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"os"
	"time"

	"real_estate_scraping/infrastructure/redis"

	"github.com/gin-gonic/gin"
)

var SCRAPER_URL string = os.Getenv("SCRAPER_URL")
var ctx = context.Background()
var ErrEmptyCollection = errors.New("Empty collection")

type UrlSetter struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func generateScore() float64 {
	return float64(time.Now().UnixNano()) + rand.Float64()
}

func postUrlToScrape(c *gin.Context) {
	var urlSetter UrlSetter
	status := http.StatusInternalServerError
	if err := c.BindJSON(&urlSetter); err != nil {
		c.JSON(status, GetJsonForPostUrlToScrape(err))
		return
	}
	collection := "to_scrape"
	err := redis.AddValueToSortedSet(urlSetter.URL, collection, generateScore(), ctx)
	if err == nil {
		// TODO: add handling duplicates? sortedSet store only the first value
		status = http.StatusCreated
	}
	c.JSON(status, GetJsonForPostUrlToScrape(err))
}

func GetErrorString(err error) (errString *string) {
	if err != nil {
		s := err.Error()
		errString = &s
	}
	return
}

func GetJsonForPostUrlToScrape(err error) gin.H {
	return gin.H{
		"error": GetErrorString(err),
	}
}

func GetJsonForUrlToScrape(result *string, err error) gin.H {
	errString := GetErrorString(err)
	return gin.H{
		"error": errString,
		"data":  result,
	}
}

func GetStatusForUrlToScrape(err error) (status int) {
	status = http.StatusOK
	if err != nil && err != ErrEmptyCollection {
		status = http.StatusInternalServerError
	}
	return
}

func getUrlToScrapeWithTheLeastScore(c *gin.Context) {
	result, err := GetUrlWithTheLeastScore("to_scrape")
	json := GetJsonForUrlToScrape(result, err)
	status := GetStatusForUrlToScrape(err)
	c.JSON(status, json)
}

func GetUrlWithTheLeastScore(collection string) (result *string, err error) {
	res, err := redis.GetStringValueWithTheLeastScore(collection, ctx)
	if len(res) == 0 {
		return nil, ErrEmptyCollection
	}
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

// 1. add url to scrape
// 2. add pagination url to scrape
func main() {
	redis.RedisConnect(ctx)

	r := gin.Default()
	r.POST("/add_url_to_scrape", postUrlToScrape)
	r.GET("/get_url_to_scrape", getUrlToScrapeWithTheLeastScore)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	redis.RedisDisconnect()
}
