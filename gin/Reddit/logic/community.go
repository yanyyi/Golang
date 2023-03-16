package logic

import (
	"Reddit/dao/mysql"
	"Reddit/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetail(id)
}
