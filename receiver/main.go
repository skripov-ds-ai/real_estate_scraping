package main

import (
	"context"
	"os"
	"real_estate_scraping/infrastructure/redis"
	"real_estate_scraping/pipeline"
	"real_estate_scraping/receiver/utils"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// 1. add url to scrape
// 2. add pagination url to scrape
func main() {
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDR")
	ginAddr := pipeline.MakeGinAddr(addr, port)

	redis.RedisConnect(ctx)

	r := gin.Default()
	r.POST("/add_item_url", func(c *gin.Context) {
		utils.PostItemUrl(c, ctx)
	})
	r.POST("/add_pagination_url", func(c *gin.Context) {
		utils.PostPaginationUrl(c, ctx)
	})
	r.Run(ginAddr)

	redis.RedisDisconnect()
}
