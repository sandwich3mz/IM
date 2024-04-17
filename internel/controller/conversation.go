package controller

import (
	"IM/configs"
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/db/ent/friend"
	"IM/internel/db/ent/msg"
	"IM/internel/db/ent/user"
	"IM/internel/types"
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

var bucket = configs.Conf.MinioConfig.Bucket

// ListConversation 获取所有对话
func ListConversation(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	conversations, err := db.DB.Friend.Query().
		Where(friend.OwnerUserIDEQ(userID),
			friend.DeletedAt(utils.ZeroTime),
			friend.RelationshipEQ(types.RelationshipTypeFriend)).
		WithFriendUser(func(query *ent.UserQuery) {
			query.Select(user.FieldAvatar, user.FieldNickname)
		}).
		Order(ent.Desc(friend.FieldLastTalkAt)).
		All(c)
	if err != nil {
		logrus.Errorf("ListConversation failed, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, conversations)
}

// GetMessage 拉取聊天内容
func GetMessage(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.GetMessageReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	if err := c.ShouldBindUri(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	query := db.DB.Msg.Query().Where(
		msg.Or(msg.And(msg.SendID(userID), msg.ReceiveIDEQ(req.ID)),
			msg.And(msg.SendIDEQ(req.ID), msg.ReceiveID(userID))),
		msg.DeletedAt(utils.ZeroTime)).
		Order(ent.Desc(msg.FieldSendAt))

	total, err := query.Count(c)
	if err != nil {
		logrus.Errorf("GetMessage failed, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	msgs, err := query.
		Offset((req.PageIndex - 1) * req.PageSize).
		Limit(req.PageSize).
		All(c)
	if err != nil {
		logrus.Errorf("CreateConversation failed, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}
	var resp types.PaginateResp
	resp.List = msgs
	resp.PageIndex = req.PageIndex
	resp.PageSize = req.PageSize
	resp.Total = total

	response.RespSuccess(c, resp)
}

func UploadPic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logrus.Errorf("UploadPic failed, err: %v", err)
		response.RespError(c, code.ServerErrUpload)
		return
	}

	fileObj, err := file.Open()
	if err != nil {
		logrus.Errorf("UploadPic failed, err: %v", err)
		return
	}

	file.Filename = fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), file.Filename)

	err = utils.MinioClientGlobal.UploadFile(bucket, file.Filename, fileObj, file.Size)
	if err != nil {
		logrus.Errorf("UploadPic failed, err: %v", err)
		response.RespError(c, code.ServerErrUpload)
		return
	}

	url, err := utils.MinioClientGlobal.GetUrl(bucket, file.Filename, 7*time.Hour)
	if err != nil {
		logrus.Errorf("get url failed, err: %v", err)
		response.RespError(c, code.ServerErr)
		return
	}
	response.RespSuccessWithMsg(c, url, "上传成功")
}
