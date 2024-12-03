package controllers

import (
	"web-app/logic"
	"web-app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


type VoteData struct {
	PostID int64 `json:"post_id,string"`		// 帖子id
	Direction int8 `json:"direction,string"` 	// 1赞成 0反对
}

/*
{
	"post_id":"",
	"direction":
}
*/
func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	user_id, err  := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNoAuth)
		return
	}
	if err := logic.VoteForPost(user_id, p); err != nil {
		if err == logic.ErrorVoteTimeExpire {
			ResponseError(c, CodeInvalidParam)
			return
		}
		zap.L().Error("logic.VoteForPost(user_id, p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, "投票成功！")
}
