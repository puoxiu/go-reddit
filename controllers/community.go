package controllers

import (
	"strconv"
	"web-app/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//
func CommunityHandler(c *gin.Context) {
	// 查询所有的社区(community_id, community_name) 以列表形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		if err == logic.ErrorNoData {
			ResponseError(c, CodeNoData)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		if err == logic.ErrorNoData {
			ResponseError(c, CodeNoData)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

