package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)


const (
	oneSeekSecond = 7 * 24 * 3600
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

// 记录新帖子
func InsertPost(postID int64) error {
	// 事务
	pipeline := rdb.TxPipeline()

	// 记录帖子时间
	pipeline.ZAdd(context.Background(), KeyPostTimeZSet, &redis.Z{
			Member: postID, 
			Score: float64(time.Now().Unix()),
		})
	// 记录帖子的分数
	pipeline.ZAdd(context.Background(), KeyPostScoreZSet, &redis.Z{
			Member: postID, 
			Score: 0,
	})

	_, err := pipeline.Exec(context.Background())

	return err
}

// 缺少并发控制，todo：加锁
func VoteForPostV0(user_id, post_id string, dir float64) (err error) {
	ctx := context.Background()

	// 获取帖子的发布时间,判断投票时间是否过期
	postTime := rdb.ZScore(ctx, KeyPostTimeZSet, post_id).Val()
	if float64(time.Now().Unix())-postTime > oneSeekSecond {
		return ErrVoteTimeExpire
	}


	// 获取之前该用户对该帖子的投票记录
	old_dir := rdb.ZScore(ctx, KeyPostVotedZSetPF+post_id, user_id).Val()
	// 如果之前投过票，且分数相同，则不做任何操作
	if old_dir == dir {
		return
	}

	// 如果用户之前投过票，且分数不同，则需要更新投票分数
	_, err = rdb.ZAdd(ctx, KeyPostVotedZSetPF+post_id, &redis.Z{
			Member: user_id, 
			Score: dir,
		}).Result()

	if err != nil {
		return err
	}

	// 更新帖子的总分数
	dif := dir - old_dir
	_, err = rdb.ZIncrBy(ctx, KeyPostScoreZSet, dif, post_id).Result()

	return
}


func VoteForPost(user_id, post_id string, dir float64) (err error) {
	ctx := context.Background()

	// 获取帖子的发布时间,判断投票时间是否过期
	postTime := rdb.ZScore(ctx, KeyPostTimeZSet, post_id).Val()
	if float64(time.Now().Unix())-postTime > oneSeekSecond {
		return ErrVoteTimeExpire
	}

	// 获取之前该用户对该帖子的投票记录
	old_dir := rdb.ZScore(ctx, KeyPostVotedZSetPF+post_id, user_id).Val()
	// 如果之前投过票，且分数相同，则不做任何操作
	if old_dir == dir {
		return
	}

	// 使用事务来确保操作的原子性
	_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// 如果用户之前投过票，且分数不同，则需要更新投票分数
		pipe.ZAdd(ctx, KeyPostVotedZSetPF+post_id, &redis.Z{
			Member: user_id,
			Score:  dir,
		})

		// 更新帖子的总分数
		dif := dir - old_dir
		pipe.ZIncrBy(ctx, KeyPostScoreZSet, dif, post_id)

		return nil
	})

	return err
}