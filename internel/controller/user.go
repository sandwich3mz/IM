package controller

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"IM/internel/db/ent/user"
	"IM/internel/dto"
	"IM/internel/myRedis"
	"IM/internel/types"
	"IM/internel/types/enums"
	"IM/response"
	"IM/response/code"
	"IM/utils"
	"IM/utils/db_utils"
	_jwt "IM/utils/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	var req types.SendCodeReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 生成随机验证码
	randomCode := utils.GenerateRegisterCode(types.RegisterCodeLength)

	// 保存到 myRedis
	rKey := fmt.Sprintf(types.RedisEmailCode, req.Action, req.Email)
	err := myRedis.Client.Set(c, rKey, randomCode, types.RedisEmailCodeExpireTime*time.Minute).Err()
	if err != nil {
		logrus.Errorf("failed to set email(%s) code(%s) to myRedis, err: %v", rKey, randomCode, err)
		response.RespError(c, code.ServerErrCache)
		return
	}

	// 发送验证码
	err = utils.SendMail(req.Email, randomCode)
	if err != nil {
		response.RespError(c, code.ServerErrThirdPartyAPI)
		return
	}

	response.RespSuccess(c, "发送验证码成功")
}

// Register 注册
func Register(c *gin.Context) {
	var req types.RegisterReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 验证两次密码是否相同
	if req.Password != req.ConfirmPwd {
		logrus.Errorf("Passwords are inconsistent")
		response.RespError(c, code.InvalidParams)
		return
	}

	rKey := fmt.Sprintf(types.RedisEmailCode, enums.ActionTypeRegister, req.Email)
	rCode, err := myRedis.Client.Get(c, rKey).Result()
	if errors.Is(err, redis.Nil) {
		response.RespError(c, code.InvalidRequest)
		return
	} else if err != nil {
		response.RespError(c, code.ServerErrCache)
		return
	}

	// 判断验证码是否正确
	if rCode != req.Code {
		logrus.Errorf("code error")
		response.RespError(c, code.EmailCodeErr)
		return
	}

	var _user *ent.User
	var exist bool
	txErr := db_utils.WithTx(c, nil, func(tx *ent.Tx) error {
		// 判断用户是否存在
		exist, err = tx.User.Query().Where(user.Email(req.Email)).Exist(c)
		if err != nil {
			logrus.Errorf("failed to query user, err: %v", err)
			return err
		}

		if exist {
			response.RespErrorWithMsg(c, code.SourceExist, "邮箱已被注册")
			return nil
		} else {
			// 用户不存在，尝试创建用户
			rKey := fmt.Sprintf(types.RedisEmailCode, enums.ActionTypeRegister, req.Email)
			rCode, err := myRedis.Client.Get(c, rKey).Result()
			if err != nil {
				logrus.Errorf("failed to get code from myRedis, err: %v", err)
				return err
			}
			if rCode != req.Code {
				response.RespErrorWithMsg(c, code.InvalidParams, "验证码错误")
				return nil
			}

			_user, err = tx.User.Create().
				SetNickname(req.NickName).
				SetPassword(req.Password).
				SetStatus(enums.UserStatusOnline).
				SetEmail(req.Email).
				SetLastOnlineAt(time.Now()).
				SetAvatar("").
				Save(c)
			if err != nil {
				logrus.Errorf("failed to create user, err: %v", err)
				return err
			}
			// 创建默认好友分组
			_, err = tx.FriendGroup.Create().
				SetOwnerID(_user.ID).
				SetGroupName(types.DefaultGroupName).
				Save(c)
			if err != nil {
				logrus.Errorf("failed to create default group, err: %v", err)
				return err
			}
		}
		return nil
	})
	if txErr != nil {
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, nil)
}

// Login 登录
func Login(c *gin.Context) {
	var req types.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	_user, err := db.DB.User.Query().Where(user.EmailEQ(req.Email), user.DeletedAtEQ(utils.ZeroTime)).First(c)
	if ent.IsNotFound(err) {
		logrus.Errorf("%s not found", req.Email)
		response.RespError(c, code.NotFound)
		return
	} else if err != nil {
		logrus.Errorf("failed to query user, err: %v", err)
		return
	}

	switch req.Way {
	// 使用密码登录
	case enums.LoginWayPwd:
		// 验证密码
		if _user.Password != req.Password {
			logrus.Errorf("user(%d) login password error", _user.ID)
			response.RespError(c, code.PwdError)
			return
		}
	// 使用验证码登录
	case enums.LoginWayCode:
		// 验证验证码
		rKey := fmt.Sprintf(types.RedisEmailCode, enums.ActionTypeLogin, req.Email)
		rCode, err := myRedis.Client.Get(c, rKey).Result()
		if !errors.Is(err, redis.Nil) {
			if err != nil {
				response.RespError(c, code.ServerErrCache)
				return
			}
		} else {
			response.RespError(c, code.InvalidParams)
			return
		}

		// 判断验证码是否正确
		if rCode != req.Code {
			logrus.Errorf("password error")
			response.RespError(c, code.InvalidKey)
			return
		}
	}

	// 通过验证 发放 JWT
	jwt, err := _jwt.GenerateJwt(_user.ID)
	if err != nil {
		logrus.Errorf("failed to generate jwt, err: %v", err)
		response.RespError(c, code.ServerErr)
		return
	}

	var resp types.LoginResp
	resp.Token = jwt

	resp.User.ID = _user.ID
	resp.User.NickName = _user.Nickname
	resp.User.Email = _user.Email
	resp.User.Status = _user.Status
	resp.User.LastOnlineAt = _user.LastOnlineAt
	resp.User.Avatar = _user.Avatar

	response.RespSuccess(c, resp)
}

// ModifyPwd 修改密码
func ModifyPwd(c *gin.Context) {
	var req types.ModifyPwdReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	// 验证两次密码是否相同
	if req.Password != req.ConfirmPwd {
		logrus.Errorf("Passwords are inconsistent")
		response.RespError(c, code.InvalidParams)
		return
	}

	rKey := fmt.Sprintf(types.RedisEmailCode, enums.ActionTypeModify, req.Email)
	rCode, err := myRedis.Client.Get(c, rKey).Result()
	if err == redis.Nil {
		response.RespError(c, code.InvalidRequest)
		return
	} else if err != nil {
		response.RespError(c, code.ServerErrCache)
		return
	}

	// 判断验证码是否正确
	if rCode != req.Code {
		logrus.Errorf("code error")
		response.RespError(c, code.EmailCodeErr)
		return
	}

	_, err = db.DB.User.Update().
		Where(user.EmailEQ(req.Email), user.DeletedAtEQ(utils.ZeroTime)).
		SetPassword(req.Password).
		Save(c)
	if err != nil {
		logrus.Errorf("failed to update user, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	response.RespSuccess(c, nil)
}

// GetOneInfo 获取用户信息
func GetOneInfo(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	_user, err := db.DB.User.Query().Where(user.IDEQ(userID)).First(c)
	if err != nil {
		logrus.Errorf("failed to get user(%d), err: %v", userID, err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	var res dto.UserDto
	res.NickName = _user.Nickname
	res.Email = _user.Email
	res.Status = _user.Status
	res.LastOnlineAt = _user.LastOnlineAt
	res.Sex = _user.Sex
	res.Avatar = _user.Avatar

	response.RespSuccess(c, res)
}

// UpdateOneInfo 更新用户信息
func UpdateOneInfo(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.RespError(c, code.UnLogin)
		return
	}

	var req types.UpdateInfoReq
	if err := c.ShouldBind(&req); err != nil {
		response.RespErrorInvalidParams(c, err)
		return
	}
	logrus.Debugf("req: %+v", req)

	_user, err := db.DB.User.
		UpdateOneID(userID).
		SetAvatar(req.Avatar).
		SetSex(req.Sex).
		SetNickname(req.NickName).
		Save(c)
	if err != nil {
		logrus.Errorf("failed to update user info, err: %v", err)
		response.RespError(c, code.ServerErrDB)
		return
	}

	var res dto.UserDto
	res.NickName = _user.Nickname
	res.Email = _user.Email
	res.Status = _user.Status
	res.LastOnlineAt = _user.LastOnlineAt
	res.Avatar = _user.Avatar
	res.Sex = _user.Sex

	response.RespSuccess(c, res)
}
