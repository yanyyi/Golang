package mysql

import (
	"Reddit/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community "
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in database")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (CommunityDetail *models.CommunityDetail, err error) {
	CommunityDetail = new(models.CommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time " +
		" from community" +
		" where community_id=?"
	err = db.Get(CommunityDetail, sqlStr, id)
	return CommunityDetail, err
}
