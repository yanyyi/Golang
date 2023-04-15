package logic

import (
	"Reddit/dao/mysql"
	"Reddit/dao/redis"
	"Reddit/models"
	"Reddit/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1. 生成post id
	p.ID = snowflake.GenID()
	// 2. 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 3. 返回
}

func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostBtId(pid) failed", zap.Error(err))
		return
	}
	user, err := mysql.GetUserByUserId(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByUserId(post.AuthorID) failed", zap.Error(err))
		return
	}
	community, _ := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return data, nil
}

// GetPostList 获取帖子列表 ApiPostDatail
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByUserId(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Error(err))
			continue
		}
		//根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailedByID(post.CommunityID) failed", zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
	}
	return
}
