package message_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容
}

//发布消息

func (MessageApi) MessageCreateView(c *gin.Context) {
	//当前用户发布消息
	//senduserid一定是登录人的id
	//要验证一下，后期处理

	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var senUser, recvUser models.UserModel
	fmt.Println(cr.SendUserID, cr.RevUserID, cr.Content)
	err = global.DB.Take(&senUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMessage("发送人不存在", c)
		return
	}
	err = global.DB.Take(&recvUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMessage("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.MessageModel{
		SendUserID:       senUser.ID,
		RevUserID:        recvUser.ID,
		SendUserAvatar:   senUser.Avatar,
		Content:          cr.Content,
		SendUserNickName: senUser.NickName,
		RevUserNickName:  recvUser.NickName,
		IsRead:           false,
		RevUserAvatar:    recvUser.Avatar,
	}).Error
	if err != nil {
		res.OkWithMessage("消息发送失败", c)
		return
	}
	res.OkWithMessage("消息发送成功", c)
	return

}
