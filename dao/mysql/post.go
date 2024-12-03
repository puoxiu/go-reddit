package mysql

import (
	"strings"
	"web-app/models"

	"github.com/jmoiron/sqlx"
)

func InsertPost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id)
			 values(?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostList(page, size int64) (data []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			 from post
			 ORDER BY create_time DESC
			 limit ? offset ?`
	err = db.Select(&data, sqlStr, size, (page-1)*size)
	return
}

func GetPostByID(id int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			 from post where post_id = ?`
	err = db.Get(data, sqlStr, id)
	return
}

// GetPostListByIDs 根据多个id查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			 from post where post_id in (?) order by FIND_IN_SET(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	// 重新绑定查询的参数, 防止sql注入
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}