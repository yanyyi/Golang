package controller

import (
	"Reddit/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("Logic get community list failed ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	IdStr := c.Param("id")
	Id, _ := strconv.ParseInt(IdStr, 10, 64)
	data, err := logic.GetCommunityDetail(Id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
