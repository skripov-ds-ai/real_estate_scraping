package redis

import (
	"context"
	"github.com/go-redis/redis/v9"
)

func AddValueToSortedSet(value, collection string, score float64, ctx context.Context) (err error) {
	err = Client.ZAdd(ctx, collection, redis.Z{Score: score, Member: value}).Err()
	return err
}

func GetValueWithTheLeastScore(collection string, ctx context.Context) (result []redis.Z, err error) {
	result, err = Client.ZPopMin(ctx, collection).Result()
	return
}

func GetStringValueWithTheLeastScore(collection string, ctx context.Context) (result []string, err error) {
	res, err := GetValueWithTheLeastScore(collection, ctx)
	for _, r := range res {
		result = append(result, r.Member.(string))
	}
	return
}
