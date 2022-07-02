package utils

import (
	"math/rand"
	"time"
)

type UrlSetter struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

func generateScore() float64 {
	return float64(time.Now().UnixNano()) + rand.Float64()
}
