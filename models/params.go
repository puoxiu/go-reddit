package models


const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID string `json:"post_id" binding:"required"`		// 帖子id
	Direction int8 		`json:"direction" binding:"oneof=1 0 -1"` 	// 1赞成 -1反对
}


type ParamPostList struct {
	Page int64 `json:"page" form:"page"`
	Size int64 `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
	CommunityID int64 `json:"community_id" form:"community_id"`	// 社区id可以为空
}