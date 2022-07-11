package pipeline

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"

	"real_estate_scraping/infrastructure/redis"
)

var ErrEmptyCollection = errors.New("Empty collection")

func GetErrorString(err error) (errString *string) {
	if err != nil {
		s := err.Error()
		errString = &s
	}
	return
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

func GetUrlToScrapeWithTheLeastScore(collection string, ctx context.Context, c *gin.Context) {
	result, err := GetUrlWithTheLeastScore(collection, ctx)
	json := GetJsonForUrlToScrape(result, err)
	status := GetStatusForUrlToScrape(err)
	c.JSON(status, json)
}

func GetUrlWithTheLeastScore(collection string, ctx context.Context) (result *string, err error) {
	res, err := redis.GetStringValueWithTheLeastScore(collection, ctx)
	if len(res) == 0 {
		return nil, ErrEmptyCollection
	}
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}