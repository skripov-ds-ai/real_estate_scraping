package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var SCRAPER_URL string = os.Getenv("SCRAPER_URL")
var ctx = context.Background()
var Client *redis.Client
var Config RedisConfig

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func (r *RedisConfig) Load() {
	r.Addr = os.Getenv("Addr")
	r.Password = os.Getenv("Password")
	DB, err := strconv.ParseInt(os.Getenv("DB"), 10, 32)
	if err != nil {
		panic(err)
	}
	r.DB = int(DB)
}

func RedisConnect() {
	Config.Load()
	options := redis.Options{
		Addr:     Config.Addr,
		Password: Config.Password,
		DB:       Config.DB,
	}
	Client = redis.NewClient(&options)

	if _, err := Client.Ping(ctx).Result(); err != nil {
		panic(err)
	}
}

func RedisDisconnect() {
	_ = Client.Close()
}

func SetUrl(url, collection string, score float64) (ok bool, err error) {
	if collection != "to_scrape" && collection != "to_paginate" {
		//panic("Wrong collection")
		return
	}
	err = Client.ZAdd(ctx, collection, redis.Z{Score: score, Member: url}).Err()
	return err == nil, err
}

type UrlSetter struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func generateScore() float64 {
	return float64(time.Now().UnixNano()) + rand.Float64()
}

func postUrlToScrape(c *gin.Context) {
	//c.Params.Get("url")

	var urlSetter UrlSetter
	if err := c.BindJSON(&urlSetter); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": urlSetter})
		return
	}
	ok, _ := SetUrl(urlSetter.URL, "to_scrape", generateScore())
	if ok {
		c.IndentedJSON(http.StatusCreated, gin.H{"data": urlSetter})
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": urlSetter})
}

func getUrlToScrapeWithTheLeastScore(c *gin.Context) {
	result, err := GetUrlWithTheLeastScore("to_scrape")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"data": "",
		})
		return
	}
	if len(result) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": "",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": result[0],
	})
}

func GetUrlWithTheLeastScore(collection string) (result []string, err error) {
	// https://stackoverflow.com/questions/20017255/how-to-get-a-member-with-maximum-or-minimum-score-from-redis-sorted-set-given
	//result, err = Client.ZRangeByScore(
	//	ctx,
	//	collection,
	//	&redis.ZRangeBy{
	//		Min:    "-inf",
	//		Max:    "inf",
	//		Offset: 0,
	//		Count:  1,
	//	},
	//).Result()
	res, err := Client.ZPopMin(ctx, collection).Result()
	for _, r := range res {
		result = append(result, r.Member.(string))
	}
	return
}

// 1. add url to scrape
// 2. add pagination url to scrape
func main() {
	RedisConnect()

	r := gin.Default()
	r.POST("/add_url_to_scrape", postUrlToScrape)
	r.GET("/get_url_to_scrape", getUrlToScrapeWithTheLeastScore)
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	RedisDisconnect()
}
