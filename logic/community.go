package logic

import (
	"database/sql"
	"web-app/dao/mysql"
	"web-app/models"

	"go.uber.org/zap"
)


func GetCommunityList() (communities []*models.Community ,err error) {
	communities, err = mysql.GetCommunityList()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNoData
		}
		zap.L().Error("mysql.GetCommunityList failed", zap.Error(err))
		return nil, ErrorServerBusy
	}
	return
}

func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	community, err = mysql.GetCommunityDetailByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNoData
		}
		zap.L().Error("mysql.GetCommunityByID failed", zap.Error(err))
		return nil, ErrorServerBusy
	}
	return
}