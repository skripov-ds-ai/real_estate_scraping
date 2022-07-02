package main

import (
	"context"
	"math/rand"
	"net/http"
	"real_estate_scraping/pipeline"
	"time"

	"real_estate_scraping/infrastructure/redis"

	"github.com/gin-gonic/gin"
)

//var SCRAPER_URL string = os.Getenv("SCRAPER_URL")
var ctx = context.Background()

type UrlSetter struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func generateScore() float64 {
	return float64(time.Now().UnixNano()) + rand.Float64()
}

func postUrl(c *gin.Context, collection string) {
	var urlSetter UrlSetter
	status := http.StatusInternalServerError
	if err := c.BindJSON(&urlSetter); err != nil {
		c.JSON(status, GetJsonForPostUrlToScrape(err))
		return
	}
	err := redis.AddValueToSortedSet(urlSetter.URL, collection, generateScore(), ctx)
	if err == nil {
		// TODO: add handling duplicates? sortedSet store only the first value
		status = http.StatusCreated
	}
	c.JSON(status, GetJsonForPostUrlToScrape(err))
}

func postItemUrl(c *gin.Context) {
	postUrl(c, "item")
}

func postPaginationUrl(c *gin.Context) {
	postUrl(c, "pagination")
}

func GetJsonForPostUrlToScrape(err error) gin.H {
	return gin.H{
		"error": pipeline.GetErrorString(err),
	}
}

// 1. add url to scrape
// 2. add pagination url to scrape
func main() {
	redis.RedisConnect(ctx)

	r := gin.Default()
	r.POST("/add_item_url", postItemUrl)
	r.POST("/add_pagination_url", postPaginationUrl)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	redis.RedisDisconnect()
}
