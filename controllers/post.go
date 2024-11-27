package controllers

import (
	"strconv"
	"web-app/logic"
	"web-app/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
{
    "community_id":1,
    "title":"c++",
    "content":"hello cpp"
}
*/
func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	author_id, err  := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNoAuth)
		return
	}
	p.AuthorID = author_id
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	page,size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		if err == logic.ErrorNoData {
			ResponseError(c, CodeNoData)
			return
		}
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID() failed", zap.Error(err))
		if err == logic.ErrorNoData {
			ResponseError(c, CodeNoData)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}