package logic

import (
	"strconv"
	"web-app/dao/redis"
	"web-app/models"

	"go.uber.org/zap"
)

/*
投票算法：
	1. 牛顿冷却定律：热度=赞成票-反对票
	2.
投票的几种情况：
direction=1时，有两种情况：
	1. 之前没有投过票，现在投赞成票--》更新帖子的分数+1
	2. 之前投反对票，现在改投赞成票--》更新帖子的分数+2

direction=0时，有两种情况：
	1. 之前投过赞成票，现在要取消投票--》更新帖子的分数-1
	2. 之前投过反对票，现在要取消投票--》更新帖子的分数+1

direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票--》更新帖子的分数-1
	2. 之前投赞成票，现在改投反对票--》更新帖子的分数-2

投票的限制：
	1. 同一个用户对同一个帖子只能投一次票
	2. 投票的数据有效期为7天，7天后redis自动删除，票数保存到mysql中
*/


func VoteForPost(user_id int64, p *models.ParamVoteData) (err error) {
	zap.L().Debug("VoteForPost", zap.Int64("user_id", user_id),
				 zap.String("post_id", p.PostID), zap.Int8("direction", p.Direction))
	err = redis.VoteForPost(strconv.Itoa(int(user_id)), p.PostID, float64(p.Direction))
	if err != nil {
		if err == redis.ErrVoteTimeExpire {
			// 转为logic层的错误类型
			err = ErrorVoteTimeExpire
		}
		zap.L().Error("VoteForPost failed", zap.Error(err))
		return
	}
	return
}
