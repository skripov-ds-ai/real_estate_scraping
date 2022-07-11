package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"real_estate_scraping/infrastructure/redis"
	"real_estate_scraping/pipeline"
)

func getJsonForPostUrlToScrape(err error) gin.H {
	return gin.H{
		"error": pipeline.GetErrorString(err),
	}
}

func postUrl(c *gin.Context, collection string, ctx context.Context) {
	var urlSetter UrlSetter
	status := http.StatusInternalServerError
	if err := c.BindJSON(&urlSetter); err != nil {
		c.JSON(status, getJsonForPostUrlToScrape(err))
		return
	}
	err := redis.AddValueToSortedSet(urlSetter.URL, collection, generateScore(), ctx)
	if err == nil {
		status = http.StatusCreated
	}
	c.JSON(status, getJsonForPostUrlToScrape(err))
}

func PostItemUrl(c *gin.Context, ctx context.Context) {
	postUrl(c, "item", ctx)
}

func PostPaginationUrl(c *gin.Context, ctx context.Context) {
	postUrl(c, "pagination", ctx)
}
