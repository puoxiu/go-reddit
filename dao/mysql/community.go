package mysql

import (
	"fmt"
	"web-app/models"
)

func GetCommunityList() ([]*models.Community, error) {
	sqlStr := "select community_id, community_name from community"

	var communityList []*models.Community
	if err := db.Select(&communityList, sqlStr); err != nil {
		return nil, err
	}
	return communityList, nil
}

func GetCommunityDetailByID(id int64) (communityDetail *models.CommunityDetail,err error) {
	communityDetail = new(models.CommunityDetail)
	fmt.Println("id:", id)
	sqlStr := `select community_id, community_name, introduction, create_time
				from community where community_id = ?`
	if err := db.Get(communityDetail, sqlStr, id); err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	return
}
