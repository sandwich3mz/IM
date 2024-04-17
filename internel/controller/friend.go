package controller

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/db/ent/friend"
	"IM/internel/db/ent/friendapply"
	"IM/internel/db/ent/friendgroup"
	"IM/internel/db/ent/user"
	"IM/internel/types"
	"IM/internel/types/enums"
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"IM/utils/db_utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// ListFriendGroup 获取好友分组列表
func ListFriendGroup(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	groups, err := db.DB.Debug().FriendGroup.Query().Where(friendgroup.OwnerIDEQ(userID)).
		WithFriends(func(query *ent.FriendQuery) {
			query.Where(friend.DeletedAt(utils.ZeroTime),
				friend.RelationshipEQ(types.RelationshipTypeFriend)).
				WithFriendUser(func(userQuery *ent.UserQuery) {
					userQuery.Select(user.FieldNickname, user.FieldAvatar, user.FieldSex)
				})
		}).All(c)
	if err != nil {
		logrus.Errorf("failed to query friend group, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, groups)
}

// ListGroupFriend 拉取指定分组中的用户信息
func ListGroupFriend(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.ListFriendGroupReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	friends, err := db.DB.Friend.Query().
		Where(friend.OwnerUserIDEQ(userID),
			friend.GroupIDEQ(req.GroupID)).
		All(c)
	if err != nil {
		logrus.Errorf("failed to query friends by friend group, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}
	response.RespSuccess(c, friends)
}

// ListFriendApply 获取好友申请列表
func ListFriendApply(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	applies, err := db.DB.FriendApply.Query().
		Where(friendapply.DeletedAt(utils.ZeroTime),
			friendapply.Or(friendapply.ToUserIDEQ(userID), friendapply.FromUserIDEQ(userID))).
		Order(ent.Desc(friendapply.FieldUpdatedAt)).
		All(c)
	if err != nil {
		logrus.Errorf("failed to query friend apply, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, applies)
}

// AddFriend 发送好友请求
func AddFriend(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.AddFriendReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 查询目标好友是否存在
	_friend, err := db.DB.User.Query().
		Where(user.EmailEQ(req.Email), user.DeletedAtEQ(utils.ZeroTime)).
		First(c)
	if ent.IsNotFound(err) {
		response.RespErrorWithMsg(c, code.NotFound, "用户不存在")
		return
	} else if err != nil {
		logrus.Errorf("failed to query user, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	// 保证好友不为本人
	if userID == _friend.ID {
		response.RespErrorWithMsg(c, code.InvalidParams, "不能加自己为好友")
		return
	}
	// 判断是否已经是好友
	exist, err := db.DB.Friend.Query().Where(friend.OwnerUserIDEQ(userID),
		friend.FriendUserIDEQ(_friend.ID),
		friend.RelationshipEQ(types.RelationshipTypeFriend)).
		Exist(c)
	if err != nil {
		logrus.Errorf("failed to query friend, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}
	if exist {
		response.RespErrorWithMsg(c, code.InvalidParams, "已经是好友")
	}

	// 获取本人信息
	_user, err := db.DB.User.Query().Where(user.IDEQ(userID), user.DeletedAtEQ(utils.ZeroTime)).First(c)
	if err != nil {
		logrus.Errorf("failed to query user, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	apply, err := db.DB.FriendApply.Create().
		SetFromUserID(userID).
		SetFromNickname(_user.Nickname).
		SetFromFaceURL(_user.Avatar).
		SetToUserID(_friend.ID).
		SetToNickname(_friend.Nickname).
		SetToFaceURL(_friend.Avatar).
		SetGroupID(req.GroupID).
		Save(c)
	if err != nil {
		logrus.Errorf("failed to add friend to user, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, apply)
}

// AcceptApply 同意好友请求
func AcceptApply(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.AcceptFriendReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	friendApply, err := db.DB.FriendApply.Query().
		Where(friendapply.IDEQ(req.ID),
			friendapply.DeletedAt(utils.ZeroTime)).
		First(c)
	if err != nil {
		logrus.Errorf("failed to query friend apply, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	// 确认还不是好友关系
	exist, err := db.DB.Friend.Query().
		Where(
			friend.Or(
				friend.And(friend.OwnerUserIDEQ(friendApply.ToUserID), friend.FriendUserIDEQ(friendApply.FromUserID)),
				friend.And(friend.OwnerUserIDEQ(friendApply.FromUserID), friend.FriendUserIDEQ(friendApply.ToUserID)),
			), friend.RelationshipEQ(types.RelationshipTypeFriend)).
		Exist(c)
	if err != nil {
		logrus.Errorf("failed to query friend, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}
	if exist {
		response.RespErrorWithMsg(c, code.InvalidParams, "已经是好友关系")
		return
	}

	txErr := db_utils.WithTx(c, nil, func(tx *ent.Tx) error {

		// 更新好友请求状态
		err = tx.FriendApply.Update().
			Where(friendapply.IDEQ(req.ID),
				friendapply.DeletedAt(utils.ZeroTime),
				friendapply.Result(enums.ApplyTypePending)).
			SetResult(enums.ApplyTypeAgree).
			Exec(c)
		if err != nil {
			logrus.Errorf("failed to update friend apply, err: %v", err)
			return err
		}

		// 添加接收者好友信息
		_, err = tx.Friend.Create().
			SetOwnerUserID(userID).
			SetFriendUserID(friendApply.FromUserID).
			SetRelationship(types.RelationshipTypeFriend).
			SetRemark("").
			SetGroupID(req.GroupID).
			SetLastTalkAt(time.Now()).
			Save(c)
		if err != nil {
			logrus.Errorf("failed to add friend, err: %v", err)
			return err
		}

		// 添加发起者好友信息
		_, err = tx.Friend.Create().
			SetOwnerUserID(friendApply.FromUserID).
			SetFriendUserID(friendApply.ToUserID).
			SetRelationship(types.RelationshipTypeFriend).
			SetRemark("").
			SetGroupID(friendApply.GroupID).
			SetLastTalkAt(time.Now()).
			Save(c)
		if err != nil {
			logrus.Errorf("failed to add friend, err: %v", err)
			return err
		}
		return nil
	})
	if txErr != nil {
		response.RespError(c, code.ServerErrDB)
		return
	}
	response.RespSuccess(c, "添加好友成功")
}

// RefuseApply 拒绝好友请求
func RefuseApply(c *gin.Context) {
	var req types.RefuseFriendReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	err := db.DB.FriendApply.UpdateOneID(req.ID).SetResult(enums.ApplyTypeRefuse).Exec(c)
	if err != nil {
		logrus.Errorf("failed to update friend apply, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, "拒绝好友请求成功")
}

// AddFriendGroup 添加好友分组
func AddFriendGroup(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.AddFriendGroupReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	friendGroup, err := db.DB.FriendGroup.Create().
		SetOwnerID(userID).
		SetGroupName(req.Name).
		Save(c)
	if err != nil {
		logrus.Errorf("failed to add friend group, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}
	response.RespSuccess(c, friendGroup)
}

// DeleteFriend 删除好友
func DeleteFriend(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.DeleteFriendReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	txErr := db_utils.WithTx(c, nil, func(tx *ent.Tx) error {
		err := tx.Friend.Update().
			Where(friend.Or(
				friend.And(friend.OwnerUserIDEQ(userID), friend.FriendUserIDEQ(req.ID)),
				friend.And(friend.OwnerUserIDEQ(req.ID), friend.FriendUserIDEQ(userID)),
			)).
			SetRelationship(types.RelationshipTypeDelete).
			Exec(c)
		if err != nil {
			logrus.Errorf("failed to update friend, err: %v", err)
			return err
		}
		return nil
	})
	if txErr != nil {
		response.RespError(c, code.ServerErrDB)
		return
	}
	response.RespSuccess(c, "删除好友成功")
}
