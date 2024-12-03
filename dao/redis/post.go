package redis

import (
	"context"
	"fmt"
	"web-app/models"
)


func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := KeyPostScoreZSet
	if p.Order == models.OrderTime {
		key = KeyPostTimeZSet
	}

	fmt.Println("redis p: ", p.Order)

	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}