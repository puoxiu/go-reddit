package logic

import (
	"database/sql"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/snowflake"

	"go.uber.org/zap"
)


func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GetID()
	err = mysql.InsertPost(p)
	if err != nil {
		zap.L().Error("mysql.CreatePost(p) failed", zap.Error(err))
		return ErrorServerBusy
	}
	return
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNoData
		}
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return nil, ErrorServerBusy
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
			continue
		}
		// 查询社区信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName: user.UserName,
			Post:       post,
			CommunityDetail:  community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostDetailByID(id int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNoData
		}
		zap.L().Error("mysql.GetPostByID() failed", zap.Error(err))
		return nil, ErrorServerBusy
	}

	// 查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID() failed", zap.Error(err))
		return
	}

	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName: user.UserName,
		Post:       post,
		CommunityDetail:  community,
	}

	return
}
	