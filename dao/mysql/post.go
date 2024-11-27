package mysql

import (
	"web-app/models"
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

