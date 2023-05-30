package controller

import (
	"Reddit/logic"
	"Reddit/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 投票
type VoteData struct {
	// UserID 可直接从请求中(c *gin.Context)获取当前用户
	PostID    int64 `json:"post_id,string"`   //帖子id
	Direction int   `json:"direction,string"` //赞成票(1)还是反对票(-1)
}

func PostVoteController(c *gin.Context) {
	//参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取当前请求的用户id
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体的投票业务
	if err = logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		//ResponseError(c, CodeServerBusy)
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}
